package handlers

import (
	"log"
	"net/http"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTicketsByService(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTotalTicketsByService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}
func (h *Handler) GetAllWaitingTickets(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetAllWaitingTickets(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsByStatus handles the route for Query 2.
func (h *Handler) GetTicketsByStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsByStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// Continue with other route handler functions...
// ... (Previous code)

// GetAvgServiceTimeByService handles the route for Query 3.
func (h *Handler) GetAvgServiceTimeByService(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetAverageServiceTimeByService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}

// GetTicketsTransferredByCounter handles the route for Query 4.
func (h *Handler) GetTicketsTransferredByCounter(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	transferred, err := h.ReportService.GetTicketsTransferredByCounter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, transferred)
}

// GetTicketsByCompanyBranch handles the route for Query 5.
func (h *Handler) GetTicketsByCompanyBranch(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsByCompanyBranch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetAvgWaitingTimeByService handles the route for Query 6.
func (h *Handler) GetAvgWaitingTimeByService(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetAverageServiceTimeByService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}

// GetTicketsServedByCounter handles the route for Query 7.
func (h *Handler) GetTicketsServedByCounter(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsServedByCounter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetPercentageTicketsFinishedByService handles the route for Query 8.
func (h *Handler) GetPercentageTicketsFinishedByService(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	percentages, err := h.ReportService.GetPercentageTicketsFinishedByService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, percentages)
}

type TicketsBetweenDates struct {
	Code      string `json:"code,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// GetTicketsBetweenDatesByStatusAndUser handles the route for Query 9.
func (h *Handler) GetTicketsBetweenDatesByStatusAndUser(c *gin.Context) {
	req := TicketsBetweenDates{}
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
	tickets, err := h.ReportService.GetTicketsBetweenDatesByStatusUser(req.Code, req.StartDate, req.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsWithUserInfo handles the route for Query 10.
func (h *Handler) GetTicketsWithUserInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsWithUserInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsWithCounterInfo handles the route for Query 11.
func (h *Handler) GetTicketsWithCounterInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsWithCounterInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsWithUserCounterInfo handles the route for Query 12.
func (h *Handler) GetTicketsWithUserCounterInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsWithUserCounterInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsServedByCounterWithUserInfo handles the route for Query 13.
func (h *Handler) GetTicketsServedByCounterWithUserInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsServedByCounterWithUserInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsTransferredWithUserInfo handles the route for Query 14.
func (h *Handler) GetTicketsTransferredWithUserInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsTransferredWithUserInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsByCompanyBranchCounterInfo handles the route for Query 15.
func (h *Handler) GetTicketsByCompanyBranchCounterInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetTicketsByCompanyBranchCounter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsWithAvgServiceTimeByCounter handles the route for Query 16.
func (h *Handler) GetTicketsWithAvgServiceTimeByCounter(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	tickets, err := h.ReportService.GetAverageServiceTimeByCounter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetAverageActiveTimeOfUser handles the route for Query 17.
func (h *Handler) GetAverageActiveTimeOfUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetAverageActiveTimeOfUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}

// GetActiveTimeOfUserPerDay handles the route for Query 18.
func (h *Handler) GetActiveTimeOfUserPerDay(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetActiveTimeOfUserPerDay(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}
func (h *Handler) GetCreatedfinishedLast8Days(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetTicketStatsForLast8Days(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}
func (h *Handler) GetHourlyCreated(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetHourlyTicketStatsForToday(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}
func (h *Handler) GetHourlyFinished(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetHourlyTicketFinishedForToday(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}
func (h *Handler) GetHourlyNoShow(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.GetHourlyTicketNoShowForToday(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}
func (h *Handler) GetEstimatedWaitingTimeHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
		return
	}
	times, err := h.ReportService.EstimatedWaitingTime(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, times)
}
