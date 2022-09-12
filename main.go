package main

import (
	"log"
	"net/http"
	"os"

	"pheet-gin-backend/auth"
	handler "pheet-gin-backend/product/http"
	"pheet-gin-backend/product/repository"
	"pheet-gin-backend/product/service"
	"pheet-gin-backend/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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
	r.Use(CORS())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Bizcuitware Web!!!")
	})

	r.POST("/login", productHand.Login)

	routeGroup := r.Group("", auth.Protect([]byte(os.Getenv("SIGN"))))
	routeGroup.Use(cors.Default())
	route := route.NewRoute(routeGroup)
	route.RegisterProduct(productHand)

	r.Run(os.Getenv("PORT"))

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
