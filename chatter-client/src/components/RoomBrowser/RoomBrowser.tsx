import { useEffect, useState } from "react";
import styles from "./RoomBrowser.module.css";

interface RoomInfo {
  name: string;
  memberCount: number;
}

const RoomBrowser = () => {
  const [roomList, setRoomList] = useState<RoomInfo[]>([
    { name: "room 1asdfdsf", memberCount: 9 },
    { name: "room 1", memberCount: 9 },
    { name: "room 1asdfsafdas", memberCount: 9 },
    { name: "room 1", memberCount: 9 },
    { name: "room 1", memberCount: 9 },
    { name: "1234567890123456", memberCount: 9 },
    { name: "wwwwwwwwwwwwwwww", memberCount: 9 },
  ]);
  const fetchRooms = () => {
    fetch("/api/chat/list")
      .then((res) => res.json())
      .then((rooms: RoomInfo[]) => {
        setRoomList(rooms);
      });
  };

  useEffect(fetchRooms, []);

  return (
    <div className={styles.container}>
      <h1 className={styles.header}>Select your room</h1>
      <div className={styles.roomList}>
        {roomList.map((room) => (
          <div id={room.name} className={styles.card}>
            <span className={styles.active}>Active</span>
            <h2 className={styles.roomName}>{room.name}</h2>
            <p className={styles.members}>{room.memberCount} Users</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default RoomBrowser;
