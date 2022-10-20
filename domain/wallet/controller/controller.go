package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"wallet-manager/domain/wallet/service"
	"wallet-manager/infrastructure/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IWalletController interface {
	Create(c *gin.Context)
	AddFund(c *gin.Context)
	SubtractFund(c *gin.Context)
	GetByID(c *gin.Context)
	GetByUserID(c *gin.Context)
}

type WalletController struct {
	service service.IWalletService
	store   database.IDatabase
}

func NewWalletController(service service.IWalletService, store database.IDatabase) *WalletController {
	return &WalletController{service: service, store: store}
}

// Run - Starts the gin engine and sets up the http routes
func (w *WalletController) Run(port string) *http.Server {
	// init gin
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	router.GET("/health", w.HealthCheck)

	v1 := router.Group("/v1")
	{
		wallet := v1.Group("/wallet")
		{
			wallet.POST("/create", w.Create)
			wallet.POST("/add-fund", w.AddFund)
			wallet.POST("/subtract-fund", w.SubtractFund)
			wallet.GET("/:id", w.GetByID)
			wallet.GET("/user/:id", w.GetByUserID)
		}
	}

	// gin middleware config
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": gin.H{"status": false, "message": fmt.Sprintf("Page not found: %s, method: %s", c.Request.URL, c.Request.Method)}})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": gin.H{"status": false, "message": "Method not found"}})
	})

	// Note: we use http server to have graceful shutdown
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		log.Printf("Listening and serving HTTP on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("gin sever stoped with err: %s \n", err)
		}
	}()

	return server
}

func (w *WalletController) Create(c *gin.Context) {
	var request createRequest
	if err := c.BindJSON(&request); err != nil {
		w.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := w.service.Create(c.Request.Context(), request.UserID)
	if err != nil {
		w.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	w.ginResponse(c, http.StatusOK, wallet)
}

func (w *WalletController) AddFund(c *gin.Context) {
	var request addFundRequest
	if err := c.BindJSON(&request); err != nil {
		w.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := w.service.AddFund(c.Request.Context(), request.ID, request.Fund)
	if err != nil {
		w.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	w.ginResponse(c, http.StatusOK, wallet)
}

func (w *WalletController) SubtractFund(c *gin.Context) {
	var request subtractFundRequest
	if err := c.BindJSON(&request); err != nil {
		w.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := w.service.SubtractFund(c.Request.Context(), request.ID, request.Fund)
	if err != nil {
		w.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	w.ginResponse(c, http.StatusOK, wallet)
}

func (w *WalletController) GetByID(c *gin.Context) {
	ID := c.Param("id")
	IDInt, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		w.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := w.service.GetByID(c.Request.Context(), IDInt)
	if err != nil {
		w.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	w.ginResponse(c, http.StatusOK, wallet)
}

func (w *WalletController) GetByUserID(c *gin.Context) {
	userID := c.Param("id")
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		w.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := w.service.GetByUserID(c.Request.Context(), userIDInt)
	if err != nil {
		w.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	w.ginResponse(c, http.StatusOK, wallet)
}

// HealthCheck - Checks database health
func (w *WalletController) HealthCheck(c *gin.Context) {
	health := map[string]interface{}{
		"store": "up",
	}

	if err := w.store.Ping(); err != nil {
		health["database"] = "down"
		w.ginResponse(c, http.StatusInternalServerError, health)
		return
	}

	w.ginResponse(c, http.StatusOK, health)
}

func (w *WalletController) ginResponse(c *gin.Context, status int, payload interface{}) {
	type Response struct {
		Status  int         `json:"status"`
		Payload interface{} `json:"payload"`
	}

	response := Response{
		Status:  status,
		Payload: payload,
	}

	c.Header("Content-Type", "application/json")
	c.Status(status)

	c.JSON(status, response)
}
