import {} from "data/types";
import { MapRenderableMessage } from "data/types/Message";
import styles from "./HistoryView.module.css";

interface HistoryViewProps {
  messages: MapRenderableMessage[];
}

const HistoryView = ({ messages }: HistoryViewProps) => {
  const min = messages.length - 10 < 0 ? 0 : messages.length - 10;
  return (
    <div className={styles.view}>
      {messages
        .slice(min, messages.length)
        .map(({ sender, message, sentOn, id }) => (
          // eslint-disable-next-line react/no-array-index-key
          <div key={id} className={styles.message}>
            <div className={styles.status}>
              <div>{sentOn}</div>
              <div data-testid="chatUser">{sender}:</div>
            </div>
            <div data-testid="chatMessage" className={styles.text}>
              &quot;{message}&quot;
            </div>
          </div>
        ))}
    </div>
  );
};

export default HistoryView;
