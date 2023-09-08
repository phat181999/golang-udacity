package routes

import (
	"log"
	"net/http"

	controller "github/phat/product/controller"
	database "github/phat/product/database"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	s := database.Server{};
	s.Init();
	crt := &controller.APIEnv{
		DB: s.DB,
	};
	log.Println("Routes");
	httpRouter := gin.Default();
	// Add template static
	httpRouter.LoadHTMLGlob("template/index.html");
	httpRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	api := httpRouter.Group("/api")
	{
		user := api.Group("/customer")
		{
			user.POST("/create", crt.Register)
			user.GET("/get-customers", crt.GetCustomers)
			user.GET("/get-customer/:id", crt.GetDetailCustomer)
			user.DELETE("/delete-customer/:id", crt.DeleteMovie)
			user.PUT("/update-customer/:id", crt.UpdateMovie)
		}
	};

	// Add a catch-all route for non-existent endpoints
	httpRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Endpoint not found 1",
		})
	});

	
	return httpRouter
}
