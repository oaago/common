package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("oaago")

var data map[string]interface{}

type Claims struct {
	Key        string        `json:"key"`
	CreateTime time.Duration `json:"createTime"`
	Data       map[string]interface{}
	jwt.StandardClaims
}

func GenerateToken(key string, data map[string]interface{}) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		key,
		time.Duration(time.Now().Unix()),
		data,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "oaago",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
		fmt.Println("验证失败")
	}
	return nil, err
}
