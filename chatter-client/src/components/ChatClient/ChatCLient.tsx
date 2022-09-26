import { Paper } from "@mui/material";
import HistoryView from "components/HistoryView";
import { dummyMessageList } from "data/mock";
import { useState } from "react";
import styles from "./ChatClient.module.css";

const ChatCLient = () => {
  const [messages] = useState(dummyMessageList);

  return (
    <Paper className={styles.view} elevation={3}>
      <HistoryView messages={messages} />
    </Paper>
  );
};

export default ChatCLient;
