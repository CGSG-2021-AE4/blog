import React from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import LoginForm from "../components/login";

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
        alignItems: "center",
        justifyContent: "center",
      }}>
        <LoginForm/>
      </div>
      <Footer/>
    </div>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)