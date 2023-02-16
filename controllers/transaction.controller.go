package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Ghezak-git/synapsis/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

func NewTransactionController(DB *gorm.DB) TransactionController {
	return TransactionController{DB}
}

func (tc *TransactionController) GetTransaction(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.GetProductsRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
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

	var transaction []models.TransactionHeader
	var count int64

	results := tc.DB.Preload("Details")
	resultscount := tc.DB.Model(&transaction)

	if payload.Search != "" {
		resultscount.Where("no_transaction LIKE ?", "%"+payload.Search+"%")
		results.Where("no_transaction LIKE ?", "%"+payload.Search+"%")
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

	results.Where("id_created = ?", currentUser.ID)

	resultscount.Count(&count)
	results.Find(&transaction)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	if resultscount.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "count": count, "data": transaction})
}

func (tc *TransactionController) GetPaymentMethod(ctx *gin.Context) {
	var paymentMethod []models.PaymentMethod
	result := tc.DB.Find(&paymentMethod)

	if result.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "false", "data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": paymentMethod})
}

func (tc *TransactionController) PostTransaction(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.TrasactionPostHeaderRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newPost := models.TransactionHeader{
		NamaCustomer:       currentUser.FullName,
		NoTransaction:      StringWithCharset(),
		PaymentMethod:      payload.PaymentMethod,
		PaymentDescription: payload.PaymentDescription,
		Total:              payload.Total,
		IDCreated:          currentUser.ID,
		IDUpdated:          currentUser.ID,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	result := tc.DB.Create(&newPost)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	for _, v := range payload.Detail {
		total := v.QTY * v.Price

		now := time.Now()
		newPostDetail := models.TransactioDetail{
			IDHeader:           newPost.ID,
			NamaProduct:        v.NamaProduct,
			DescriptionProduct: v.DescriptionProduct,
			Price:              v.Price,
			QTY:                v.QTY,
			Total:              total,
			IDCreated:          currentUser.ID,
			IDUpdated:          currentUser.ID,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		resultDetail := tc.DB.Create(&newPostDetail)
		if resultDetail.Error != nil {
			tc.DB.Delete(&models.TransactionHeader{}, "id_transaction_header = ? ", newPost.ID)
			tc.DB.Delete(&models.TransactioDetail{}, "id_transaction_header = ? ", newPost.ID)
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Transaction has been successfully"})
}
