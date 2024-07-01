package handlers

import (
	"net/http"
	"time"

	"github.com/bentenison/fastq-server/api/handlers/middleware"
	"github.com/bentenison/fastq-server/api/reports"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService    models.UserService
	TokenService   models.TokenService
	AdminService   models.AdminService
	TicketService  models.TicketService
	LicenseService models.LicenseService
	ConfigService  models.ConfigService
	ReportService  *reports.ReportsDAO
	MaxBodyBytes   int64
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	UserService     models.UserService
	TokenService    models.TokenService
	AdminService    models.AdminService
	TicketService   models.TicketService
	LicenseService  models.LicenseService
	ConfigService   models.ConfigService
	ReportService   *reports.ReportsDAO
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:    c.UserService,
		TokenService:   c.TokenService,
		AdminService:   c.AdminService,
		TicketService:  c.TicketService,
		LicenseService: c.LicenseService,
		ConfigService:  c.ConfigService,
		ReportService:  c.ReportService,
		MaxBodyBytes:   c.MaxBodyBytes,
	} // currently has no properties

	// Create an account group
	g := c.R.Group(c.BaseURL)
	go handleMessage()
	// if gin.Mode() != gin.TestMode {
	// 	g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
	// 	g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
	// 	g.PUT("/details", middleware.AuthUser(h.TokenService), h.Details)
	// 	g.POST("/image", middleware.AuthUser(h.TokenService), h.Image)
	// 	g.DELETE("/image", middleware.AuthUser(h.TokenService), h.DeleteImage)
	// } else {
	g.GET("/me", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "running",
			"time":   time.Now(),
		})
	})
	// 	g.POST("/signout", h.Signout)
	// 	g.PUT("/details", h.Details)
	// 	g.POST("/image", h.Image)
	// 	g.DELETE("/image", h.DeleteImage)
	// }
	g.StaticFS("/uploaded", http.Dir("uploads"))
	// g.StaticFS("/", http.Dir("/"))
	g.POST("/signup", h.signupHandler)
	g.POST("/signin", h.signinHandler)
	g.POST("/branch-login", h.branchLoginHandler)
	g.POST("/forgot-password", h.forgotPasswordHandler)
	g.POST("/token/refresh", h.TokensHandler)
	g.POST("/reset-password", h.signinHandler)
	g.GET("/ws", h.wsHandler)
	g.GET("/getconnectedclients", h.getConnectedClients)
	g.GET("/get-countries", h.getCountriesHandler)
	g.GET("/get-tz/:code", h.getTimezonesHandler)
	g.POST("/upload", h.uploadVideoHandlerfunc)
	//? branch handlers
	g.POST("/branch/addbranch", middleware.AuthUser(h.TokenService), h.addBranchHandler)
	g.POST("/branch/getbranch", middleware.AuthUser(h.TokenService), h.getBranchHandler)
	g.POST("/branch/updatebranch", middleware.AuthUser(h.TokenService), h.updateBranchHandler)
	g.GET("/branch/deletebranch/:id", middleware.AuthUser(h.TokenService), h.deleteBranchHandler)
	g.GET("/branch/getAllBranch/:code", h.getAllBranchHandler)
	//? service handlers
	g.POST("/service/addservice", middleware.AuthUser(h.TokenService), h.AddServiceHandler)
	g.POST("/service/getservice", middleware.AuthUser(h.TokenService), h.GetServiceHandler)
	g.POST("/service/updateservice", middleware.AuthUser(h.TokenService), h.UpdateServiceHandler)
	g.GET("/service/deleteservice/:id", middleware.AuthUser(h.TokenService), h.DeleteServiceHandler)
	g.GET("/service/getAllservice/:code", h.GetAllServicesHandler)
	g.GET("/service/assignService/:id", h.GetAssignedServices)
	g.GET("/service/getAllassignedServices", h.GetAllAssignedServices)
	g.POST("/service/updateassignedServices", h.UpdateCounteervicesHandler)
	g.POST("/service/addassignedServices", h.AddCountervicesHandler)
	g.GET("/service/deleteassignedServices/:id", h.DeleteCounterServicesHandler)
	//? section handlers
	g.POST("/section/addsection", middleware.AuthUser(h.TokenService), h.AddSectionHandler)
	g.POST("/section/getsection", middleware.AuthUser(h.TokenService), h.GetSectionHandler)
	g.POST("/section/updatesection", middleware.AuthUser(h.TokenService), h.UpdateSectionHandler)
	g.GET("/section/deletesection/:id", middleware.AuthUser(h.TokenService), h.DeleteSectionHandler)
	g.GET("/section/getAllsection/:code", h.UpdateSectionHandler)
	//? user handlers
	g.POST("/user/adduser", middleware.AuthUser(h.TokenService), h.AddUserHandler)
	g.POST("/user/getuser", middleware.AuthUser(h.TokenService), h.GetUserHandler)
	g.POST("/user/getAlluser", h.GetAllUsersHandler)
	g.POST("/user/updateuser", middleware.AuthUser(h.TokenService), h.UpdateUserHandler)
	g.POST("/user/updateemail", middleware.AuthUser(h.TokenService), h.UpdateUserEmailHandler)
	g.POST("/user/updatepassword", middleware.AuthUser(h.TokenService), h.UpdateUserPasswordHandler)
	g.GET("/user/deleteuser/:id", middleware.AuthUser(h.TokenService), h.DeleteUserHandler)
	g.GET("/user/resetlogin/:id", middleware.AuthUser(h.TokenService), h.ResetUserLoginsHandler)
	//? counter handlers
	g.POST("/counter/addcounter", middleware.AuthUser(h.TokenService), h.AddCounterHandler)
	g.POST("/counter/getcounter", h.GetCounterHandler)
	g.POST("/counter/updatecounter", h.UpdateCounterHandler)
	g.POST("/counter/interchangeCounters", h.InterchangeCounterHandler)
	g.GET("/counter/getAllcounter/:code", h.GetAllCountersHandler)
	g.GET("/counter/getinactivecounter/:code", h.GetAllUnActivatedCountersHandler)
	g.GET("/counter/deletecounter/:id", middleware.AuthUser(h.TokenService), h.DeleteCounterHandler)
	// g.POST("/signup", h.Signup)
	// g.POST("/signin", h.Signin)
	// g.POST("/tokens", h.Tokens)
	//?ticket handlers
	g.POST("/ticket/addticket", h.AddTicketHandler)
	g.POST("/ticket/deleteticket/:id", h.DeleteTicketHandler)
	g.POST("/ticket/getallticketsforday", h.GetAllTicketsForDayHandler)
	g.POST("/ticket/getticket", h.GetTicketHandler)
	g.POST("/ticket/getticketbytransfer", h.GetTicketsByTransferHandler)
	g.POST("/ticket/getticketbybranch", h.GetTicketsByBranchHandler)
	g.POST("/ticket/getticketbystatus", h.GetTicketsByStatusHandler)
	g.POST("/ticket/getticketbyuser", h.GetTicketsByUserHandler)
	g.POST("/ticket/updatetickettime", h.UpdateTicketStartTimeHandler)
	g.POST("/ticket/updateticketstatus", h.UpdateTicketStatusHandler)
	g.POST("/ticket/updatetickettransferto", h.UpdateTicketTransferedToHandler)
	g.POST("/ticket/updateticketuser", h.UpdateTicketUserHandler)
	g.GET("/ticket/getLastNumber/:service", h.GetTicketNumberHandler)
	g.POST("/ticket/updateendtime", h.UpdateTicketEndTimeHandler)
	g.POST("/ticket/gettickettoprocess", h.GetTicketToProcessHandler)
	g.GET("/ticket/getwaiting", h.GetWaitingTicketsHandler)
	g.GET("/ticket/getnoshow/:id", h.GetNoShowTicketsHandler)
	g.GET("/ticket/getfinished/:id", h.GetFinishedTicketsHandler)
	g.GET("/ticket/getlastcalled", h.GetLastCalledTicketHandler)
	g.GET("/ticket/getlaststarted/:id", h.GetLastStartedTicketHandler)
	g.GET("/ticket/getestimated/:id", h.GetEstimatedWaitingTimeHandler)

	//?license handlers
	// g.POST("/license/updateBranch", middleware.AuthUser(h.TokenService), h.UpdateLicenseBranchHandler)
	// g.POST("/license/updateCounters", h.UpdateLicenseCounterHandler)
	g.POST("/license/counteractivate", h.ActivateLicenseHandler)
	g.POST("/license/UpdateLicense", h.UpdateLicenseHandler)
	g.GET("/license/updatecounteruser/:id", h.UpdateCounterUserHandler)
	g.GET("/license/updateFirmName/:id", h.UpdateLicenseFirm)
	// g.GET("/license//:id", h.UpdateCounterUserHandler)
	// g.GET("/license/updatecounteruser/:id", h.UpdateCounterUserHandler)
	g.POST("/app/checkActivation", h.CheckCounterActivationHandler)
	g.GET("/app/checkLicense", h.CheckLicensehandler)
	g.POST("/app/checkandgetcounter", h.CheckAndGetActiveCounterHandler)
	g.POST("/counter/auth", h.AuthCounterHandler)
	// g.POST("/counter/logout/:id", h.AuthCounterHandler)
	// g.POST("/license/updateNumberLicense", middleware.AuthUser(h.TokenService))
	//misc
	g.GET("/app/getCompany/:id", h.GetCompanyFromApp)
	g.GET("/app/getServer/:id", h.GetServerFromApp)
	g.POST("/app/updatesystemip", h.UpdateSystemIPHandler)
	g.GET("/app/getAllClients", h.GetAllClientsRoute)

	g.GET("/config/getvideo/:id", h.GetVideoConfigHandler)
	g.GET("/config/getallvideo/:id", h.GetAllVideoHandler)
	g.POST("/config/setvideo", h.AddVideoHandler)
	g.GET("/config/getallschedularconfig/:id", h.GetSchedularConfigHandler)
	g.POST("/config/setschedularconfig", h.SetSchedularConfigHandler)
	// g.POST("/config/setformconfig")
	g.POST("/config/setdsconfig", h.SetDSConfigHandler)
	g.GET("/config/getdsconfig/:id", h.GetDSConfigHandler)
	g.POST("/config/updateds", h.UpdateDSConfigHandler)
	g.POST("/config/addannounce", h.AddAnnouncementHandler)
	g.POST("/config/updateannounce", h.UpdateAnnouncementHandler)
	g.GET("/config/getannounce/:id", h.GetAnnouncementHandler)
	g.GET("/config/getallannounce/:id", h.GetAllAnnouncementHandler)
	g.GET("/config/selecttodisplay/:id", h.SelectOneAnnouncement)
	g.POST("/config/addaudioconf", h.AddAudioHandler)
	g.POST("/config/updateaudioconf", h.UpdateAudioHandler)
	g.GET("/config/getaudioconf/:id", h.GetAudioConfHandler)
	g.POST("/config/addticketconf", h.AddTicketConfHandler)
	g.POST("/config/updateticketconf", h.UpdateTicketConfHandler)
	g.GET("/config/getticketconf/:id", h.GetTicketConfHandler)
	g.GET("/config/getallconfig/:id", h.GetAllConfigHandler)
	g.GET("/config/getServer/:id", h.GetServerByIDHandler)
	g.POST("/config/updateIp", h.UpdateServerIPByCodeHandler)
	g.GET("/config/deleteuploaded/:id", h.DeleteUploadedHandler)

	// Reports Handlers
	g.GET("/report/tickets-by-service/:id", h.GetTicketsByService)
	g.GET("/report/allwaitingTickets/:id", h.GetAllWaitingTickets)
	g.GET("/report/tickets-by-status/:id", h.GetTicketsByStatus)
	g.GET("/report/avg-service-time-by-service/:id", h.GetAvgServiceTimeByService)
	g.GET("/report/tickets-transferred-by-counter/:id", h.GetTicketsTransferredByCounter)
	g.GET("/report/tickets-by-company-branch/:id", h.GetTicketsByCompanyBranch)
	g.GET("/report/avg-waiting-time-by-service/:id", h.GetAvgWaitingTimeByService)
	g.GET("/report/tickets-served-by-counter/:id", h.GetTicketsServedByCounter)
	g.GET("/report/percentage-tickets-finished-by-service/:id", h.GetPercentageTicketsFinishedByService)
	g.POST("/report/tickets-between-dates-by-status-and-user", h.GetTicketsBetweenDatesByStatusAndUser)

	g.GET("/report/tickets-with-user-info/:id", h.GetTicketsWithUserInfo)
	g.GET("/report/tickets-with-counter-info/:id", h.GetTicketsWithCounterInfo)
	g.GET("/report/tickets-with-user-counter-info/:id", h.GetTicketsWithUserCounterInfo)
	g.GET("/report/tickets-served-by-counter-with-user-info/:id", h.GetTicketsServedByCounterWithUserInfo)
	g.GET("/report/tickets-transferred-with-user-info/:id", h.GetTicketsTransferredWithUserInfo)
	g.GET("/report/tickets-by-company-branch-counter-info/:id", h.GetTicketsByCompanyBranchCounterInfo)
	g.GET("/report/tickets-with-avg-service-time-by-counter/:id", h.GetTicketsWithAvgServiceTimeByCounter)
	g.GET("/report/average-active-time-of-user/:id", h.GetAverageActiveTimeOfUser)
	g.GET("/report/active-time-of-user-per-day/:id", h.GetActiveTimeOfUserPerDay)
	g.GET("/report/ticket-stats/:id", h.GetCreatedfinishedLast8Days)
	g.GET("/report/ticket-today-created/:id", h.GetHourlyCreated)
	g.GET("/report/ticket-today-finished/:id", h.GetHourlyFinished)
	g.GET("/report/ticket-today-noshow/:id", h.GetHourlyNoShow)
	///Client Customizations from here
	g.GET("/customize/changelanguage/:id", h.ChangeLanguageHandler)
	g.GET("/customize/changetemplate/:id", h.ChangeTemplateHandler)
	g.GET("/customize/getcustomizations", h.GetCustomizationHandler)
}
