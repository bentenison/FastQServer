package handlers

import (
	"log"
	"net/http"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/api/services"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ChangeLanguageHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err := services.ChangeTTSLanguage(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "OK")

}
func (h *Handler) ChangeTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err := services.ChangeTemplate(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "OK")

}
func (h *Handler) GetCustomizationHandler(c *gin.Context) {

	res, err := services.GetClientCustomozations()
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
