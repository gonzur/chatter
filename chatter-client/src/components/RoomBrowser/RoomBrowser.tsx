import { useEffect, useState } from "react";
import { AiOutlineSearch } from "react-icons/ai";
import styles from "./RoomBrowser.module.css";

interface RoomInfo {
  name: string;
  memberCount: number;
}

interface RoomBrowserProps {
  transition: (roomID: string) => void;
}

const RoomBrowser = (props: RoomBrowserProps) => {
  const { transition } = props;
  const [roomList, setRoomList] = useState<RoomInfo[]>([
    { name: "room 1asdfdsf", memberCount: 9 },
    { name: "room 1", memberCount: 9 },
    { name: "room 1asdfsafdas", memberCount: 9 },
    { name: "room 2", memberCount: 9 },
    { name: "room 3", memberCount: 9 },
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

  const [query, setQuery] = useState("");
  const [roomName, setRoomName] = useState("");

  useEffect(fetchRooms, []);

  return (
    <div className={styles.container}>
      <h1 className={styles.header}>Select your room</h1>
      <div className={styles.searchBackdrop}>
        <input
          value={query}
          onChange={(ev) => {
            setQuery(ev.target.value);
          }}
          className={styles.searchBar}
          placeholder="Search..."
        />
        <button
          type="button"
          aria-label="Search"
          className={styles.searchButton}
        >
          <AiOutlineSearch size="1.5rem" />
        </button>
      </div>
      <div className={styles.roomListBorder}>
        <div className={styles.roomList}>
          <div className={styles.card}>
            <span className={styles.inactive}>Inactive</span>
            <input
              placeholder="Room name..."
              value={roomName}
              onChange={(e) => setRoomName(e.target.value)}
              className={styles.roomNameInput}
            />
            <div className={styles.bottomGroup}>
              <p className={styles.members}>0 Users</p>
              <button
                className={styles.joinButton}
                type="button"
                onClick={() => transition(roomName)}
              >
                Create
              </button>
            </div>
          </div>
          {roomList
            .filter((value) => value.name.includes(query))
            .map((room) => (
              <div id={room.name} className={styles.card}>
                <span className={styles.active}>Active</span>
                <h2 className={styles.roomName}>{room.name}</h2>
                <div className={styles.bottomGroup}>
                  <p className={styles.members}>{room.memberCount} Users</p>
                  <button
                    className={styles.joinButton}
                    type="button"
                    onClick={() => transition(room.name)}
                  >
                    Join
                  </button>
                </div>
              </div>
            ))}
        </div>
      </div>
    </div>
  );
};

export default RoomBrowser;
