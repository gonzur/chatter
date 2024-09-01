import { apiBaseUrl } from "data/constants";
import { useState } from "react";
import styles from "./Login.module.css";

interface LoginProps {
  onError: () => void;
  onSuccess: () => void;
}

const Login = (props: LoginProps) => {
  const { onError, onSuccess } = props;

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const verifyCredentials = () => {
    fetch(`${apiBaseUrl}/login`, {
      body: JSON.stringify({ username, password }),
    }).then((res) => {
      if (res.ok) {
        onSuccess();
      } else {
        onError();
      }
    });
  };

  return (
    <div className={styles.content}>
      <div className={styles.card}>
        <h1 className={styles.header}>Sign in</h1>

        <input
          className={styles.loginInput}
          id="username"
          type="text"
          placeholder="Username..."
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />

        <input
          className={styles.loginInput}
          id="password"
          type="password"
          placeholder="Password..."
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className={styles.signIn}
          type="button"
          onClick={verifyCredentials}
        >
          Sign In
        </button>
      </div>
    </div>
  );
};

export default Login;
