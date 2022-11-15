import { ChatCLient } from "components";
import "./colors.css";
import style from "./App.module.css";

const App = () => (
  <div className={style.center}>
    <ChatCLient userID="dummyUser" />
  </div>
);

export default App;
