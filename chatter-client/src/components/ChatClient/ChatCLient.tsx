import ChatInput from "components/ChatInput";
import HistoryView from "components/HistoryView";
import { Message } from "data/types";
import useTwelveHourDate from "hooks/useTwelveHourDate";
import { useState } from "react";
import styles from "./ChatClient.module.css";

interface ChatCLientProps {
  userID: string;
}

const ChatCLient = ({ userID }: ChatCLientProps) => {
  const [messages, setMessages] = useState([] as Message[]);

  return (
    <div className={styles.view}>
      <HistoryView data-testid="oneday" messages={messages} />
      <div className={styles["chat-input"]}>
        <ChatInput
          onSubmit={(data) => {
            const { formattedDate } = useTwelveHourDate(new Date());
            const message: Message = {
              sender: userID,
              message: data,
              sentOn: formattedDate,
            };
            setMessages([...messages, message]);
          }}
        />
      </div>
    </div>
  );
};

export default ChatCLient;
