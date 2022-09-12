package main

import (
	"log"
	"net/http"
	"os"

	"pheet-gin-backend/product/http"
	"pheet-gin-backend/product/repository"
	"pheet-gin-backend/product/service"
	"pheet-gin-backend/route"
	"pheet-gin-backend/auth"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func main() {
	var err error
	err = godotenv.Load("local.env") /*Load Env*/
	if err != nil {
		log.Printf("please consider environment variable: %s", err)
	}

	db, err := sqlx.Open("mysql", os.Getenv("DB_CONN"))
	if err != nil {
		panic(err)
	}
	
	productRepo := repository.NewProductRepo(db)
	productServ := service.NewProductService(productRepo)
	productHand := handler.NewProductHandler(productServ)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Bizcuitware Web!!!")
	})

	r.POST("/login", productHand.Login)

	routeGroup := r.Group("", auth.Protect([]byte(os.Getenv("SIGN"))))
	routeGroup.Use(cors.New(config))
	route := route.NewRoute(routeGroup)
	route.RegisterProduct(productHand)

	r.Run(os.Getenv("PORT"))

}
