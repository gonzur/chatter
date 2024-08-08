import { apiBaseUrl } from "data/constants";
import { useState } from "react";

interface LoginProps {
  onError: () => {};
  onSuccess: () => {};
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
    <div className="card">
      <h2>Login</h2>
      <label htmlFor="username">
        Username
        <input
          id="username"
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
      </label>
      <label htmlFor="password">
        Password
        <input
          id="password"
          type="text"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </label>
      <button type="button" onClick={verifyCredentials}>
        Sign In
      </button>
    </div>
  );
};

export default Login;
