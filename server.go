package main

import (
	"log"
	"net"

	"first_website/proto"
	"first_website/server/controllers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterFirstWebsiteServer(s, &controllers.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

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
