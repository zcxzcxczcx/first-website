package main

import (
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"

	"github.com/gin-gonic/gin"
)

func main() {
	profiler.Start(profiler.Config{
		ApplicationName: "simple.golang.app",
		ServerAddress:   "http://121.196.11.0:4040",
	})
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {

		c.JSON(200, gin.H{"hello": 11})
	})
	r.Run(":8000")
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
