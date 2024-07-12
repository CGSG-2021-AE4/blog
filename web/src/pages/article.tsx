import React, { useState } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import { useEffect } from "react";
import * as Notifications from "../components/notification";

class article {
  id: string
	title: string
  text: string
}

const rowDivStyle = {
  
} as React.CSSProperties

function Article() {
  let [article, setArticle] = useState<article>({} as article)

  useEffect(()=>{
    let fetchArticle = async () => {
      let resp = await fetch("/api/article?id=" + (new URLSearchParams(window.location.search)).get("id"));
      let res = await resp.json();
      if (res.err != undefined) {
        // Got error
        console.log(res.err)
        return
      }
      setArticle(res as article)
    }

    fetchArticle();
  }, [])
  
  return (<>
    <div style={{flex: 1, margin: "1em"}}>
      <h1>{article.title}</h1>
      <div>{article.text}</div>
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