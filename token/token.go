package token

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"strings"
	"time"
)

func Generatetoken(username string) (string, error) {
	tokenexpirationtime, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_TIME"))
	if err != nil {
		return "", err
	}

	var claims = jwt.MapClaims{}

	claims["iis"] = "kiki"
	claims["sub"] = "forfreelink"
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(tokenexpirationtime)).Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(os.Getenv("API_SECRET"))
}

func Tokenvalid(tokenString string) error {
	if _, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	}); err != nil {
		return err
	}

	return nil
}

func Extractclaims(tokenstr string) (*jwt.MapClaims, error) {
	if token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	}); err != nil {
		return nil, err
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return &claims, nil
		} else {
			return nil, errors.New("invalid token")
		}
	}
}

func Extracttokenstr(ctx *gin.Context) string {
	if len(strings.Split(ctx.GetHeader("Authorization"), " ")) == 2 {
		return strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	}
	return ""
}
