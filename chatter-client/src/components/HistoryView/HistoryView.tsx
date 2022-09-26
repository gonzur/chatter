import { Message } from "data/types";
import styles from "./HistoryView.module.css";

interface HistoryViewProps {
  messages: Message[];
}

const MessageElement = ({ sender, sentOn, message }: Message) => (
  <div className={styles.message}>
    <div className={styles.status}>
      <div>{sentOn}</div>
      <div>{sender}:</div>
    </div>
    <div className={styles.text}>&quot;{message}&quot;</div>
  </div>
);

const HistoryView = ({ messages }: HistoryViewProps) => (
  <div className={styles.view}>
    {messages.map(({ sender, message, sentOn }) => (
      <MessageElement message={message} sender={sender} sentOn={sentOn} />
    ))}
  </div>
);

export default HistoryView;
