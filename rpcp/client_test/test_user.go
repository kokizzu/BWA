package main

import (
	"BWA/rpcp"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(`localhost:9090`, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpcp.NewUserServiceClient(conn)

	in := &rpcp.RegisterUserInput{
		Name:       "yay",
		Occupation: "foo",
		Email:      "bar@baz.com",
		Password:   "test123",
	}
	out, err := client.RegisterUserGrpc(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
