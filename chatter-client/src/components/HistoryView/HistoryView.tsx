import { Message } from "data/types";
import styles from "./HistoryView.module.css";

interface HistoryViewProps {
  messages: Message[];
}

const HistoryView = ({ messages }: HistoryViewProps) => {
  const min = messages.length - 10 < 0 ? 0 : messages.length - 10;
  return (
    <div className={styles.view}>
      {messages
        .slice(min, messages.length)
        .map(({ sender, message, sentOn }) => (
          <div className={styles.message}>
            <div className={styles.status}>
              <div>{sentOn}</div>
              <div>{sender}:</div>
            </div>
            <div className={styles.text}>&quot;{message}&quot;</div>
          </div>
        ))}
    </div>
  );
};

export default HistoryView;
