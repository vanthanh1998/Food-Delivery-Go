package main

import (
	"Food-Delivery/component/appctx"
	"Food-Delivery/middleware"
	"Food-Delivery/module/restaurant/transport/ginrestaurant"
	"Food-Delivery/module/upload/uploadtransport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`     // id of the restaurant
	Name string `json:"name" gorm:"column:name"` // TODO
	Addr string `json:"addr" gorm:"column:addr"` // address
}

func (Restaurant) TableName() string {
	return "restaurants" // table name
}

// Trường hợp muốn update 1 filed nào đó về nil ~~ "" thì bắt buộc phải dùng con trỏ *,&
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName() // table name
}

func main() {
	//test := Restaurant{
	//	Id:   1,
	//	Name: "200lab",
	//	Addr: "what the hell",
	//}
	//
	//jsByte, err := json.Marshal(test)
	//log.Println(string(jsByte), err) // {"id":1,"name":"200lab","addr":"what the hell"}
	//
	//json.Unmarshal([]byte("{\"id\":2,\"name\":\"200lab1998\",\"addr\":\"what the hell do you want naruto?\"}"), &test)
	//
	//log.Println(test)

	dsn := os.Getenv("MYSQL_CONN_STRING") // database connection string
	// MYSQL_CONN_STRING=>food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // database connection

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default() // server connection
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // H là map []
			"message": "pong",
		})
	})

	r.Static("/static", "./static")

	// POST restaurant
	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appContext))

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	// get by id
	restaurants.GET("/:id", func(c *gin.Context) {
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: string -> (int)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data Restaurant

		db.Where("id = ?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// get all
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	// update by id
	restaurants.PUT("/:id", func(c *gin.Context) {
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: string -> (int)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			// ShouldBind: dùng để đọc và gán giá trị từ các request parameter vào các struct
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Where("id = ?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"data":   data,
		})
	})

	// delete by id
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
