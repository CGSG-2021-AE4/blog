import React, { useState } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import { useEffect } from "react";
import * as Auth from "../auth";
import * as Notifications from "../components/notification";
import { articleHeader } from "../types";

async function HandleCreate() {
  let resp = await fetch(window.location.origin + "/api/article/create", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + Auth.GetToken()
    },
    body: JSON.stringify({
      title: (document.getElementById("title")! as HTMLInputElement).value,
      id: Auth.GetId(),
    }),
  });
  console.log(resp)
  let res = await resp.json();
  if (res.err != undefined) {
    Notifications.Push({type: "error", msg: "Error: " + res.err});
    return;    
  }
  Notifications.Push({type: "ok", msg: "Success creation"});
  setTimeout(() => {
    window.location.href = window.location.origin + "/article/edit?id=" + res.id;
  }, 1000)
}

function HandleCancel() {
  window.location.href = window.location.origin;
}

function ArticleCreationForm() {
  useEffect(()=>{
    if (!Auth.IsAuthorized()) {
      window.location.href = window.location.origin + "/login";
      return;
    }
  }, [])
  
  const buttomStyle: React.CSSProperties = {
    marginInline: "0.7em",
    paddingInline: "1em",
  }

  return (<>
    <div style={{flex: 1, margin: "1em", display: "flex", flexDirection: "column"}}>
      <h2>Title: <input id="title" type="text"/></h2>
      <div style={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "end",
        margin: "1em"
      }}>
        <input type="button" value="Create" onClick={HandleCreate} style={buttomStyle}/>
        <input type="button" value="Cancel" onClick={HandleCancel} style={buttomStyle}/>
      </div>
    </div>
  </>);
}

function Application() {
  return (<>
    <Header/>
      <ArticleCreationForm />
    <Footer/>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)