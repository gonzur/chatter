import ChatInput from "components/ChatInput";
import HistoryView from "components/HistoryView";
import { dummyMessageList } from "data/mock";
import { useState } from "react";
import styles from "./ChatClient.module.css";

const ChatCLient = () => {
  const [messages] = useState(dummyMessageList);

  return (
    <div className={styles.view}>
      <HistoryView data-testid="oneday" messages={messages} />
      <div className={styles["chat-input"]}>
        <ChatInput onSubmit={() => {}} />
      </div>
    </div>
  );
};

export default ChatCLient;
