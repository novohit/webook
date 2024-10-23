package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey []byte

type MyCustomClaims struct {
	Identify string `json:"identify"`
	jwt.RegisteredClaims
}

func initSecret() {
	mySigningKey = []byte("secret")
}

func GenToken(identify string) (string, error) {
	initSecret()
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		identify,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	fmt.Printf("identify: %v\n", claims.Identify)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	//fmt.Println(tokenString, err)
	return tokenString, err
}

func VerifyToken(tokenString string) (*MyCustomClaims, error) {
	initSecret()
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, errors.New("token is invalid")
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.UserID, claims.RegisteredClaims.Issuer)
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}
