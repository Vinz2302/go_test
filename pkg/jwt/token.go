package jwt

import (
	"errors"
	"fmt"
	"rest-api/app/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	conf, _ = config.Init()
)

func ValidateToken(encodedToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

/* func Refresh(refreshToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
} */

func GenerateToken(roleId uint) (string, error) {

	claim := jwt.MapClaims{}
	claim["role_id"] = roleId
	claim["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	// if tokenName == "at" {
	// 	claim["token_name"] = "access_token"
	// 	claim["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	// } else {
	// 	claim["token_name"] = "refresh_token"
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret_key := []byte(conf.App.Secret_key)

	signedToken, err := token.SignedString(secret_key)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

/* func GenerateToken(uuID string, email string, tokenName string, roleId *int) (string, error) {

	claim := jwt.MapClaims{}
	claim["uuid"] = uuID
	claim["email"] = email
	claim["role_id"] = roleId

	if tokenName == "at" {
		claim["token_name"] = "access_token"
		claim["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	} else {
		claim["token_name"] = "refresh_token"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	secret_key := []byte(conf.App.Secret_key)

	signedToken, err := token.SignedString(secret_key)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
} */

func ExtractTokenUUID(token *jwt.Token) (string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	uid := fmt.Sprintf("%v", claims["uuid"])

	return uid, nil
}

func ExtractTokenRoleID(token *jwt.Token) (*int, error) {

	claims, _ := token.Claims.(jwt.MapClaims)
	role_id := fmt.Sprintf("%v", claims["role_id"])
	roleID, _ := strconv.Atoi(role_id)
	fmt.Println("roleid = ", roleID)

	return &roleID, nil
}
