import React from "react";
import ReactDOM from "react-dom/client";
import * as Auth from "../auth";

class loginReq {
  username: string;
  password: string;
}
class loginResp {
  token: string;
  msg: string;
}

async function HandleSubmit(e) {
  e.preventDefault();
  let username = (document.getElementById("username")! as HTMLInputElement).value;
  let password = (document.getElementById("password")! as HTMLInputElement).value;
  const r = await fetch("/login", {
    method: "POST",
    body: JSON.stringify({username: username, password: password} as loginReq),
  });
  // console.log(r);
  //Auth.Login("test", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA2NjAxMzAsImlzcyI6InRlc3QifQ.B43QQxQ8ZndTxw5FapK2rzXGY6HiL0FnpP6KuhTUO_A")
  const resp = (await r.json()) as loginResp;
  console.log(resp);
  Auth.Login(username, resp.token)
  window.location.href = "/"; // Redirect
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
      }} value="Login" />
    </form>
  );
}