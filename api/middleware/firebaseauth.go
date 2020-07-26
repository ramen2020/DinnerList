package middleware

import (
	"context"
	"net/http"

	"app/config"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func unAuthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	})
	c.Abort()
}

func FirebaseAuth(c *gin.Context) {

	// Firebase SDK のセットアップ
	opt := option.WithCredentialsFile(config.Config.FIREBASE)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		firebaseUninitialized(c)
		return
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		firebaseUninitialized(c)
		return
	}

	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		unAuthorized(c)
		return
	}
	if authorizationHeader[:7] != "Bearer " {
		unAuthorized(c)
		return
	}
	tokenString := authorizationHeader[7:]

	// JWT が無効かどうか
	token, err := auth.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		unAuthorized(c)
		return
	}
	c.Set("uid", token.UID)
}
