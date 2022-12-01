import { useState } from "react";
import { AiOutlineSend } from "react-icons/ai";
import styles from "./ChatInput.module.css";

interface ChatInputProps {
  onSubmit: (data: string) => void;
}

const ChatInput = ({ onSubmit }: ChatInputProps) => {
  const [text, setText] = useState("");

  return (
    <div className={styles.flex}>
      <input
        className={styles["chat-line"]}
        onChange={(e) => setText(e.target.value)}
        placeholder="Type to chat..."
      />

      <button
        className={styles["btn-round"]}
        onClick={() => onSubmit(text)}
        type="button"
      >
        <AiOutlineSend size="1.5rem" />
      </button>
    </div>
  );
};

export default ChatInput;
