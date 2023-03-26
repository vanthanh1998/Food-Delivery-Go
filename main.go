package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	Id int `json:"id" gorm:"column:id"` // id of the restaurant
	Name string `json:"name" gorm:"column:name"` // TODO
	Addr string `json:"addr" gorm:"column:addr"` // address
}

func (Restaurant) TableName() string  {
	return "restaurants" // table name
}

func main()  {
	dsn := os.Getenv("MYSQL_CONN_STRING") // database connection string

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // database connection

	if err != nil {
		log.Fatalln(err)
	}

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
	var restaurants Restaurant
    if err := db.Where("id = ?", 6).Find(&restaurants).Error; err!= nil {
        log.Println(err)
    }

    log.Println(restaurants)

    // ----------------------------Update--------------------------------
	restaurants.Name = "Megumi Fushigoro"
	if err := db.Where("id = ?", 6).Updates(&restaurants).Error; err!= nil {
		log.Println(err)
	}
	log.Println(restaurants)

	// ----------------------------Delete--------------------------------
	if err := db.Table(Restaurant{}.TableName()).Where("id =?", 1).Delete(nil).Error; err!= nil {
        log.Println(err)
    }
    log.Println(restaurants)
}