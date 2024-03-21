package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// normal func version
// func AccessToken(c *gin.Context) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
// 		"exp": time.Now().Add(1 * time.Minute).Unix(),
// 	})

// 	ss, err := token.SignedString([]byte("+++signature+++"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": ss,
// 	})
// }

// middleware version
func AccessToken(signature []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
			"exp": time.Now().Add(1 * time.Minute).Unix(),
			"aud": "Touch",
		})

		ss, err := token.SignedString(signature)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": ss,
		})
	}
}
