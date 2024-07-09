import { ChatCLient, RoomBrowser } from "components";
import style from "./App.module.css";

const App = () => {
  const user = `dummyUser${new Date().getSeconds()}`;
  return (
    <div className={style.center}>
      <RoomBrowser />
      <ChatCLient username={user} roomID="test" />
    </div>
  );
};

export default App;
