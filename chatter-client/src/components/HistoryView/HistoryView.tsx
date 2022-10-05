import { Message } from "data/types";
import styles from "./HistoryView.module.css";

interface HistoryViewProps {
  messages: Message[];
}

const HistoryView = ({ messages }: HistoryViewProps) => (
  <div className={styles.view}>
    {messages.map(({ sender, message, sentOn }) => (
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

export default HistoryView;
