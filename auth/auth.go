package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ochom/sdp-lib/utils"
	"golang.org/x/crypto/bcrypt"
)

//ContextKey ...
type ContextKey string

//ContextUser ...
const ContextUser ContextKey = "user"

var secreteKey string = utils.GetEnvOrDefault("AUTH_SECRETE_KEY", "nmU9AHJEyhZODIWT4sLBZtCX3k6fGhEY")

//SignedDetails ...
type SignedDetails struct {
	UID       string
	FirstName string
	LastName  string
	Email     string
	Mobile    string
	jwt.StandardClaims
}

//ValidateToken validates the jwt token
func ValidateToken(signedToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secreteKey), nil
		},
	)

	if err != nil {
		return nil, errors.New("invalid auth token")
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, errors.New("invalid auth token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("expired auth token")
	}

	return claims, nil
}

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) (*string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}
	resp := string(hashed)
	return &resp, nil
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	return err == nil
}

// GenerateAuthTokens generates both the detailed token and refresh token
func GenerateAuthTokens(claims SignedDetails) (string, string, error) {
	expiry := time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiry,
	}

	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secreteKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secreteKey))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

//Middleware is the authentication middleware for basic jwt authentication
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaders := c.Request.Header.Get("Authorization")

		if authHeaders == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		splitToken := strings.Split(authHeaders, " ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided. e.g `Bearer <token>`"})
			c.Abort()
			return
		}

		token := splitToken[1]
		claims, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		ctx := context.WithValue(c, ContextUser, claims)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

//GetAuthClaims retrieves a user from context
func GetAuthClaims(ctx context.Context) (*SignedDetails, error) {
	data := ctx.Value(ContextUser)

	claims, ok := data.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("claims not in context")
	}
	return claims, nil
}
