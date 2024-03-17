package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTicketHandler(c *gin.Context) {
	req := models.AddTicketParams{}
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
	ctx := c.Request.Context()
	res, err := h.TicketService.AddTicketService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in adding ticket to DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) DeleteTicketHandler(c *gin.Context) {
	// req := models.GetAllTicketsForDayParams{}
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.DeleteTicketService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in deleting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetAllTicketsForDayHandler(c *gin.Context) {
	req := models.GetAllTicketsForDayParams{}
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
	ctx := c.Request.Context()
	res, err := h.TicketService.GetAllTicketsForDayService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h *Handler) GetTicketHandler(c *gin.Context) {
	req := models.GetTicketParams{}
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
	ctx := c.Request.Context()
	res, err := h.TicketService.GetTicketService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetTicketToProcessHandler(c *gin.Context) {
	req := models.GetTicketParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.GetTicketToProcessService(ctx, req)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, "NO_RESULT")
		return
	}
	if err != nil {
		log.Println("error getting data :", err)
		// err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetTicketsByBranchHandler(c *gin.Context) {
	req := models.GetTicketsByBranchParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.GetTicketsByBranchService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetTicketsByStatusHandler(c *gin.Context) {
	req := models.GetTicketsByStatusParams{}
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
	ctx := c.Request.Context()
	res, err := h.TicketService.GetTicketsByStatusService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket by status from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetTicketsByTransferHandler(c *gin.Context) {
	req := models.GetTicketsByTransferParams{}
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
	ctx := c.Request.Context()
	res, err := h.TicketService.GetTicketsByTransferService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetTicketsByUserHandler(c *gin.Context) {
	req := models.GetTicketsByUserParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.GetTicketsByUserService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) UpdateTicketStartTimeHandler(c *gin.Context) {
	req := models.UpdateTicketStartTimeParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	err = h.TicketService.UpdateTicketStartTimeService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) UpdateTicketEndTimeHandler(c *gin.Context) {
	req := models.UpdateTicketEndTimeParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	err = h.TicketService.UpdateTicketEndTimeService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) UpdateTicketStatusHandler(c *gin.Context) {
	req := models.UpdateTicketStatusParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	err = h.TicketService.UpdateTicketStatusService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in updating ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
func (h *Handler) UpdateTicketTransferedToHandler(c *gin.Context) {
	req := models.UpdateTicketTransferedToParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	err = h.TicketService.UpdateTicketTransferedToService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
func (h *Handler) UpdateTicketUserHandler(c *gin.Context) {
	req := models.UpdateTicketUserParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	err = h.TicketService.UpdateTicketUserService(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in updating ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
func (h *Handler) GetTicketNumberHandler(c *gin.Context) {
	service := c.Param("service")
	if service == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	log.Println("handlers", h.TicketService)
	result, err := h.TicketService.GetLastTicketNumberService(ctx, service)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket number from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetWaitingTicketsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	log.Println("handlers", h.TicketService)
	result, err := h.TicketService.GetWaitingTicketsService(ctx)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket number from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetNoShowTicketsHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	log.Println("handlers", h.TicketService)
	result, err := h.TicketService.GetNoShowTicketsService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket number from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetFinishedTicketsHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	log.Println("handlers", h.TicketService)
	result, err := h.TicketService.GetFinishedTicketsService(ctx, id)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket number from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetLastCalledTicketHandler(c *gin.Context) {
	ctx := c.Request.Context()
	log.Println("handlers", h.TicketService)
	result, err := h.TicketService.GetLastCalledService(ctx)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, "NO_RESULT")
		return
	}
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting ticket number from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetLastStartedTicketHandler(c *gin.Context) {
	// req := models.GetAllTicketsForDayParams{}
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.GetLastStartedTicket(ctx, id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, "NO_RESULT")
		return
	}
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in deleting ticket from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) UpdateSystemIPHandler(c *gin.Context) {
	req := models.UpdateSystemParams{}
	err := c.Bind(&req)
	log.Println("data::::::", req)
	if err != nil {
		log.Println("error binding data :", err)
		err := apperrors.NewExpectationFailed("invalid data provided")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	ctx := c.Request.Context()
	res, err := h.TicketService.UpdateSystemIp(ctx, req)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in updating system ip from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetAllClientsRoute(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := h.TicketService.GetAllSystemService(ctx)
	if err != nil {
		log.Println("error getting data :", err)
		err := apperrors.NewExpectationFailed("error in getting systems ip from DB")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// req := signUpReq{}
// 	err := c.Bind(&req)
// 	log.Println("data::::::", req)
// 	if err != nil {
// 		log.Println("error binding data :", err)
// 		err := apperrors.NewExpectationFailed("invalid data provided")
// 		c.JSON(err.Status(), gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	addr, _ := mail.ParseAddress(req.Email)
// 	if addr == nil {
// 		log.Println("error in parsing email address:", err)
// 		err := apperrors.NewExpectationFailed(fmt.Sprintf("invalid email address %v provided", req.Email))
// 		c.JSON(err.Status(), gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	ctx := c.Request.Context()
// 	user := models.UserAccount{
// 		Email:     req.Email,
// 		Password:  req.Password,
// 		FirstName: req.FirstName,
// 		LastName:  req.LastName,
// 		Company:   req.Company,
// 		Country:   req.Country,
// 		Timezone:  req.Timezone,
// 	}
// 	err = h.UserService.SignUp(ctx, &user)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, apperrors.UserAlreadyExists):
// 			log.Println("error in signup service reason:", err)
// 			err := apperrors.NewExpectationFailed("User already exist. Kindly login")
// 			c.JSON(err.Status(), gin.H{
// 				"error": err,
// 			})
// 			return
// 		default:
// 			log.Println("error in signing up user:", err)
// 			err := apperrors.NewInternal()
// 			c.JSON(err.Status(), gin.H{
// 				"error": err,
// 			})
// 			return
// 		}
// 	}
