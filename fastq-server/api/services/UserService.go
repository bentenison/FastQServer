package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/api/helpers"
	"github.com/bentenison/fastq-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
* userService acts as a struct for injecting an implementation of UserRepository
* for use in service methods
 */

type userService struct {
	ResetTokenExp  int64
	ResetSectret   string
	UserRepository models.UserRepository
	MailService    models.MailService
	// ImageRepository model.ImageRepository
}

/*
* USConfig will hold repositories that will eventually be injected into this service layer
* now it means the USConfig is used to hide the implementation of service layer
 */
type USConfig struct {
	ResetTokenExp  int64
	ResetSectret   string
	UserRepository models.UserRepository
	MailService    models.MailService
	// ImageRepository model.ImageRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) models.UserService {
	return &userService{
		UserRepository: c.UserRepository,
		ResetTokenExp:  c.ResetTokenExp,
		ResetSectret:   c.ResetSectret,
		MailService:    c.MailService,
		// ImageRepository: c.ImageRepository,
	}
}

func (us *userService) SignIn(ctx context.Context, user *models.UserAccount) error {
	userDB, err := us.UserRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Println("error getting user from db reson:", err)
		return err
	}
	log.Println(userDB)
	matched, err := comparePasswords(userDB.Password, user.Password)
	if err != nil {
		log.Println("error occured while comparing passwords reason:", err)
		return err
	}
	if !matched {
		log.Println("error passwords mismatched:", err)
		return apperrors.PasswordMismatched
	}
	*user = *userDB
	return nil
}
func (us *userService) BranchLogin(ctx context.Context, user *models.UserAccount) error {
	userDB, err := us.UserRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Println("error getting user from db reson:", err)
		return err
	}
	log.Println(userDB)
	matched, err := comparePasswords(userDB.Password, user.Password)
	if err != nil {
		log.Println("error occured while comparing passwords reason:", err)
		return err
	}
	if !matched {
		log.Println("error passwords mismatched:", err)
		return apperrors.PasswordMismatched
	}
	*user = *userDB
	return nil
}
func (us *userService) SignUp(ctx context.Context, user *models.UserAccount) error {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("error generating uid:", err)
		return err
	}
	company_code := generateRandomString(7)
	userDB, err := us.UserRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			hashedPassword, err := hashPassword(user.Password)
			if err != nil {
				log.Println("error hashing password:", err)
				return err
			}
			user.Password = hashedPassword
			user.UserID = uid
			t := time.Now()
			ts := t.Format("2006-01-02 15:04:05")
			user.CreatedAt = ts
			user.UpdatedAt = ts
			user.UpdatedBy = uid.String()
			user.CompanyCode = company_code
			err = us.UserRepository.Create(ctx, user)
			if err != nil {
				log.Println("error occured while signUp:", err)
				return err
			}
			cpu, err := helpers.GetCPUId()
			if err != nil {
				log.Println("error occured while GetCPUId :", err)
				return err
			}
			disk, err := helpers.GetHDDId()
			if err != nil {
				log.Println("error occured while GetHDDId:", err)
				return err
			}
			ip := helpers.GetOutboundIP()
			if ip == "" {
				log.Println("error occured while GetIpAddr:", err)
				return err
			}
			details := models.ServerDetails{
				ServerCPU:    cpu,
				ServerIP:     ip,
				ServerDiskId: disk,
				Id:           company_code,
			}
			err = us.UserRepository.AddServer(ctx, details)
			if err != nil {
				log.Println("error occured while Addserver:", err)
				return err
			}
			return nil
		default:
			log.Println("error getting user details reason:", err)
			return err
		}
	}
	if userDB != nil {
		return apperrors.UserAlreadyExists
	}
	return err
}
func (us *userService) ForgotPassword(ctx context.Context, email string) error {
	userDB, err := us.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		log.Println("error getting user from db reson:", err)
		return err
	}
	log.Println(userDB.UserID)
	resetToken, err := us.generateResetToken(ctx, us.ResetSectret, us.ResetTokenExp, userDB)
	if err != nil {
		log.Println("error generating reset token reason:", err)
		return err
	}
	// parese tempplates with data to s4nd email
	td := make(map[string]interface{})
	//data creation for dynamic templates
	td["email"] = email
	td["link"] = "localhost:8080/#/token"
	td["username"] = "username"
	cwd, _ := os.Getwd()
	file := filepath.Join(cwd, "/api/templates/passwordReset.html")
	err = us.MailService.ParseTemplate(file, td)
	if err != nil {
		log.Println("error parsing template:", err)
		return err
	}
	err = us.MailService.SendMail(email, "")
	if err != nil {
		log.Println("error sending mail:", err)
		return err
	}
	log.Println("token", resetToken)

	return nil
}
func (us *userService) ResetPassword(ctx context.Context, token, password string) error {
	claims, err := us.verifyResetToken(ctx, token, us.ResetSectret)
	if err != nil {
		log.Println("error verifying token , try again")
		return err
	}
	log.Println("claims", claims)
	return nil
}
func (us *userService) generateResetToken(ctx context.Context, refreshKey string, exp int64, user *models.UserAccount) (string, error) {
	currentTime := time.Now()
	expireTime := currentTime.Add(time.Duration(exp) * time.Second)
	tokenID, err := uuid.NewRandom()
	if err != nil {
		log.Println("error generating uuid:", err)
		return "", err
	}
	claims := resetTokenClaims{
		Id: user.UserID.String(),
		StandardClaims: jwt.StandardClaims{
			Id:        tokenID.String(),
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(refreshKey))
	if err != nil {
		log.Println("error in signing refresh token:", err)
		return "", err
	}
	log.Println("token generated...", ss)
	return ss, nil

}
func (us *userService) verifyResetToken(ctx context.Context, tokenStr, key string) (*resetTokenClaims, error) {
	claims := &resetTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		log.Println("error while parsing token reason:", err)
		return nil, err
	}
	if !token.Valid {
		log.Println("inalid token provided:", err)
		return nil, fmt.Errorf("JWT token is invalid")
	}
	claims, ok := token.Claims.(*resetTokenClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing the claims from token")
	}
	return claims, nil
}

