package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-shopee/config"
	"golang-shopee/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func getGithubOauthConfig() *oauth2.Config {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	redirectURL := os.Getenv("GITHUB_REDIRECT_URL")

	log.Printf("Debug OAuth Config - ClientID: %s, RedirectURL: %s", clientID, redirectURL)

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

// GithubLogin initiates the GitHub OAuth login
func GithubLogin(c *gin.Context) {
	url := getGithubOauthConfig().AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GithubCallback handles the GitHub OAuth callback
func GithubCallback(c *gin.Context) {
	code := c.Query("code")
	githubConfig := getGithubOauthConfig()
	token, err := githubConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	client := githubConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var githubUser struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&githubUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	// Save or update user in database
	var user models.User
	result := config.DB.Where("github_id = ?", fmt.Sprintf("%d", githubUser.ID)).First(&user)

	if result.Error != nil {
		// Create new user
		user = models.User{
			Name:      githubUser.Name,
			Email:     githubUser.Email,
			AvatarURL: githubUser.AvatarURL,
			GithubID:  fmt.Sprintf("%d", githubUser.ID),
		}
		if user.Name == "" {
			user.Name = githubUser.Login
		}
		config.DB.Create(&user)
	} else {
		// Update existing user
		user.Name = githubUser.Name
		user.Email = githubUser.Email
		user.AvatarURL = githubUser.AvatarURL
		if user.Name == "" {
			user.Name = githubUser.Login
		}
		config.DB.Save(&user)
	}

	// Generate JWT or session here (omitted for simplicity, just returning user info)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token.AccessToken,
	})
}
