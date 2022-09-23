import { TextField } from "@mui/material";

const ChatBox = () => (
  <div>
    <TextField fullWidth label="fullwidth" />
  </div>
);

export default ChatBox;

/**
 * * Requirements:
 * 1. must accept text into input field
 * 2. when enter is pressed it will send a message
 * 3. clears out message on enter
 */
