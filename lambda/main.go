package main

import (
    "context"
    "fmt"
    "os"

    "github.com/aws/aws-lambda-go/lambda"
    "google.golang.org/grpc"
    "lambda/proto"
)

type MyEvent struct {
    ID string `json:"id"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
    address := os.Getenv("GRPC_SERVER_ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        return "", fmt.Errorf("did not connect: %v", err)
    }
    defer conn.Close()

    client := proto.NewUserServiceClient(conn)
    req := &proto.GetUserRequest{Id: event.ID}
    res, err := client.GetUser(ctx, req)
    if err != nil {
        return "", fmt.Errorf("could not get user: %v", err)
    }

    return fmt.Sprintf("User: %s, Age: %d", res.GetName(), res.GetAge()), nil
}

func main() {
    lambda.Start(HandleRequest)
}