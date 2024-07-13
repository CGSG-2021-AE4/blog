import React from "react";
import ReactDOM from "react-dom/client";
import * as Auth from "../auth";

export default function Header() {
  let accountInfo = (
    <a href={window.location.origin + "/login"}>
      <h3>Login</h3>
    </a>
  );
  if (Auth.IsAuthorized()) {
    accountInfo = (<>
      <a href={window.location.origin + "/account"}>
        <h3>{Auth.GetUsername()}</h3>
      </a>
    </>);
  }

  return (<div style={{
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",

    margin: "1em",
    marginBottom: "var(--vertical-gap)",
  }}>
    <a href="/">
      <img id="logo" src={window.location.origin + "/bin/imgs/logo.svg"}/>
    </a>
    <a href={window.location.origin + "/article/create"}>
        <h3>Create article</h3>
    </a>
    {accountInfo}
  </div>);
}