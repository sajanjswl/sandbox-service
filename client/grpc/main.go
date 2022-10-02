package main

import (
	"context"
	"flag"
	"time"

	v1alpha1 "github.com/sajanjswl/sandbox-service/gen/go/sandbox/v1alpha1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	address := flag.String("server", "localhost:8000", "gRPC server in format host:port")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1alpha1.NewSandboxServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// register(c, ctx)

	login(c, ctx)

}

func register(c v1alpha1.SandboxServiceClient, ctx context.Context) {

	req := &v1alpha1.RegisterUserRequest{
		User: &v1alpha1.User{
			EmailId:      "sjnjaiswal@gmail.com",
			Password:     "password1",
			Name:         "Sajan",
			MobileNumber: "+91 789076552",
		},
	}

	resp, err := c.RegisterUser(ctx, req)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)

}

func login(c v1alpha1.SandboxServiceClient, ctx context.Context) {

	req1 := &v1alpha1.LoginUserRequest{
		EmailId:  "sjnjaiswal@gmail.com",
		Password: "password1",
	}
	res1, err := c.LoginUser(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("login result: <%+v>\n\n", res1)

}
