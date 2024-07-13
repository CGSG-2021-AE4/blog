import React from "react";
import ReactDOM from "react-dom/client";
import * as Auth from "../auth";
import * as Notifications from "../components/notification";

class loginReq {
  username: string;
  password: string;
}

class loginResp {
  token: string;
  id: string;
  msg: string;
}

async function HandleSubmit(e) {
  e.preventDefault();
  let username = (document.getElementById("username")! as HTMLInputElement).value;
  let password = (document.getElementById("password")! as HTMLInputElement).value;

  const resp = await fetch(window.location.origin + "/api/user/login", {
    method: "POST",
    body: JSON.stringify({username: username, password: password} as loginReq),
  });

  const res = await resp.json();
  if (res.err != undefined) {
    // Got error
    Notifications.Push({type: "error", msg: "Error: " + res.err});
    return;
  }

  Auth.Login(res.id, username, res.token);
  Notifications.Push({type: "ok", msg: "Login complete"});
  (document.getElementById("username")! as HTMLInputElement).value = "";
  (document.getElementById("password")! as HTMLInputElement).value = "";
  setTimeout(() => {
    window.location.href = window.location.origin; // Redirect
  }, 1000);
}

// Some style constants
const labelDivStyle = {
  display: "flex",
  justifyContent: "space-between",
  marginBottom: "1em",
} as React.CSSProperties

export default function LoginForm() {
  return (
    <form onSubmit={HandleSubmit} style={{
      display: "flex",
      flexDirection: "column",
      padding: "1.5em",
      border: "1px solid var(--main-color)",
      borderRadius: "1em",
      minWidth: "18em",
    }}>
      <div style={labelDivStyle}>
        <label htmlFor="username">Username:</label><input id="username" name="username" type="text"/>
      </div>
      <div style={labelDivStyle}>
        <label htmlFor="password">Password:</label><input id="password" name="password" type="text"/>
      </div>
      <input onSubmit={HandleSubmit} type="submit" style={{
        marginTop: "2em",
        marginBottom: "0.5em",
      }} value="Login" />
      <div style={{display: "flex", justifyContent: "center"}}><div><a href={window.location.origin + "/signup"}>Signup</a> for new users</div></div>
    </form>
  );
}