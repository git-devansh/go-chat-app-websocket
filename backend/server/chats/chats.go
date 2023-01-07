package chats

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/chat-app/proto"
)

// Initilize and get the stored value.
func InitChatHistory() []string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	resp, err := c.GetChats(context.Background(), &pb.NoParam{})
	if err != nil {
		panic(err)
	}

	fmt.Println("RESPONSE::", resp.Chats)
	return resp.Chats
}

// This function use to add new message
func StoreMessage(newMessage string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	resp, err := c.AddMessage(context.Background(), &pb.Message{Message: newMessage})
	if err != nil {
		panic(err)
	}

	fmt.Println("AddMessage::", resp)
}
