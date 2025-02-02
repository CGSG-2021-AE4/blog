import React, { useState } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import { useEffect } from "react";
import * as Auth from "../auth";
import * as Notifications from "../components/notification";
import { articleHeader } from "../types";

let articleId: string = (new URLSearchParams(window.location.search)).get("id")!;

function Article() {
  let [header, setHeader] = useState<articleHeader>({} as articleHeader);
  let [editButton, setEditButton] = useState<React.JSX.Element>(<></>);
  
  useEffect(()=>{
    let fetchHeader = async () => {
      let resp = await fetch(window.location.origin + "/api/article/header?id=" + articleId);
      let res = await resp.json();
      if (res.err != undefined) {
        // Got error
        Notifications.Push({type: "error", msg: "Error: " + res.err});
        return;
      }
      setHeader(res);
      console.log(res);
    }
    fetchHeader();
  }, [])

  useEffect(() => {
    let fetchContent = async () => {
      if (header.contentId == undefined) {
        return;
      }
      let resp = await fetch(window.location.origin + "/api/article/contentHTML?id=" + header.contentId);
      let res = await resp.json();
      if (res.err != undefined) {
        // Got error
        Notifications.Push({type: "error", msg: "Error: " + res.err});
        return
      }
      (document.getElementById("articleContent")!).innerHTML = res.text;
      if (header.authorId == Auth.GetId()) {
        setEditButton(<>
          <div>
            <input type="button" onClick={() => {
              window.location.href = window.location.origin + "/article/edit?id=" + header.id;
            }} value="Edit"/>
          </div>
        </>);
      }
    }
    fetchContent()
  }, [header])

  return (<>
    <div style={{flex: 1, margin: "1em"}}>
      <div style={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        marginBottom: "1em",
      }}>
        <h1 style={{flex: 1, margin: 0}}>{header.title}</h1>
        <a href={window.location.origin + "user?id" + header.authorId}><h3>{header.authorUsername}</h3></a>
      </div>
      <hr/>
      <div id="articleContent"></div>
      {editButton}
    </div>
  </>);
}

function Application() {
  return (<>
    <Header/>
      <div className="board" style={{
        width: "60%",
        alignSelf: "center",
      }}>
        <Article />
      </div>
    <Footer/>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)