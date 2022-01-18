package main

import (
	"first_website/proto"
	"log"

	"github.com/gin-gonic/gin"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(ctx *gin.Context) {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := proto.NewFirstWebsiteClient(conn)

		// Contact the server and print out its response.
		// name := defaultName
		// if len(os.Args) > 1 {
		// 	name = os.Args[1]
		// }
		resp, err := client.Login(context.Background(), &proto.LoginReq{
			Name:     "zcx",
			Password: "123456",
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("####### get server Greeting response: %s", resp.Token)

		ctx.JSON(200, gin.H{"hello": ""})
	})
	r.Run(":8001")
}
