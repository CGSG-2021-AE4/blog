import React, { useState } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import { useEffect } from "react";
import * as Auth from "../auth";
import * as Notifications from "../components/notification";
import { articleHeader } from "../types";

let articleId: string = (new URLSearchParams(window.location.search)).get("id")!;

async function HandleSave() {
  let resp = await fetch(window.location.origin + "/api/article/edit", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + Auth.GetToken()
    },
    body: JSON.stringify({
      id: articleId,
      title: (document.getElementById("title")! as HTMLInputElement).value,
      content: (document.getElementById("content")! as HTMLInputElement).value,
    }),
  });
  let res = await resp.json();
  if (res.err != undefined) {
    Notifications.Push({type: "error", msg: "Error: " + res.err});
    return;    
  }
  console.log(res)
  Notifications.Push({type: "ok", msg: "Success: " + res.msg});
  setTimeout(() => {
    window.location.href = window.location.origin + "/article?id=" + articleId;
  }, 1000)
}


async function HandleDelete() {
  let resp = await fetch(window.location.origin + "/api/article/delete", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + Auth.GetToken()
    },
    body: JSON.stringify({
      id: articleId,
    }),
  });
  let res = await resp.json();
  if (res.err != undefined) {
    Notifications.Push({type: "error", msg: "Error: " + res.err});
    return;    
  }
  console.log(res)
  Notifications.Push({type: "ok", msg: "Success: " + res.msg});
  setTimeout(() => {
    window.location.href = window.location.origin;
  }, 1000)
}

async function HandleCancel() {
  window.location.href = window.location.origin;
}

function Article() {
  let [header, setHeader] = useState<articleHeader>({} as articleHeader)
  let [content, setContent] = useState<string>("")

  useEffect(()=>{
    let fetchHeader = async () => {
      let resp = await fetch(window.location.origin + "/api/article/header?id=" + articleId);
      let res = await resp.json();
      if (res.err != undefined) {
        // Got error
        Notifications.Push({type: "error", msg: "Error: " + res.err});
        return
      }
      setHeader(res)
      console.log(res);
    }
    fetchHeader()
  }, [])

  useEffect(() => {
    (document.getElementById("title")! as HTMLInputElement).value = header.title;

    let fetchContent = async () => {
      if (header.contentId == undefined) {
        return
      }
      let resp = await fetch(window.location.origin + "/api/article/content?id=" + header.contentId);
      let res = await resp.json();
      if (res.err != undefined) {
        // Got error
        Notifications.Push({type: "error", msg: "Error: " + res.err});
        return
      }
      setContent(res.text)
    }
    fetchContent()
  }, [header])
  
  useEffect(() => {
    (document.getElementById("content")! as HTMLInputElement).value = content;
  }, [content])
  
  const buttomStyle: React.CSSProperties = {
    marginInline: "0.7em",
    paddingInline: "1em",
  }

  return (<>
    <div style={{flex: 1, margin: "1em", display: "flex", flexDirection: "column"}}>
      <h2>Title: <input id="title" type="text"/></h2>
      <p><textarea id="content" rows={10} cols={115}></textarea></p>
      <div style={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "end",
        margin: "1em"
      }}>
        <input type="button" value="Save" onClick={HandleSave} style={buttomStyle}/>
        <input type="button" value="Cancel" onClick={HandleCancel} style={buttomStyle}/>
        <input type="button" value="Delete" onClick={HandleDelete} style={buttomStyle}/>
      </div>
    </div>
  </>);
}

function Application() {
  return (<>
    <Header/>
      <Article />
    <Footer/>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)