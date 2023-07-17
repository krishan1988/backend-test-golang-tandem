// Package middleware contains an middleware for the server.
package middleware

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/KryptoKnight/backend-test-golang/api/server/schema"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

var key *rsa.PublicKey

func getPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	if key != nil {
		return key, nil
	}
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	key, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	return key, err
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("no authorization token")
	}
	// extract token from {Bearer <token>}
	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	return jwtToken[1], nil
}

// GetJwtValidationMiddleware returns the JWT validation middleware.
func GetJwtValidationMiddleware(publicKeyPath string, logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, schema.Error{
				Message: err.Error(),
			})
			return
		}

		publicKey, err := getPublicKey(publicKeyPath)
		if err != nil {
			logger.Error().Err(err).Msg("unable to get the public key")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		switch err.(type) {
		case nil:
			if !token.Valid {
				logger.Error().Msg("unable to validate the token")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
			return
		case *jwt.ValidationError:
			validationErr := err.(*jwt.ValidationError)
			switch validationErr.Errors {
			case jwt.ValidationErrorExpired:
				logger.Error().Msg("token has been expired")
				c.AbortWithStatusJSON(http.StatusBadRequest, schema.Error{Message: "token has been expired"})
				return
			default:
				logger.Error().Err(validationErr).Msg("unable to validate the token")
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

		default:
			logger.Error().Err(err).Msg("jwt parsing error")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
