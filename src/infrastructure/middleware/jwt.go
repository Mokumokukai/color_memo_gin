package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type jwtMiddlewareHandler struct {
	client *auth.Client
}
type JwtMiddlewareHandler interface {
	SetJWTToken() gin.HandlerFunc
}

func NewJWTMiddlwareHandler(filePath string, ProjectID string) JwtMiddlewareHandler {

	return &jwtMiddlewareHandler{get_client(filePath, ProjectID)}
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

func (handler *jwtMiddlewareHandler) SetJWTToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		bearer_tok := c.Request.Header.Get("Authorization")
		splited_tok := strings.Split(bearer_tok, " ")
		if len(splited_tok) != 2 {
			c.JSON(http.StatusUnauthorized, fmt.Errorf("Error: %s", "this is not bearer token format."))
			c.Abort()
			return
		}
		token, err := handler.client.VerifyIDToken(context.Background(), splited_tok[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, fmt.Errorf("error verifying ID token: %v", err))
			c.Abort()
			return

		}
		c.Set("UID", token.UID)
		c.Next()
	}

}
