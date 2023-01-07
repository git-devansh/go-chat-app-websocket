import { useEffect, useState, useRef } from "react";
import ChatBox from "./components/Chats";
import Header from "./components/Header";
import { connect, sendMessage } from "./websocket";
import { namesList } from "./utils/namelist";

let userName = "";

function App() {
  const [chats, setChats] = useState([]);
  const [inputMsg, setInputMsg] = useState("");

  // returns random name value from namesList.
  const getRandomName = () => {
    return namesList[Math.floor(Math.random() * namesList.length)];
  };

  useEffect(() => {
    userName = getRandomName();
    // websocket
    connect((msg) => {
      setChats((prevState) => [...prevState, msg]);
    });
  }, []);

  const onMessageSend = () => {
    let jsonData = {};
    jsonData["action"] = "message";
    jsonData["message"] = inputMsg;
    jsonData["username"] = userName;
    sendMessage(JSON.stringify(jsonData));
  };

  return (
    <div>
      <Header />
      <ChatBox chats={chats} />
      <input
        type="text"
        placeholder="Enter message.."
        onChange={(e) => setInputMsg(e.target.value)}
      ></input>
      <button onClick={onMessageSend}>Send message</button>
    </div>
  );
}

export default App;
