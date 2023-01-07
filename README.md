# go-chat-app-websocket

Websocket client able to connect to the backend server. Users can send messages to other clients connected to the same server.

for websocket used this package: https://github.com/gorilla/websocket

---
There is a client written using React (Create React App). 

# To install required dependencies: 
npm install
# To Run the project
npm start

The backend is of two main parts
1) server - where handlers WebSocket server code.
2) grpc   - Additional Go program to store and retrieve the chats using protobuf.

Please run both in different terminals in order to work.

go run main.go
