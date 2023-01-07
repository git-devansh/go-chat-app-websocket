package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/chat-app/proto"
	"google.golang.org/grpc"
)

var chatHistory []string

type chatServiceServer struct {
	pb.ChatServiceServer
}

func (s *chatServiceServer) AddMessage(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	fmt.Println("chatHistory", chatHistory)
	chatHistory = append(chatHistory, message.Message)
	fmt.Println("chatHistory", chatHistory)
	return &pb.Message{Message: "Stored in the slice"}, nil
}

func (s *chatServiceServer) GetChats(ctx context.Context, message *pb.NoParam) (*pb.Chats, error) {
	return &pb.Chats{Chats: chatHistory}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &chatServiceServer{})

	fmt.Println("Server starting at :4040")
	log.Fatal(grpcServer.Serve(listener))

}
