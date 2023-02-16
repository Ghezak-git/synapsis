package routes

import (
	"github.com/Ghezak-git/synapsis/controllers"
	"github.com/Ghezak-git/synapsis/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewProductRouteController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {

	router := rg.Group("product")
	router.Use(middleware.DeserializeUser())
	router.POST("/listproducts", pc.productController.GetProduct)
	router.POST("/addcart", pc.productController.PostCart)
	router.GET("/listcart", pc.productController.GetCart)
	router.POST("/updatecart", pc.productController.UpdateCart)

}
