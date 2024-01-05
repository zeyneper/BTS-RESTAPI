package main

import (
	"fmt"
	"mentalartsapi/database"
	"mentalartsapi/docs"
	"mentalartsapi/handlers"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	r := gin.Default()

	db, err := database.InitDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	h := handlers.NewHandler(db)

	v1 := r.Group("/v1")
	{

		book := v1.Group("/book")
		{
			book.GET("", h.GetBooks)
			book.GET(":id", h.GetBookById)
			book.PUT(":id", h.UpdateBook)
			book.DELETE(":id", h.DeleteBookById)

		}
		user := v1.Group("/user")
		{

			user.GET("", h.GetUsers)
			user.POST("", h.CreateUser)
			user.GET(":id", h.GetUser)
			user.PUT(":id", h.UpdateUser)
			user.DELETE(":id", h.DeleteUserById)

		}
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Run() // listen and serve on 0.0.0.0:8080
	}

}
