package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/temideewan/go-social/internal/store"
)

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
				return
			}
			// parse it -> get base 64
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}
			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.unauthorizedBasicErrorResponse(w, r, err)
				return
			}

			// check the credentials
			user := app.config.auth.basic.user
			pass := app.config.auth.basic.pass
			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != user || creds[1] != pass {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("Invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
			return
		}

		// ensure bearer token is passed in
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}
		fmt.Println("Token:", parts[1])
		token := parts[1]
		jwtToken, err := app.authenticator.ValidateToken(token)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}
		ctx := r.Context()
		user, err := app.store.Users.GetById(ctx, userId)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}
		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) checkPostOwnership(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromContext(r)
		post := getPostFromCtx(r)
		// if it is the user's post
		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}
		// role precedence check
		allowed, err := app.checkRolePrecedence(r.Context(), user, requiredRole)

		if err != nil {
			switch err {
			case store.ErrNotFound:
				app.badRequestResponse(w, r, err)
				return
			default:
				app.internalServerError(w, r, err)
				return
			}
		}

		if !allowed {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("You are not allowed to perform this action"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkRolePrecedence(ctx context.Context, user *store.User, requiredRole string) (bool, error) {
	role, err := app.store.Roles.GetByName(ctx, requiredRole)

	if err != nil {
		return false, err
	}
	return user.Role.Level >= role.Level, nil
}
