package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error", "error", err.Error(), "path", r.URL.Path, "method", r.Method)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request", "error", err.Error(), "path", r.URL.Path, "method", r.Method)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("not found error", "error", err.Error(), "path", r.URL.Path, "method", r.Method)
	writeJSONError(w, http.StatusNotFound, "not found")
}
func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("conflict", "error", err.Error(), "path", r.URL.Path, "method", r.Method)
	writeJSONError(w, http.StatusConflict, err.Error())
}
func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("unauthorized error", "error", err.Error(), "path", r.URL.Path, "method", r.Method)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}
func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("unauthorized error", "error", err.Error(), "path", r.URL.Path, "method", r.Method)
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}
