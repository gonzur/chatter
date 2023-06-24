import { useCallback, useState } from "react";
import { AiOutlineSend } from "react-icons/ai";
import styles from "./ChatInput.module.css";

interface ChatInputProps {
  onSubmit: (data: string) => void;
}

const ChatInput = ({ onSubmit }: ChatInputProps) => {
  const [message, setMessage] = useState("");

  const sendMessage = useCallback(() => {
    onSubmit(message);
    setMessage("");
  }, [message]);

  return (
    <div className={styles.flex}>
      <input
        className={styles["chat-line"]}
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            sendMessage();
          }
        }}
        placeholder="Type to chat..."
      />

      <button
        className={styles["btn-round"]}
        onClick={() => sendMessage()}
        type="button"
      >
        <AiOutlineSend size="1.5rem" />
      </button>
    </div>
  );
};

export default ChatInput;
