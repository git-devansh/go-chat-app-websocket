package handlers

import (
	"fmt"
	"net/http"

	"example.com/chat-app/server/chats"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsJsonRequest)
var clients = make(map[WebSocketConnection]string)

// var chatSlice []string

// wrapper to websocket connection
type WebSocketConnection struct {
	*websocket.Conn
}

// reponse send back
type WsJsonResponse struct {
	Action      string   `json:"action"`
	Message     string   `json:"message"`
	ChatHistory []string `json:"chathistory"`
}

// JSOn request coming
type WsJsonRequest struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Upgrades connection to websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Client connected..")

	var response WsJsonResponse
	response.Action = "stats"
	response.Message = "Connected to server"

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		fmt.Println(err)
	}

	chats := chats.InitChatHistory()

	response.Action = "message"
	response.Message = ""
	response.ChatHistory = chats

	err = ws.WriteJSON(response)
	if err != nil {
		fmt.Println(err)
	}

	go ListenForWs(&conn)
}

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var request WsJsonRequest

	for {
		err := conn.ReadJSON(&request)
		if err != nil {
			// nothing
		} else {
			request.Conn = *conn
			wsChan <- request
		}
	}
}

func ListenToWsChannel() {
	var resp WsJsonResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "message":
			resp.Action = "message"
			newMessage := fmt.Sprintf("%s: %s", e.Username, e.Message)
			resp.Message = newMessage
			broadcastToAll(resp)
			chats.StoreMessage(newMessage)
			// Before gRPC implementation
			// chatSlice = append(chatSlice, fmt.Sprintf("%s: %s", e.Username, e.Message))
			// fmt.Println("Chats", chatSlice)
			//
			// case "left":
			// 	delete(clients, e.Conn)
			// 	broadcastToAll(resp)
		}
	}
}

func broadcastToAll(resp WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(resp)
		if err != nil {
			fmt.Println(err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}
