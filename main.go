package main

import (
	"Food-Delivery/component/appctx"
	"Food-Delivery/component/uploadprovider"
	"Food-Delivery/middleware"
	"Food-Delivery/module/restaurant/transport/ginrestaurant"
	"Food-Delivery/module/upload/uploadtransport/ginupload"
	"Food-Delivery/module/user/transport/ginuser"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // database connection

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider)

	r := gin.Default() // server connection
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // H là map []
			"message": "pong",
		})
	})

	//r.Static("/static", "./static")

	// POST restaurant
	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appContext))

	// user
	v1.POST("/register", ginuser.Register(appContext))

	restaurants := v1.Group("/restaurants")

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.GetIdRestaurant(appContext))
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//newRestaurant := Restaurant{
	//	Name: "megumi fushigoro",
	//	Addr: "Shukuna",
	//}
	////----------------------------create--------------------------------
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	// truyền &(con trỏ): để thay đổi dữ liệu trong table Restaurant
	//	log.Println(err)
	//}
	//log.Println("New success Id = ", newRestaurant.Id)

	// ----------------------------Read--------------------------------
	//var restaurants Restaurant
	//if err := db.Where("id = ?", 6).Find(&restaurants).Error; err!= nil {
	//    log.Println(err)
	//}
	//
	//log.Println(restaurants)
	//
	//// ----------------------------Update--------------------------------
	//restaurants.Name = "Megumi Fushigoro"
	//if err := db.Where("id = ?", 6).Updates(&restaurants).Error; err!= nil {
	//	log.Println(err)
	//}
	//log.Println(restaurants)
	//
	//// ----------------------------Update name = nil ~~ ""--------------------------------
	//newName := "ThanhRain"
	//updateData := RestaurantUpdate{Name: &newName}
	//if err := db.Where("id = ?", 3).Updates(&updateData).Error; err!= nil {
	//	log.Println(err)
	//}
	//log.Println(restaurants)
	//// ----------------------------Delete--------------------------------
	//if err := db.Table(Restaurant{}.TableName()).Where("id =?", 1).Delete(nil).Error; err!= nil {
	//    log.Println(err)
	//}
	//log.Println(restaurants)
}
