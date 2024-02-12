import { ChatCLient } from "components";
import style from "./App.module.css";

const App = () => {
  const user = `dummyUser${new Date().getSeconds()}`;
  return (
    <div className={style.center}>
      <ChatCLient username={user} roomID="test" />
    </div>
  );
};

export default App;
