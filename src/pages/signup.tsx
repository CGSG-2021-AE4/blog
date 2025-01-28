import React from "react";
import ReactDOM from "react-dom/client";
import Header from "../components/header";
import Footer from "../components/footer";
import SignupForm from "../components/signup";

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
        <SignupForm/>
      </div>
      <Footer/>
    </div>
  </>);
}
const root = ReactDOM.createRoot(document.getElementById("app")!)
root.render(<Application />)