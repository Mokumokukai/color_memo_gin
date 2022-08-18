package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type jwtMiddlewareHandler struct {
	client *auth.Client
	db     *gorm.DB
}
type JwtMiddlewareHandler interface {
	SetUID() gin.HandlerFunc
	SetUserID() gin.HandlerFunc
}

func NewJWTMiddlwareHandler(filePath string, ProjectID string, db *gorm.DB) JwtMiddlewareHandler {

	return &jwtMiddlewareHandler{get_client(filePath, ProjectID), db}
}
func get_client(filePath string, ProjectID string) *auth.Client {
	opt := option.WithCredentialsFile(filePath) //("/go/src/firebase-adminsdk.json")
	config := &firebase.Config{ProjectID: ProjectID}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client

}
func parse(bearer_token string, client *auth.Client) (string, error) {
	splited_tok := strings.Split(bearer_token, " ")
	if len(splited_tok) != 2 {
		return "", fmt.Errorf("Error: %s", "this is not bearer token format.")
	}
	token, err := client.VerifyIDToken(context.Background(), splited_tok[1])
	if err != nil {
		return "", fmt.Errorf("error verifying ID token: %v", err)

	}
	return token.UID, nil
}

func (handler *jwtMiddlewareHandler) SetUID() gin.HandlerFunc {

	return func(c *gin.Context) {
		bearer_token := c.Request.Header.Get("Authorization")
		uid, err := parse(bearer_token, handler.client)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Set("UID", uid)
		c.Next()
	}
}
func (handler *jwtMiddlewareHandler) SetUserID() gin.HandlerFunc {

	return func(c *gin.Context) {
		bearer_token := c.Request.Header.Get("Authorization")
		uid, err := parse(bearer_token, handler.client)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		var user models.User
		result := handler.db.Where("uid = ?", uid).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, fmt.Errorf("You are not signed up yet."))
			c.Abort()
			return
		}
		c.Set("user_id", user.ID)
		c.Next()
	}
}
