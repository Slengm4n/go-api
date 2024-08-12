package main

import (
	"go-ai/controller"
	"go-ai/db"
	"go-ai/repository"
	"go-ai/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//Conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	productRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	productUseCase := usecase.NewProductUseCase(productRepository)
	//Camada de controllers
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("product/:productId", productController.GetProductById)
	server.Run(":8000")
}
