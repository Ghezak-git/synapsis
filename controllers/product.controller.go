package controllers

import (
	"net/http"
	"strconv"

	"github.com/Ghezak-git/synapsis/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(DB *gorm.DB) ProductController {
	return ProductController{DB}
}

func (pc *ProductController) GetProduct(ctx *gin.Context) {
	var payload *models.GetProductsRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fails", "message": err.Error()})
		return
	}

	var pageInt int = 0
	var perPageInt int = 10

	if payload.Page != "" {
		pageInt, _ = strconv.Atoi(payload.Page)
	}

	if payload.PerPage != "" {
		perPageInt, _ = strconv.Atoi(payload.PerPage)
	}

	var product []models.Product
	var count int64

	results := pc.DB.Preload("Category")
	resultscount := pc.DB.Model(&product)

	if payload.Search != "" {
		resultscount.Where("name LIKE ?", "%"+payload.Search+"%")
		results.Where("name LIKE ?", "%"+payload.Search+"%")
	}

	if payload.Category != "" {
		resultscount.Where("id_category = ?", payload.Category)
		results.Where("id_category = ?", payload.Category)
	}

	if payload.Sort != "" && payload.SortColumn != "" {
		results.Order(payload.SortColumn + " " + payload.Sort)
	} else {
		results.Order("updated_at desc")
	}

	if payload.Page != "" && pageInt > 1 {
		noskip := 0
		notake := perPageInt
		noskip = (notake * pageInt) - notake
		results.Limit(noskip).Offset(notake)
	} else {
		results.Limit(perPageInt)
	}

	resultscount.Count(&count)
	results.Find(&product)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	if resultscount.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "count": count, "data": product})
}

func (pc *ProductController) PostCart(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.PostCartRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var shopCart models.ShopCart
	var results *gorm.DB
	result := pc.DB.Where("id_product = ? AND id_user = ?", payload.IDProduct, currentUser.ID).First(&shopCart)

	if result.RowsAffected > 0 {
		shopCartToUpdate := models.ShopCart{
			IDProduct: payload.IDProduct,
			IDUser:    currentUser.ID,
			QTY:       shopCart.QTY + 1,
		}

		results = pc.DB.Model(&shopCart).Updates(shopCartToUpdate)
	} else {
		newPost := models.ShopCart{
			IDProduct: payload.IDProduct,
			IDUser:    currentUser.ID,
			QTY:       1,
		}

		results = pc.DB.Create(&newPost)
	}

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": "Success Add Product to Cart"})
}

func (pc *ProductController) GetCart(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var shopCart []models.ShopCart
	result := pc.DB.Preload("Product.Category").Preload("User").Where("id_user = ?", currentUser.ID).Find(&shopCart)

	if result.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "false", "data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": shopCart})
}

func (pc *ProductController) UpdateCart(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.UpdateCartRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var shopCart models.ShopCart
	var results *gorm.DB
	result := pc.DB.Where("id_product = ? AND id_user = ?", payload.IDProduct, currentUser.ID).First(&shopCart)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	if payload.Method == "update" {
		cardToUpdate := models.ShopCart{
			IDProduct: payload.IDProduct,
			IDUser:    currentUser.ID,
			QTY:       payload.QTY,
		}

		results = pc.DB.Model(&shopCart).Updates(cardToUpdate)
	} else {
		results = pc.DB.Delete(&models.ShopCart{}, "id_product = ? AND id_user = ?", payload.IDProduct, currentUser.ID)
	}

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": "Success Update Cart"})
}
