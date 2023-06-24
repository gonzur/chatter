import ChatInput from "components/ChatInput";
import HistoryView from "components/HistoryView";
import { Message } from "data/types";
import { MapRenderableMessage } from "data/types/Message";
import { formatToTwelveHourDate } from "helpers/format";

import { useCallback, useEffect, useState } from "react";
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
  const [messages, setMessages] = useState<MapRenderableMessage[]>([]);

  const sendMessage = useCallback(
    (message: Message) => {
      if (connected) {
        socket?.send(JSON.stringify(message));
      }

      setMessages([...messages, { ...message, id: messages.length - 1 }]);
    },
    [messages, socket, connected]
  );

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
      socket.onopen = () => {
        setConnected(true);
        const message: Message = {
          sender: username,
          sentOn: formatToTwelveHourDate(new Date()),
          message: `New user "${username}" has joined. Say hello everyone!`,
        };
        sendMessage(message);
      };

      socket.onmessage = (event) => {
        const message = JSON.parse(event.data) as Message;
        setMessages([...messages, { ...message, id: messages.length - 1 }]);
      };
    }
  }, [socket, messages]);

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
