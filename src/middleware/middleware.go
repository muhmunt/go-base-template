package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const TOKEN_FAILED = "Token Invalid!"
const HEADER_ABSENT = "Authorization Header Absent!"

func HandleAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := extractBearerToken(c.Request)

		actualToken := "TES"

		validate := bcrypt.CompareHashAndPassword([]byte(token), []byte(actualToken))
		if validate != nil {
			response := gin.H{"message": TOKEN_FAILED}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			fmt.Println("Test")
			return
		}
		c.Next()
	}
}

func extractBearerToken(r *http.Request) (string, bool) {
	header := r.Header.Get("Authorization")

	if header == "" || !strings.HasPrefix(header, "Bearer ") {

		return "", false
	}

	token := strings.TrimPrefix(header, "Bearer ")

	return token, true
}

func Authorization(level string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("client-authorization")
		emptyToken := errorTokenEmpty(token, c)

		if emptyToken {
			return
		}
		
		level := validateLevel(token, level, c)

		if !level {
			return
		}

		c.Next()
	}
}

// func PatientAuthorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		token := r.Header.Get("client-authorization")

// 		emptyToken := errorTokenEmpty(token, w)

// 		if emptyToken {
// 			return
// 		}

// 		level := validateLevel(token, "PATIENT", w)

// 		if !level {
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func errorTokenEmpty(token string, c *gin.Context) bool {
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": TOKEN_FAILED})

		return true
	}

	return false
}

type Token struct {
	ID     string
	UserID string
	Type   string
}

func validateLevel(token string, level string, c *gin.Context) bool {
	if token != "" {
		var userData Token

		data, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			fmt.Println("Error decoding Base64:", err)
			return false
		}
		// Unmarshal the JSON data into a struct
		err = json.Unmarshal(data, &userData)
		fmt.Println(userData.Type)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return false
		}

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": TOKEN_FAILED})
			c.Abort()
			return false
		}

		if userData.Type != level {
			c.JSON(http.StatusUnauthorized, gin.H{"message": TOKEN_FAILED})
			c.Abort()

			return false
		}
	}

	return true

}
