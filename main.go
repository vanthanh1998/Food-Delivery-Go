package main

import (
	"Food-Delivery/component/appctx"
	"Food-Delivery/component/tokenprovider/jwt"
	"Food-Delivery/component/uploadprovider"
	"Food-Delivery/middleware"
	userstore "Food-Delivery/module/user/store"
	"Food-Delivery/pubsub/localpb"
	"Food-Delivery/routes"
	"Food-Delivery/skio"
	"Food-Delivery/subscriber"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
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

	ps := localpb.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	// setup subscribers
	//subscriber.Setup(appContext, context.Background())
	_ = subscriber.NewEngine(appContext).Start()

	r := gin.Default() // serv

	r.StaticFile("/demo/", "./demo.html")

	// er connection
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

	//startSocketIOServer(r, appContext)
	skio.NewEngine().Run(appContext, r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
	server, _ := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("")
		fmt.Println("Socket connected:", s.ID(), " IP:", s.RemoteAddr())

		s.Join("Shipper")
		//ticker := time.NewTicker(time.Second)
		//i := 0
		//for {
		//	<-ticker.C
		//	i++
		//	s.Emit("test", i)
		//}

		return nil
	})

	//go func() {
	//	for range time.NewTicker(time.Second).C {
	//		server.BroadcastToRoom("/", "Shipper", "test", "thanhrain")
	//	}
	//}()

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	// authenticate
	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		// validate token
		// If false: s.Close(), and return
		// If true
		// => UserID
		// Fetch db find user by ID
		// Here: s belongs to who? (user_id)
		// We need a map[user_id][]skio.Conn

		db := appCtx.GetMailDBConnection()
		store := userstore.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			s.Close()
			return
		}

		user.Mask(false)

		s.Emit("your_profile", user)
	})
	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
		log.Println("test: ", msg)
	})

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
		fmt.Println("server receive notice", p.Name, p.Age)

		p.Age = 33
		s.Emit("notice", p)
	})

	go server.Serve()

	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))
}
