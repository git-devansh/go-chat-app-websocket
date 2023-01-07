import React from "react";

function ChatBox({ chats }) {
  return (
    <div className="chatbox">
      <h2>Chats</h2>
      {chats.map((chat, index) => {
        if (chat.chathistory) {
          return chat.chathistory.map((history, index) => (
            <p key={index + "__history"} className="chat__text">
              {history}
            </p>
          ));
        }
        return (
          <p key={index + "__newMessage"} className="chat__text">
            {chat.message}
          </p>
        );
      })}
    </div>
  );
}

export default ChatBox;
