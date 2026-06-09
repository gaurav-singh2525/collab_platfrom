import { useEffect, useState } from "react";

function WebSocketTest() {
  const [socket, setSocket] = useState(null);

  const [message, setMessage] = useState("");

  const [messages, setMessages] = useState([]);

  const roomId = "room1";

  useEffect(() => {
    const ws = socketService.connect(roomId);

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      if (message.type === "code_change") {
        setCode(message.code);
      }
    };

    return () => ws.close();
  }, []);



  
  const sendMessage = () => {
    if (!socket || socket.readyState !== WebSocket.OPEN) {
      return;
    }

    socket.send(message);
    setMessage("");
  };

  return (
    <div>
      <h1>WebSocket Test</h1>

      <input value={message} onChange={(e) => setMessage(e.target.value)} />

      <button onClick={sendMessage}>Send</button>

      <ul>
        {messages.map((msg, index) => (
          <li key={index}>{msg}</li>
        ))}
      </ul>
    </div>
  );
}
export default WebSocketTest;
