package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddVideoHandler(c *gin.Context) {
	req := models.Video{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.ConfigService.AddVideoService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetVideoConfigHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.ConfigService.GetVideoService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetAllVideoHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.ConfigService.GetAllVideoService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetSchedularConfigHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.ConfigService.GetAllSchedularConfService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) SetSchedularConfigHandler(c *gin.Context) {
	req := models.VideoSchedular{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.ConfigService.AddSchedularConfigService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) SetDSConfigHandler(c *gin.Context) {
	req := models.DSConfig{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.AddDSConfigService(req)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) UpdateDSConfigHandler(c *gin.Context) {
	req := models.DSConfig{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.UpdateDsConfigService(req)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) GetDSConfigHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetDSConfigService(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetAllConfigs(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.DeleteTicketService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in deleting ticket from DB")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetAnnouncementHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetAnnouncement(id)
	if err != nil {
		log.Println("error getting data :", err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetAllAnnouncementHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetAllAnnouncement(id)
	if err != nil {
		log.Println("error getting data :", err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) SelectOneAnnouncement(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	anId := c.Query("id")
	if anId == "" {
		err := apperrors.NewExpectationFailed("anid is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err := h.ConfigService.SelectAnnouncementToDisplay(anId, id)
	if err != nil {
		log.Println("error getting data :", err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) AddAnnouncementHandler(c *gin.Context) {
	req := models.AnnouncementConf{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.AddAnnouncementService(req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in adding announcement to DB")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}

type UpdateAnnouncementParams struct {
	Text        string `json:"text,omitempty"`
	Speed       int    `json:"speed,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
}

func (h *Handler) UpdateAnnouncementHandler(c *gin.Context) {
	req := UpdateAnnouncementParams{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.UpdateAnnouncement(req.CompanyCode, req.Speed, req.Text)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in updating announcement to DB")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}

func (h *Handler) AddAudioHandler(c *gin.Context) {
	req := models.AudioConfig{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.AddAudioConfigService(req)
	if err != nil {
		log.Println("error adding data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) UpdateAudioHandler(c *gin.Context) {
	req := models.AudioConfig{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.UpdateAudioConfigService(req)
	if err != nil {
		log.Println("error updating data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) GetAudioConfHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetAudionConfigService(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) AddTicketConfHandler(c *gin.Context) {
	req := models.TicketConf{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.AddTicketConf(req)
	if err != nil {
		log.Println("error adding data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) UpdateTicketConfHandler(c *gin.Context) {
	req := models.TicketConf{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.UpdateTicketConf(req.CompanyCode, req)
	if err != nil {
		log.Println("error adding data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")

}
func (h *Handler) GetTicketConfHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetTicketConf(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetAllConfigHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetAllConfService(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetServerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.GetServerByIDService(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) UpdateServerIPByCodeHandler(c *gin.Context) {
	req := models.ServerDetails{}
	err := c.Bind(&req)
	// log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err = h.ConfigService.UpdateIPByCode(req)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "OK")

}
func (h *Handler) DeleteCounterServicesHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	err := h.ConfigService.DeleteCounterServices(id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "OK")

}
func (h *Handler) DeleteUploadedHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	// ctx := c.Request.Context()
	res, err := h.ConfigService.DeleteVideoService(context.TODO(), id)
	if err != nil {
		log.Println("error getting data :", err)

		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
