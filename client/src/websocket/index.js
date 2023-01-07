var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (callBack) => {
  console.log("WEBSOCKET: Attempting Connection...");

  socket.onopen = () => {
    console.log("WEBSOCKET: Successfully Connected");
  };

  socket.onclose = () => {
    console.log("WEBSOCKET: Closed Connection");
  };

  socket.onerror = (error) => {
    console.log("WEBSOCKET: Socket Error: ", error);
  };

  socket.onmessage = (msg) => {
    console.log("Message from server: ");
    callBack(JSON.parse(msg.data));
  };
};

let sendMessage = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

let leftUser = () => {
  window.onbeforeunload = () => {
    console.log("Leaving..");
    let jsonData = {};
    jsonData["action"] = "left";
    socket.send(JSON.stringify(jsonData));
  };
};

export { connect, sendMessage, leftUser };
