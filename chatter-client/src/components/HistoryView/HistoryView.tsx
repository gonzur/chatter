import { MapRenderableMessage } from "data/types/Message";
import { useEffect, useRef } from "react";
import styles from "./HistoryView.module.css";

interface HistoryViewProps {
  messages: MapRenderableMessage[];
}

const HistoryView = ({ messages }: HistoryViewProps) => {
  const endAnchor = useRef<HTMLDivElement | null>(null);
  useEffect(() => {
    if (endAnchor.current) {
      endAnchor.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);
  return (
    <div className={styles.view}>
      {messages.map(({ sender, message, sentOn, id }) => (
        <div key={id} className={styles.message}>
          <div className={styles.status}>
            <div className={styles.sentOn}>{sentOn}</div>
            <div style={{ width: ".5rem" }} />
            <div data-testid="chatUser">{sender}:</div>
          </div>
          <div data-testid="chatMessage" className={styles.text}>
            &quot;{message}&quot;
          </div>
        </div>
      ))}
      <div ref={endAnchor} />
    </div>
  );
};

export default HistoryView;
