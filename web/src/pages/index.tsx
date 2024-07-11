import React, { useState } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import { useEffect } from "react";
import * as Notifications from "../components/notification";

class articleHeader {
  id: string
	title: string
}

class article {
  id: string
	title: string
  text: string
}

const rowDivStyle = {
  
} as React.CSSProperties

function ArticleTitle( props: {a: articleHeader} ) {
  return (<>
    <div style={{
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-between",
        
      padding: "1.5em",
      border: "1px solid var(--main-color)",
      borderRadius: "1em",
    }}>
      <div>Title: {props.a.title}</div>
      <a href={"/article?id=" + props.a.id}>Read</a>
    </div>
  </>);
}

function Articles() {
  let [articles, setArticles] = useState<articleHeader[]>([])

  useEffect(()=>{
    let fetchArticles = async () => {
      let resp = await fetch("/api/article/list");
      let res = await resp.json();
      if (res.err != undefined) {
        // Got error
        console.log(res.err)
        return
      }
      setArticles(res as articleHeader[])
    }

    fetchArticles();
  }, [])
  
  return (<>
    <div style={{flex: 1}}>
      {articles.map(a=>(<ArticleTitle a={a}/>))}
    </div>
  </>);
}

function Application() {
  return (<>
    <Header/>
      <Articles />
    <Footer/>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)