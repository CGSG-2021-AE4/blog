import React, { useState } from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import { useEffect } from "react";
import * as Notifications from "../components/notification";
import { articleHeader } from "../types";

function ArticleTitle( props: {a: articleHeader} ) {
  return (<>
    <div className="board" style={{
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-between",
      marginBlock: "1em",
    }}>
      <div style={{
        flex: 1,
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        paddingInline: "1em",
      }}><h2>{props.a.title}</h2></div>
      <div style={{
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        paddingInline: "1em",
      }}><h3>{props.a.authorUsername}</h3></div>
      <div style={{
        margin: "0.3em",
        display: "flex",
        alignItems: "center",
        flexDirection: "row",
        borderTopRightRadius: "0.7em",
        borderBottomRightRadius: "0.7em",
        backgroundColor: "var(--light-shadow-color)",
        paddingInline: "1em",
      }}><a href={window.location.origin + "/article?id=" + props.a.id}>Read</a></div>
    </div>
  </>);
}

function Articles() {
  let [articles, setArticles] = useState<articleHeader[]>([])

  useEffect(()=>{
    let fetchArticles = async () => {
      let resp = await fetch(window.location.origin + "/api/article/list");
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
  
  return (<>{articles.map(a=>(<ArticleTitle a={a}/>))}</>);
}

function Application() {
  return (<>
    <Header/>
      <div style={{
        width: "60%",
        alignSelf: "center",
      }}>
        <Articles />
      </div>
    <Footer/>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)