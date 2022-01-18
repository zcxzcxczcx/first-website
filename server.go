package main

import (
	"fmt"
	"time"
)

const (
	port = ":50051"
)

func main() {
	for i := 0; i < 10; i++ {

		go func() {
			fmt.Printf("ddddddddddddddddd=%v\n", i)
		}()
	}
	time.Sleep(100 * time.Second)
}

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	proto.RegisterFirstWebsiteServer(s, &controllers.Server{})
// 	// Register reflection service on gRPC server.
// 	reflection.Register(s)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

// func main() {
// 	r := gin.Default()
// 	store := cookie.NewStore([]byte("secret"))
// 	r.Use(sessions.Sessions("mysession", store))
// 	r.GET("/hello", func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		if session.Get("hello") != "world" {
// 			session.Set("hello", "world")
// 			session.Save()
// 		}
// 		c.JSON(200, gin.H{"hello": session.Get("hello")})
// 	})
// 	r.RunTLS(":8000", "/Users/abc/127.0.0.1.pem",
// 		"/Users/abc/127.0.0.1-key.pem")

// }
