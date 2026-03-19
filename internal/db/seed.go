package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/temideewan/go-social/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "david", "emma",
	"frank", "grace", "henry", "isabel", "jack",
	"kate", "liam", "mia", "noah", "olivia",
	"peter", "quinn", "rachel", "sam", "tyler",
	"uma", "victor", "wendy", "xavier", "yara",
	"zach", "adam", "bella", "chris", "diana",
	"ethan", "fiona", "george", "hannah", "ian",
	"julia", "kevin", "luna", "mike", "nina",
	"oscar", "penny", "quinn", "ryan", "sofia",
	"theo", "uma", "violet", "will", "xena",
	"yale", "zoe", "alex", "ben", "cora",
}

var titles = []string{
	"10 Tips for Better Code",
	"Getting Started with Go",
	"Understanding Algorithms",
	"Web Security Basics",
	"Clean Code Principles",
	"Database Design 101",
	"API Best Practices",
	"Cloud Computing Guide",
	"Docker Essentials",
	"Testing Strategies",
	"Git Workflow Tips",
	"Microservices Explained",
	"Performance Tuning",
	"DevOps Fundamentals",
	"Code Review Guidelines",
	"System Design Basics",
	"Frontend Optimization",
	"Data Structures Guide",
	"Machine Learning Intro",
	"Software Architecture",
}

var content = []string{
	"Learn the fundamentals of clean code architecture and SOLID principles for building maintainable software systems.",
	"A comprehensive guide to implementing RESTful APIs using modern best practices and security measures.",
	"Deep dive into container orchestration with Kubernetes - from basics to advanced deployment strategies.",
	"Understanding memory management in different programming languages and how to avoid common pitfalls.",
	"Building scalable microservices: patterns, practices, and real-world implementation examples.",
	"Essential data structures every programmer should know and when to use them effectively.",
	"Practical guide to implementing authentication and authorization in web applications.",
	"Optimizing database performance: indexing strategies, query optimization, and best practices.",
	"Test-driven development workflow: writing maintainable and reliable test suites.",
	"Cloud-native application development: principles, tools, and deployment strategies.",
	"Understanding concurrency patterns and their implementation across different languages.",
	"Frontend performance optimization techniques for modern web applications.",
	"Implementing secure coding practices to protect against common vulnerabilities.",
	"Design patterns in practice: real-world examples and implementation strategies.",
	"Building resilient systems: fault tolerance, circuit breakers, and error handling.",
	"Version control best practices and advanced Git workflows for team collaboration.",
	"Introduction to distributed systems: concepts, challenges, and solutions.",
	"API security best practices: authentication, rate limiting, and data validation.",
	"Continuous Integration/Continuous Deployment (CI/CD) pipeline implementation guide.",
	"Performance monitoring and optimization techniques for production systems.",
}

var tags = []string{
	"programming", "technology", "webdev", "coding", "software",
	"database", "cloud", "security", "devops", "architecture",
	"testing", "api", "frontend", "backend", "performance",
	"algorithms", "development", "engineering", "infrastructure", "design",
}

var comments = []string{
	"Great article, very informative!",
	"This helped me understand the concept better.",
	"Thanks for sharing your knowledge!",
	"Could you elaborate more on point #3?",
	"Excellent explanation of complex topics.",
	"Looking forward to more content like this.",
	"This solved my problem, thank you!",
	"Very well written and structured.",
	"I'll definitely bookmark this for future reference.",
	"The examples really helped clarify things.",
	"Interesting perspective on this topic.",
	"Amazing resource for beginners.",
	"Would love to see a follow-up post.",
	"Clear and concise explanation.",
	"This is exactly what I was looking for.",
	"Great insights, especially the practical tips.",
	"Really useful information here.",
	"The code examples are very helpful.",
	"Nice breakdown of the concepts.",
	"Can't wait to try this out!",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)

	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}
	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: content[rand.Intn(len(content))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}
func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostId:  posts[rand.Intn(len(posts))].ID,
			UserId:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
