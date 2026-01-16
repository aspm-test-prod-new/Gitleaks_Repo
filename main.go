package main

import (
	"fmt"
	"os"
)

// Fake secrets for Gitleaks testing ONLY

// AWS credentials
var awsAccessKey = "AKIA1234567890ABCDE"
var awsSecretKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

// GitHub token
var githubToken = "ghp_1234567890abcdefghijklmnopqrstuvwxyz"

// GitLab token
var gitlabToken = "glpat-abcdefghijklmnopqrstuvwxyz123456"

// Slack token
var slackToken = "xoxb-123456789012-123456789012-abcdefghijklmnopqrstuvwx"

// Google API Key
var googleAPIKey = "AIzaSyD-EXAMPLEKEY1234567890abcdef"

// JWT secret
var jwtSecret = "jwt_super_secret_key_123456"

// Stripe secret key
var stripeKey = "sk_test_51H8abc1234567890abcdefghijklmnopqrstuvwxyz"

// Database credentials
var dbConn = "postgres://admin:password123@localhost:5432/mydb"

// Generic password
var adminPassword = "SuperSecretPassword123!"

func main() {
	fmt.Println("Gitleaks vulnerability test file")

	// Access environment variable (also detectable)
	dbPassword := os.Getenv("DB_PASSWORD")
	fmt.Println("DB password:", dbPassword)

	// Prevent unused variable errors
	_ = awsAccessKey
	_ = awsSecretKey
	_ = githubToken
	_ = gitlabToken
	_ = slackToken
	_ = googleAPIKey
	_ = jwtSecret
	_ = stripeKey
	_ = dbConn
	_ = adminPassword
}
