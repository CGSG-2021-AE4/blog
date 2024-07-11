import React from "react";
import ReactDOM from "react-dom/client";

export default function Footer() {
  return (<div style={{
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",

    marginTop: "var(--vertical-gap)",
    marginInline: "4em",
  }}>
    <div>CGSG</div>
    <div>Copyright 2024 AE4</div>
  </div>);
}