package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ActivateLicenseHandler(c *gin.Context) {
	var activeCounter models.ActiveCounter
	err := c.Bind(&activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	active, err := h.LicenseService.ActivateCounterService(activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, active)

}
func (h *Handler) CheckAndGetActiveCounterHandler(c *gin.Context) {
	var activeCounter models.ActiveCounter
	err := c.Bind(&activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	active, err := h.LicenseService.CheckAndGetActiveCounter(activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, active)

}

func (h *Handler) CheckCounterActivationHandler(c *gin.Context) {
	var activeCounter models.ActiveCounter
	err := c.Bind(&activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	active, err := h.LicenseService.CheckCounterActivationService(activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, active)

}
func (h *Handler) AuthCounterHandler(c *gin.Context) {
	var user models.AuthCounterUser
	err := c.Bind(&user)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	if user.Email == "" || user.Password == "" {
		// log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	counterUser, err := h.LicenseService.AuthCounterUserService(user)
	if err != nil {
		if strings.Contains(err.Error(), "user is already logged") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			log.Println("error binding data :", err)
			err := apperrors.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, counterUser)

}
func (h *Handler) CounterLogoutHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	user := models.ManageUser{}
	err := h.LicenseService.CounterLogoutService(id, user)
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
func (h *Handler) UpdateLicenseHandler(c *gin.Context) {
	var licensePayload models.LicensePayload
	err := c.Bind(&licensePayload)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	active, err := h.LicenseService.UpdateLicenseService(licensePayload)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, active)

}

func (h *Handler) UpdateLicenseCounterHandler(c *gin.Context) {
	var activeCounter models.ActiveCounter
	err := c.Bind(&activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	active, err := h.LicenseService.UpdateCounterDetails(activeCounter)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, active)

}
func (h *Handler) CheckLicensehandler(c *gin.Context) {
	active, err := h.LicenseService.CheckLicensePresent()
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, active)
}

func (h *Handler) GetCompanyFromApp(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	user, err := h.UserService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) GetServerFromApp(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	user, err := h.UserService.GetServerByID(c.Request.Context(), id)
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateCounterUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	err := h.LicenseService.UpdateCounterUserService(id)
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) GetAssignedServices(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	res, err := h.ConfigService.GetServicesByCounterIDService(id)
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetAllAssignedServices(c *gin.Context) {
	res, err := h.ConfigService.GetAllAssignedServices()
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) UpdateCounteervicesHandler(c *gin.Context) {
	var counterService models.CounterService
	err := c.Bind(&counterService)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	err = h.ConfigService.UpdateAssignedServices(counterService)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) AddCountervicesHandler(c *gin.Context) {
	var counterService models.CounterService
	err := c.Bind(&counterService)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	err = h.ConfigService.AssignCounterServices(counterService)
	if err != nil {
		log.Println("error assigning services :", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) UpdateLicenseFirm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	res, err := h.LicenseService.UpdateCompanyNameInLicense(id)
	if err != nil {
		log.Println("error occured while getting user by id", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) CheckLicenseFirm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	res, err := h.LicenseService.CheckFirmNameInLicense(id)
	if err != nil {
		log.Println("error occured while chcking firm in license", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
