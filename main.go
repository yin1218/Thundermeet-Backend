package main

import (
	"fmt"
	"log"
	"os"
	"thundermeet_backend/app/config"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"

	"thundermeet_backend/app/middleware/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	_ "thundermeet_backend/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title ThunderMeet APIs
// @version 1.0
// @description This is the backend server for the Thundermeet App.

// @contact.name Wu, Chien Yin and Yeh, Hsiao Li

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host thundermeet-backend.herokuapp.com/
// schemes http
func main() {
	fmt.Println("Good Morning!")
	//read env
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	//get env
	// port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	// db, ormErr := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	db, ormErr := dao.Initialize(dbUrl)
	if ormErr != nil {
		panic(ormErr)
	}
	migrateErr := db.AutoMigrate(&model.User{})
	if migrateErr != nil {
		return
	}

	//init server
	app := gin.Default()

	app.Use(cors.CORSMiddleware())

	// set swagger docs
	url := ginSwagger.URL("https://thundermeet-backend.herokuapp.com/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	app.GET("/hc", func(c *gin.Context) {
		fmt.Println("Good hc!")
		c.JSON(200, gin.H{
			"message": "health check",
		})
	})
	config.RouteUsers(app)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
