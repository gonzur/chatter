import { ChatCLient, RoomBrowser } from "components";
import { useState } from "react";
import style from "./App.module.css";

enum Screens {
  roomBrowser = 0,
  chatClient = 1,
}

const App = () => {
  const user = `dummyUser${new Date().getSeconds()}`;

  const [room, setRoom] = useState<string>("");
  const [screen, setScreeen] = useState<Screens>(Screens.roomBrowser);

  const roomSelectedTransition = (selectedRoom: string) => {
    setRoom(selectedRoom);
    setScreeen(Screens.chatClient);
  };

  return (
    <div className={style.center}>
      {screen === Screens.roomBrowser && (
        <RoomBrowser transition={roomSelectedTransition} />
      )}
      {screen === Screens.chatClient && (
        <ChatCLient username={user} roomID={room} />
      )}
    </div>
  );
};

export default App;
