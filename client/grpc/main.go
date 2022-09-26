package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	v1 "github.com/sajanjswl/sandbox-service/gen/go/sandbox/v1alpha1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const apiVersion = "v1"

func main() {

	address := flag.String("server", "localhost:8000", "gRPC server in format host:port")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	register(c, ctx)

	// login(c, ctx)

	// update(c, ctx)

}

func register(c v1.AuthServiceClient, ctx context.Context) {

	fmt.Println("I  was here")
	req := &v1.RegisterUserRequest{
		User: &v1.User{
			EmailId:      "sjnjaiswal3@gmail.com",
			Password:     "password1",
			Name:         "Sajan",
			MobileNumber: "+917064274923",
		},
	}

	resp, err := c.RegisterUser(ctx, req)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)

}

func login(c v1.AuthServiceClient, ctx context.Context) {

	req1 := &v1.LoginUserRequest{
		EmailId:  "sjnjaiswal@gmail.com",
		Password: "password1",
	}
	res1, err := c.LoginUser(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("login result: <%+v>\n\n", res1)

}
