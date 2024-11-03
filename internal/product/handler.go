package product

import (
	"fmt"
	"net/http"

	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (h *ProductHandler) GetProductsHandler(c *gin.Context) {
	products, err := h.Service.GetProductsService()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusInternalServerError, "message": fmt.Sprintf("Error when attempting to get all products: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message:":   "Successfully retrieved all products.",
		"result":     products,
	})
}

func (h *ProductHandler) GetTrendingProductsHandler(c *gin.Context) {
	products, err := h.Service.GetProductsService()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusInternalServerError, "message": fmt.Sprintf("Error when attempting to get all trending products: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message:":   "Successfully retrieved all trending products.",
		"result":     products,
	})
}

func (h *ProductHandler) CreateProductsHandler(c *gin.Context) {
	var product models.Product
	c.ShouldBindJSON(&product)

	err := h.Service.CreateProductService(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusInternalServerError, "message": fmt.Sprintf("Error when attempting to create new product: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message:":   "Successfully created new product.",
	})

}
