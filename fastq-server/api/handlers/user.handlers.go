package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-gonic/gin"
)

type signUpReq struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Company         string `json:"company"`
	Country         string `json:"country"`
	Timezone        string `json:"timezone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (h *Handler) signupHandler(c *gin.Context) {
	req := signUpReq{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	addr, _ := mail.ParseAddress(req.Email)
	if addr == nil {
		log.Println("error in parsing email address:", err)
		err := apperrors.NewExpectationFailed(fmt.Sprintf("invalid email address %v provided", req.Email))
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx := c.Request.Context()
	user := models.UserAccount{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Company:   req.Company,
		Country:   req.Country,
		Timezone:  req.Timezone,
	}
	err = h.UserService.SignUp(ctx, &user)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.UserAlreadyExists):
			log.Println("error in signup service reason:", err)
			err := apperrors.NewExpectationFailed("User already exist. Kindly login")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		default:
			log.Println("error in signing up user:", err)
			err := apperrors.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup Successful. Kindly Login.",
	})
}
func (h *Handler) signinHandler(c *gin.Context) {
	req := signUpReq{}
	err := c.Bind(&req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	addr, _ := mail.ParseAddress(req.Email)
	if addr == nil {
		log.Println("error in parsing email address:", err)
		err := apperrors.NewExpectationFailed(fmt.Sprintf("invalid email address %v provided", req.Email))
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx := c.Request.Context()
	user := models.UserAccount{
		Email:    req.Email,
		Password: req.Password,
	}
	err = h.UserService.SignIn(ctx, &user)
	if err != nil {
		switch {
		case err == apperrors.PasswordMismatched:
			err := apperrors.NewExpectationFailed("Incorrect password. Enter correct password")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		case errors.Is(err, sql.ErrNoRows):
			err := apperrors.NewExpectationFailed("User does not exist. Kindly SignUp first")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		default:
			log.Println("error in signinHandler reason:", err)
			err := apperrors.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
	}

	tokens, err := h.TokenService.GeneratePairFromUser(ctx, &user, "")
	if err != nil {
		log.Println("error in creating user tokens:", err)
		err := apperrors.NewServiceUnavailable()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": tokens,
	})
}
func (h *Handler) branchLoginHandler(c *gin.Context) {
	req := signUpReq{}
	err := c.Bind(&req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	// addr, _ := mail.ParseAddress(req.Email)
	// if addr == nil {
	// 	log.Println("error in parsing email address:", err)
	// 	err := apperrors.NewExpectationFailed(fmt.Sprintf("invalid email address %v provided", req.Email))
	// 	c.JSON(err.Status(), gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// }

	ctx := c.Request.Context()
	user := models.UserAccount{
		Email:    req.Email,
		Password: req.Password,
	}
	err = h.UserService.BranchLogin(ctx, &user)
	if err != nil {
		switch {
		case err == apperrors.PasswordMismatched:
			err := apperrors.NewExpectationFailed("Incorrect password. Enter correct password")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		case errors.Is(err, sql.ErrNoRows):
			err := apperrors.NewExpectationFailed("User does not exist. Kindly SignUp first")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		default:
			log.Println("error in signinHandler reason:", err)
			err := apperrors.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
	}

	tokens, err := h.TokenService.GeneratePairFromUser(ctx, &user, "")
	if err != nil {
		log.Println("error in creating user tokens:", err)
		err := apperrors.NewServiceUnavailable()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": tokens,
	})
}

func (h *Handler) forgotPasswordHandler(c *gin.Context) {
	req := signUpReq{}
	err := c.Bind(&req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	err = h.UserService.ForgotPassword(c.Request.Context(), req.Email)
	if err != nil {
		log.Println("error in forgotpassword service :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "email sent.. check your inbox.",
	})
}
func (h *Handler) resetPasswordHandler(c *gin.Context) {
	req := signUpReq{}
	err := c.Bind(&req)
	token := c.Param("t")
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	h.UserService.ResetPassword(c.Request.Context(), req.Email, token)
}

func (h *Handler) getCountriesHandler(c *gin.Context) {
	data, err := h.UserService.GetMasters(c.Request.Context())
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func (h *Handler) getTimezonesHandler(c *gin.Context) {
	code := c.Param("code")
	data, err := h.UserService.GetTimezonesByCountry(c.Request.Context(), code)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

type tokensReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Tokens handler
func (h *Handler) TokensHandler(c *gin.Context) {
	// bind JSON to req of type tokensRew
	req := tokensReq{}

	err := c.Bind(&req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx := c.Request.Context()

	// verify refresh JWT
	refreshToken, err := h.TokenService.ValidateRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// get up-to-date user
	u, err := h.UserService.GetUserByID(ctx, refreshToken.UID)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create fresh pair of tokens
	tokens, err := h.TokenService.GeneratePairFromUser(ctx, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %+v. Error: %v\n", u, err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}
