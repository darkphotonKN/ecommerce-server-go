package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/darkphotonKN/ecommerce-server-go/internal/rating"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

var validate = validator.New()

func (h *ProductHandler) GetProductsHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("page", "10"))
	pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))

	offset := (pageNumber - 1) * limit
	products, err := h.Service.GetProductsService(limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to get all products: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Successfully retrieved all products.",
		"result":     products,
	})
}

func (h *ProductHandler) GetTrendingProductsHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))

	products, err := h.Service.GetProductsService(limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to get all trending products: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Successfully retrieved all trending products.",
		"result":     products,
	})
}

func (h *ProductHandler) GetProductByIdHandler(c *gin.Context) {

	idParams := c.Param("id")

	id, err := uuid.Parse(idParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to parse id as UUID: %s", err.Error())})
		return
	}

	product, err := h.Service.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to get product: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Successfully retrieved product.",
		"result":     product,
	})

}

func (h *ProductHandler) CreateProductsHandler(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error parsing json from the payload: %s", err.Error())})
		return
	}

	// validation
	err = validate.Struct(product)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "Validation failed",
			"errors":     validationErrors.Error(),
		})
		return
	}

	err = h.Service.CreateProductService(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to create new product: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Successfully created new product.",
	})
}

func (h *ProductHandler) CreateProductRating(c *gin.Context) {
	var ratingReq rating.RatingRequest
	err := c.ShouldBindJSON(&ratingReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error parsing json from the payload: %s", err.Error())})
		return
	}

	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to parse id as UUID: %s", err.Error())})
		return
	}

	err = h.Service.PostRatingService(id, ratingReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to create new product: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    fmt.Sprintf("Successfully created rating product with id: %s", id)})
}
