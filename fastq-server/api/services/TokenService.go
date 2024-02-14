package services

import (
	"context"
	"crypto/rsa"
	"log"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/models"
)

// tokenService used for injecting an implementation of TokenRepository
// for use in service methods along with keys and secrets for
// signing JWTs
type tokenService struct {
	TokenRepository       models.TokenRepository
	PrivKey               *rsa.PrivateKey
	PubKey                *rsa.PublicKey
	RefreshSecret         string
	IDExpirationSecs      int64
	RefreshExpirationSecs int64
}

// TSConfig will hold repositories that will eventually be injected into this
// this service layer
type TSConfig struct {
	TokenRepository       models.TokenRepository
	PrivKey               *rsa.PrivateKey
	PubKey                *rsa.PublicKey
	RefreshSecret         string
	IDExpirationSecs      int64
	RefreshExpirationSecs int64
}

func NewTokenService(ts *TSConfig) models.TokenService {
	return &tokenService{
		TokenRepository:       ts.TokenRepository,
		PrivKey:               ts.PrivKey,
		PubKey:                ts.PubKey,
		RefreshSecret:         ts.RefreshSecret,
		IDExpirationSecs:      ts.IDExpirationSecs,
		RefreshExpirationSecs: ts.RefreshExpirationSecs,
	}
}

func (ts *tokenService) GenerateIDToken(u *models.UserAccount) {
	token, err := generateIDToken(u, ts.PrivKey, ts.IDExpirationSecs)
	if err != nil {
		log.Println("error in generating token:", err)
	}

	claims, err := validateIDToken(token, ts.PubKey)
	if err != nil {
		log.Println("error in validating token", err)
	}

	// ts.TokenRepository.SetRefreshToken(context.Background(), u.UserID)
	log.Println("claims", claims)

}

func (ts *tokenService) GeneratePairFromUser(ctx context.Context, user *models.UserAccount, prevToken string) (*models.TokenPair, error) {
	if prevToken == "" {
		idToken, err := generateIDToken(user, ts.PrivKey, ts.IDExpirationSecs)
		if err != nil {
			log.Println("error generating id token:", err)
			return nil, err
		}
		refreshToken, err := generateRefreshToken(user.UserID, ts.RefreshSecret, ts.RefreshExpirationSecs)
		if err != nil {
			log.Println("error generating refresh token:", err)
			return nil, err
		}
		// err = ts.TokenRepository.SetRefreshToken(context.Background(), user.UserID.String(), refreshToken.ID.String(), refreshToken.ExpiresIn)
		// if err != nil {
		// 	log.Println("error setting refresh token in repository:", err)
		// 	return nil, err
		// }
		return &models.TokenPair{
			IdToken: models.IdToken{
				SS: idToken,
			},
			RefreshToken: models.RefreshToken{
				SS:  refreshToken.SS,
				ID:  refreshToken.ID.String(),
				UID: user.UserID.String(),
			},
		}, nil

	}
	return nil, nil
}

func (s *tokenService) ValidateIDToken(tokenString string) (*models.UserAccount, error) {
	claims, err := validateIDToken(tokenString, s.PubKey) // uses public RSA key

	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		log.Printf("Unable to validate or parse idToken - Error: %v\n", err)
		return nil, apperrors.NewAuthorization("Unable to verify user from idToken")
	}

	return claims.User, nil
}

// ValidateRefreshToken checks to make sure the JWT provided by a string is valid
// and returns a RefreshToken if valid
func (s *tokenService) ValidateRefreshToken(tokenString string) (*models.RefreshToken, error) {
	// validate actual JWT with string a secret
	claims, err := validateRefreshToken(tokenString, s.RefreshSecret)

	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		log.Printf("Unable to validate or parse refreshToken for token string: %s\n%v\n", tokenString, err)
		return nil, apperrors.NewAuthorization("Unable to verify user from refresh token")
	}

	// Standard claims store ID as a string. I want "model" to be clear our string
	// is a UUID. So we parse claims.Id as UUID
	// tokenUUID, err := uuid.Parse(claims.Id)

	// if err != nil {
	// 	log.Printf("Claims ID could not be parsed as UUID: %s\n%v\n", claims.Id, err)
	// 	return nil, apperrors.NewAuthorization("Unable to verify user from refresh token")
	// }

	return &models.RefreshToken{
		SS:  tokenString,
		ID:  claims.Id,
		UID: claims.Uuid,
	}, nil
}
