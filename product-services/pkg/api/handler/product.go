package handler

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler/interfaces"
	"github.com/gin-gonic/gin"

	services "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/usecase/interfaces"
)

type ProductHandler struct {
	ProductUseCase services.ProductUseCase
}

func NewProductHandler(productUsecase services.ProductUseCase) handlerinterfaces.ProductHandler {
	return &ProductHandler{ProductUseCase: productUsecase}
}

// AddProducts godoc
// @summary api for admin to update a product
// @id AddProducts
// @tags Admin Products
// @Param input body req.ReqProduct{} true "inputs"
// @Router /products [post]
// @Success 200 {object} res.Response{} "successfully product added"
// @Failure 400 {object} res.Response{} "invalid input"
func (P *ProductHandler) AddProduct(ctx *gin.Context) {

}

func (P *ProductHandler) UpdateProduct(ctx *gin.Context) {

}

func (P *ProductHandler) DeteleProduct(ctx *gin.Context) {

}

func (P *ProductHandler) ListFullProducts(ctx *gin.Context) {

}

func (P *ProductHandler) GetProductById(ctx *gin.Context) {

}
