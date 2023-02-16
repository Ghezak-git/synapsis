package routes

import (
	"github.com/Ghezak-git/synapsis/controllers"
	"github.com/Ghezak-git/synapsis/middleware"
	"github.com/gin-gonic/gin"
)

type TransactionRouteController struct {
	transactionController controllers.TransactionController
}

func NewTransactionRouteController(trasactionController controllers.TransactionController) TransactionRouteController {
	return TransactionRouteController{trasactionController}
}

func (pc *TransactionRouteController) TransactionRoute(rg *gin.RouterGroup) {

	router := rg.Group("transaction")
	router.Use(middleware.DeserializeUser())
	router.POST("/listtransaction", pc.transactionController.GetTransaction)
	router.GET("/listpaymentmethod", pc.transactionController.GetPaymentMethod)
	router.POST("/addtransaction", pc.transactionController.PostTransaction)

}