func (us *userService) GetCompany(ctx context.Context, key string) (*models.UserAccount, error) {
	user, err := us.UserRepository.GetUserById(ctx, key)
	if err != nil {
		log.Println("error in getting countries from DB reason:", err)
		return user, err
	}
	return user, nil
}

func (us *userService) GetMasters(ctx context.Context) (*models.MasterResult, error) {
	masterResult := &models.MasterResult{}
	countries, err := us.UserRepository.GetCountries(ctx)
	if err != nil {
		log.Println("error in getting countries from DB reason:", err)
		return nil, err
	}

	masterResult.Countries = countries
	// masterResult.TimeZones = timezones
	// log.Println("countries::::::", countries)
	return masterResult, nil
}
func (us *userService) GetTimezonesByCountry(ctx context.Context, country string) ([]*models.TimeZones, error) {
	timezones, err := us.UserRepository.GetTimezone(ctx, country)
	if err != nil {
		log.Println("error in getting timezones from DB reason:", err)
		return nil, err
	}
	for _, tz := range timezones {
		gmtOffsetSeconds, _ := time.ParseDuration(tz.GMT + "s")
		offset := strings.Trim(gmtOffsetSeconds.String(), "0s")
		// Calculate the hour offset by dividing the GMT offset in seconds by 3600
		tz.GMT = offset

	}
	return timezones, nil
}
func (us *userService) GetUserByID(ctx context.Context, id string) (*models.UserAccount, error) {
	result, err := us.UserRepository.GetUserById(ctx, id)
	if err != nil {
		log.Println("error in getting user from DB reason:", err)
		return nil, err
	}

	return result, nil
}
func (us *userService) GetServerByID(ctx context.Context, id string) (models.ServerDetails, error) {
	result, err := us.UserRepository.GetServerDetailsByIP(id)
	if err != nil {
		log.Println("error in getting user from DB reason:", err)
		return models.ServerDetails{}, err
	}
	return result, nil
}
func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
