/* eslint-disable jsx-a11y/no-noninteractive-tabindex */
import { MapRenderableMessage } from "data/types/Message";
import { useEffect, useRef } from "react";
import styles from "./HistoryView.module.css";

interface HistoryViewProps {
  username: string;
  messages: MapRenderableMessage[];
}

const HistoryView = ({ username, messages }: HistoryViewProps) => {
  const endAnchor = useRef<HTMLDivElement | null>(null);
  useEffect(() => {
    if (endAnchor.current) {
      endAnchor.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);
  return (
    <div className={styles.view} tabIndex={0}>
      {messages.map(({ sender, message, sentOn, id }) => (
        <div key={id} className={styles.message}>
          <div
            style={{ justifyContent: username === sender ? "end" : "start" }}
            className={styles.flex}
          >
            <p
              data-testid="chatMessage"
              className={`${styles.text} ${
                username === sender ? styles.right : styles.left
              }`}
            >
              {message}
            </p>
          </div>
          <div
            style={{ justifyContent: username === sender ? "end" : "start" }}
            className={styles.status}
          >
            <p data-testid="chatUser">{sender}</p>
            <div style={{ width: ".25rem" }} />
            <p className={styles.sentOn}>{sentOn}</p>
          </div>
        </div>
      ))}
      <div ref={endAnchor} />
    </div>
  );
};

export default HistoryView;
