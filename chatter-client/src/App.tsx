import { ChatCLient, Login, RoomBrowser } from "components";
import { useState } from "react";
import style from "./App.module.css";

enum Screens {
  roomBrowser = 0,
  chatClient,
  login,
}

const App = () => {
  const user = `dummyUser${new Date().getSeconds()}`;

  const [room, setRoom] = useState<string>("");
  const [screen, setScreen] = useState<Screens>(Screens.login);

  const roomSelectedTransition = (selectedRoom: string) => {
    setRoom(selectedRoom);
    setScreen(Screens.chatClient);
  };

  return (
    <div className={style.center}>
      {screen === Screens.login && (
        <Login onSuccess={() => {}} onError={() => {}} />
      )}
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
