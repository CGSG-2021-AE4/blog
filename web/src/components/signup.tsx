import React from "react";
import ReactDOM from "react-dom/client";
import * as Auth from "../auth";
import * as Notifications from "./notification";

class signinReq {
  username: string;
  email:    string;
  password: string;
}

class signinResp {
  msg: string;
}

async function HandleSubmit(e) {
  e.preventDefault();

  let username =     (document.getElementById("username")!     as HTMLInputElement).value;
  let email    =     (document.getElementById("email")!        as HTMLInputElement).value;
  let password =     (document.getElementById("password")!     as HTMLInputElement).value;
  let passwordConf = (document.getElementById("passwordConf")! as HTMLInputElement).value;

  if (password != passwordConf) {
    Notifications.Push({type: "error", msg: "Error: passwords do not match"})
  }
  
  const resp = await fetch(window.location.origin + "/api/user/reg", {
    method: "POST",
    body: JSON.stringify({
      username: username,
      email: email,
      password: password,
    } as signinReq),
  });
  const res = await resp.json();
  if (res.err != undefined) {
    // Got error
    Notifications.Push({type: "error", msg: "Error: " + res.err});
    return;
  }
  Notifications.Push({type: "ok", msg: res.msg});
  setTimeout(() => {
    window.location.href = "/"; // Redirect
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
      minWidth: "22em",
    }}>
      <div style={labelDivStyle}>
        <label htmlFor="username">Username:</label><input id="username" name="username" type="text"/>
      </div>
      <div style={labelDivStyle}>
        <label htmlFor="email">Email:</label><input id="email" name="email" type="text"/>
      </div>
      <div style={labelDivStyle}>
        <label htmlFor="password">Password:</label><input id="password" name="password" type="text"/>
      </div>
      <div style={labelDivStyle}>
        <label htmlFor="passwordConf">Confirm password:</label><input id="passwordConf" name="passwordConf" type="text"/>
      </div>
      <input onSubmit={HandleSubmit} type="submit" style={{
        marginTop: "2em",
      }} value="Sign up" />
    </form>
  );
}