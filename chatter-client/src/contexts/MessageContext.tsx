import { Message } from "data/types";
import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";

interface MessageContextValue {
  messages: Message[];
}

interface MessageContextProps {
  children: ReactNode;
  username: string;
}

const MessageContextInit: MessageContextValue = {
  messages: [],
};

const MessageContext = createContext(MessageContextInit);

const MessageContextProvider = ({
  children,
  username,
}: MessageContextProps) => {
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<Message[]>(
    MessageContextInit.messages
  );

  useEffect(() => {
    setSocket(new WebSocket(""));
  }, [username]);

  useEffect(() => {
    if (socket !== null) {
      socket.addEventListener("open", () => {
        const message: Message = {
          sender: username,
          sentOn: Date().toString(),
          message: `New user "${username}" has joined. Say hello everyone!`,
        };
        socket.send(message.toString());
      });

      socket.addEventListener("message", (event) => {
        const message = event.data as Message;
        setMessages([...messages, message]);
      });
    }
  }, [socket]);

  <MessageContext.Provider value={{ messages }}>
    {children}
  </MessageContext.Provider>;
};

const useMessageContext = () => useContext(MessageContext);

export { MessageContextProvider, useMessageContext };
