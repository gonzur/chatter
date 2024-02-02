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
          <div style={{ display: "flex", justifyContent: "end" }}>
            <div data-testid="chatMessage" className={styles.text}>
              {message}
            </div>
          </div>
          <div className={styles.status}>
            <div data-testid="chatUser">{sender}</div>
            <div style={{ width: ".25rem" }} />
            <div className={styles.sentOn}>{sentOn}</div>
          </div>
        </div>
      ))}
      <div ref={endAnchor} />
    </div>
  );
};

export default HistoryView;
