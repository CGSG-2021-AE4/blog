import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import LoginForm from "../components/login";
import * as Auth from "../auth";

const infoRowStyle: React.CSSProperties = {
  display: "flex",
  flexDirection: "row",
  justifyContent: "space-between",
  paddingBlock: "0.4em",
  // fontSize: "2em",
};

class UserInfo {
  id: string;
  email: string;
  username: string;
  password: string;
}

function LogoutHandler() {
  Auth.Logout()
  window.location.href = "/"; // Redirect
}

function AccountInfo() {
  let [user, setUser] = useState<UserInfo>({} as UserInfo)

  useEffect(() => {
    fetch("/getUser", {
      method: "POST",
      headers: {
        "Authorization": "Bearer " + Auth.GetToken()
      },
      body: JSON.stringify({
        username: Auth.GetUsername()
      }),
    })
    .then(resp => resp.json())
    .then(u => {
      if (u.err != undefined) {
        // Got error
        window.location.href = "/login"
        return
      }
      setUser(u as UserInfo)
    })
  }, [])
  return (<>
    <div>
      <div style={infoRowStyle}>
        <div>Id:</div>
        <div>{user.id}</div>
      </div>
      <div style={infoRowStyle}>
        <div>Email:</div>
        <div>{user.email}</div>
      </div>
      <div style={infoRowStyle}>
        <div>Username:</div>
        <div>{user.username}</div>
      </div>
      <div style={infoRowStyle}>
        <div>Password:</div>
        <div>{user.password}</div>
      </div>
      <input onClick={LogoutHandler} type="button" style={{
        marginTop: "2em",
      }} value="Logout" />
    </div>
  </>);
}

function Application() {
  return (<>
    <div style={{
      flex: 1,
      display: "flex",
      flexDirection: "column"
    }}>
      <Header/>
      <div style={{
        flex: 1,
        display: "flex",
        flexDirection: "column",
        marginInline: "10em",
        marginBlock: "4em",
      }}>
        <AccountInfo/>
      </div>
      <Footer/>
    </div>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)