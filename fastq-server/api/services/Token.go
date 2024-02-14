package services

import (
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	// "github.com/bentennison/standardServer/cmd/api/models"
	"github.com/bentenison/fastq-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type idTokenClaims struct {
	//* this struct field should be capitalized to avoid nil user in token decode
	User *models.UserAccount `json:"user"`
	jwt.StandardClaims
}

func generateIDToken(u *models.UserAccount, key *rsa.PrivateKey, exp int64) (string, error) {
	unixTime := time.Now().Unix()
	// log.Println(key)
	tokenExp := unixTime + exp
	claims := idTokenClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExp,
			IssuedAt:  unixTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Println("error signing token :", err)
		return "", err
	}
	return ss, nil

}

type refreshTokenData struct {
	SS        string
	ID        uuid.UUID
	ExpiresIn time.Duration
}
type refreshTokenClaims struct {
	Uuid string `json:"uid"`
	jwt.StandardClaims
}
type resetTokenClaims struct {
	Id string `json:"Id"`
	jwt.StandardClaims
}

func generateRefreshToken(uid uuid.UUID, key string, exp int64) (*refreshTokenData, error) {
	currentTime := time.Now()
	expireTime := currentTime.Add(time.Duration(exp) * time.Second)
	tokenID, err := uuid.NewRandom()
	if err != nil {
		log.Println("error generating uuid:", err)
		return &refreshTokenData{}, err
	}
	claims := refreshTokenClaims{
		Uuid: uid.String(),
		StandardClaims: jwt.StandardClaims{
			Id:        tokenID.String(),
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("error in signing refresh token:", err)
		return &refreshTokenData{}, err
	}

	return &refreshTokenData{
		SS:        ss,
		ID:        tokenID,
		ExpiresIn: expireTime.Sub(currentTime),
	}, nil

}

func validateIDToken(tokenstring string, publicKey *rsa.PublicKey) (*idTokenClaims, error) {
	claims := &idTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("JWT token is invalid")
	}
	claims, ok := token.Claims.(*idTokenClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing the claims from token")
	}
	return claims, nil

}
func validateRefreshToken(tokenstring string, key string) (*refreshTokenClaims, error) {
	claims := &refreshTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("JWT token is invalid")
	}
	claims, ok := token.Claims.(*refreshTokenClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing the claims from token")
	}
	return claims, nil

}
