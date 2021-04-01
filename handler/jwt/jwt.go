package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/Artpou/wiki_golang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("cle_tres_secrete")

func GetToken(r *http.Request) (*jwt.Token, error) {
	tknCookie, err := r.Cookie("token")
	tknHeader := ""

	if err == http.ErrNoCookie {
		// Try to get cookie with Header Authorization
		tknHeader = strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	} else {
		tknHeader = tknCookie.Value
	}

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	return jwt.ParseWithClaims(tknHeader, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}

func SetToken(user models.User, w http.ResponseWriter) (*http.Cookie, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	tknCookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	http.SetCookie(w, tknCookie)

	return tknCookie, err
}

func GetRole(r *http.Request) (models.Role, error) {
	claims := &Claims{}
	tknHeader, err := GetToken(r)
	if err != nil {
		return models.UserRole, err
	}
	jwt.ParseWithClaims(tknHeader.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claims.Role, nil
}
