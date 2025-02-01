package middlewares

import (
	"net/http"

	"example.com/event-booking/api-test/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization") // always return string even though empty.
	if token == "" {
		// Remeber that this function is executed in the middle of a request.
		// if we use context.JSON() here, we will send back multiple response
		// it aborts current request and no other request handlers will be executed later.
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
	}

	// Set() adds some data to this context.
	context.Set("userId", userId)

	// ensure that the next request handler will execute correctly.
	context.Next()
}
