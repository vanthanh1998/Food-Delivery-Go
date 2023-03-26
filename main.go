package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id int `json:"id" gorm:"column:id"` // id of the restaurant
	Name string `json:"name" gorm:"column:name"` // TODO
	Addr string `json:"addr" gorm:"column:addr"` // address
}

func (Restaurant) TableName() string  {
	return "restaurants" // table name
}

// Trường hợp muốn update 1 filed nào đó về nil ~~ "" thì bắt buộc phải dùng con trỏ *,&
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string  {
	return Restaurant{}.TableName() // table name
}

func main()  {
	dsn := os.Getenv("MYSQL_CONN_STRING") // database connection string

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // database connection

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)

	r := gin.Default() // server connection
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // H là map []
			"message": "pong",
		})
	})
	// POST restaurant
	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", func(c *gin.Context) {
		var data Restaurant

		if err := c.ShouldBind(&data); err != nil {
			// ShouldBind: dùng để đọc và gán giá trị từ các request parameter vào các struct
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Create(&data)
		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	// get by id
	restaurants.GET("/:id", func(c *gin.Context) {
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: chuyển đổi một chuỗi số sang dạng số nguyên (int)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
		}
		var data Restaurant

		db.Where("id = ?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	// get all
	restaurants.GET("", func(c *gin.Context) {
		var data []Restaurant

		// pagging data
		type Paging struct {
			Page int `json:"page" form:"page"`
		    Limit int `json:"limit" form:"limit"`
        }

        var pagingData Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		if pagingData.Limit <= 0 {
            pagingData.Limit = 5
        }

		db.Offset((pagingData.Page - 1 ) * pagingData.Limit).
			Order("id desc").
			Limit(pagingData.Limit).
			Find(&data)

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	// update by id
	restaurants.PUT("/:id", func(c *gin.Context) { // use postman nhớ chuyển sang kiểu json
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: chuyển đổi một chuỗi số sang dạng số nguyên (int)
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
			"data" : data,
		})
	})

	// delete by id
	restaurants.DELETE("/:id", func(c *gin.Context) {
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: chuyển đổi một chuỗi số sang dạng số nguyên (int)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"data" : "Delete successful",
		})
	})


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