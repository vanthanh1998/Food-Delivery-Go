package main

import (
	"Food-Delivery/component/appctx"
	"Food-Delivery/component/uploadprovider"
	"Food-Delivery/middleware"
	"Food-Delivery/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	//jsByte, err := json.Marshal(test)
	//log.Println(string(jsByte), err) // {"id":1,"name":"200lab","addr":"what the hell"}
	//
	//json.Unmarshal([]byte("{\"id\":2,\"name\":\"200lab1998\",\"addr\":\"what the hell do you want naruto?\"}"), &test)
	//
	//log.Println(test)

	dsn := os.Getenv("MYSQL_CONN_STRING") // database connection string
	// MYSQL_CONN_STRING=>food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // database connection

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default() // server connection
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // H là map []
			"message": "pong",
		})
	})

	//r.Static("/static", "./static")

	v1 := r.Group("/v1")
	routes.SetupRoute(appContext, v1)
	routes.SetupAdminRoute(appContext, v1)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
