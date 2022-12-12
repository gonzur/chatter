import ChatInput from "components/ChatInput";
import HistoryView from "components/HistoryView";
import { Message } from "data/types";
import { formatToTwelveHourDate } from "helpers/format";

import { useEffect, useState } from "react";
import styles from "./ChatClient.module.css";

interface ChatCLientProps {
  username: string;
  roomID: string;
}

const ChatCLient = ({ username, roomID }: ChatCLientProps) => {
  const port = 8080;
  const socketUrl = `ws://localhost:${port}/api/chat/join-room?userID=${username}&roomID=${roomID}`;

  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [connected, setConnected] = useState(false);
  const [messages, setMessages] = useState<Message[]>([]);

  const sendMessage = (message: Message) => {
    if (connected) {
      socket?.send(message.toString());
    }

    setMessages([...messages, message]);
  };

  useEffect(() => {
    try {
      if (socket === null) {
        const ws = new WebSocket(socketUrl);
        setSocket(ws);
        return () => {
          ws.close();
          setSocket(null);
        };
      }
    } catch (e) {
      return () => {};
    }
    return () => {};
  }, [username, roomID]);

  useEffect(() => {
    if (socket !== null) {
      socket.addEventListener("open", () => {
        setConnected(true);
        const message: Message = {
          sender: username,
          sentOn: formatToTwelveHourDate(new Date()),
          message: `New user "${username}" has joined. Say hello everyone!`,
        };
        sendMessage(message);
      });

      socket.addEventListener("message", (event) => {
        const message = event.data as Message;
        setMessages([...messages, message]);
      });
    }
  }, [socket]);

  return (
    <div className={styles.view}>
      <HistoryView messages={messages} />
      <div className={styles["chat-input"]}>
        <ChatInput
          onSubmit={(data) => {
            const formattedDate = formatToTwelveHourDate(new Date());
            const message: Message = {
              sender: username,
              message: data,
              sentOn: formattedDate,
            };
            sendMessage(message);
          }}
        />
      </div>
    </div>
  );
};

export default ChatCLient;
