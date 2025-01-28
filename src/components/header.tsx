import React from "react";
import ReactDOM from "react-dom/client";
import * as Auth from "../auth";

export default function Header() {
  let Username = "GUEST";
  let UsernameButton = "Login";
  let UsernameLink = window.location.origin + "/login";

  if (Auth.IsAuthorized()) {
    Username = Auth.GetUsername();
    UsernameButton = "Account";
    UsernameLink = window.location.origin + "/account";
  
  }

  return (<div style={{
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    
    marginTop: "1em",
    marginBottom: "var(--vertical-gap)",
  }}>
    <div className="board" style={{ // Icon
      marginInline: "2em",
      width: "7em",
      height: "7em",
    }}>
      <a href={window.location.origin}><img id="logo" src={window.location.origin + "/bin/imgs/logo.svg"}/></a>
    </div>
    <div className="board" style={{
      height: "3.5em",
      width: "21em",
      marginInline: "2em",
      paddingInline: "1em",
      display: "flex",
      alignItems: "center",
      flexDirection: "row",
    }}>
      <h1>47 FPS forever!!!</h1>
    </div>
    <div className="board" style={{
      height: "3.5em",
      flex: 1,
      marginInline: "2em",
      paddingInline: "0.15em",
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-between",
    }}>
      <div className="headerButton"><a href={window.location.origin + "/articles"}><h3 >Articles</h3></a></div>
      <div className="headerButton"><a href={window.location.origin + "/users"}><h3 >Uses</h3></a></div>
      <div className="headerButton"><a href={window.location.origin + "/projects"}><h3 >Projects</h3></a></div>
      <div className="headerButton"><a href={window.location.origin + "/about"}><h3 >About</h3></a></div>
    </div>
    <div className="board" style={{
      height: "2.9em",
      width: "20em",
      marginInline: "2em",
      padding: "0.3em",
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-between",
    }}>
      <div style={{
        display: "flex",
        alignItems: "center",
        flexDirection: "row",
        flex: 1,
        paddingInline: "1em"
      }}><h3>{Username}</h3></div>
      <div style={{
        display: "flex",
        alignItems: "center",
        flexDirection: "row",
        borderTopRightRadius: "0.7em",
        borderBottomRightRadius: "0.7em",
        backgroundColor: "var(--light-shadow-color)",
        paddingInline: "1em",
      }}><a href={UsernameLink}><h3>{UsernameButton}</h3></a></div>
    </div>
  </div>);
}