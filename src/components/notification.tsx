import React from "react";
import ReactDOM from "react-dom/client";
import * as Auth from "../auth";
import { useState } from "react";

export const ShowTimeout = 3000; // in ms

export type NotificationType = "ok" | "error" | "status";

class notification {
  type: NotificationType;
  msg: string;
}

class NBoxState {
  ns: notification[];
}

class NBox extends React.Component<{}, NBoxState> {

  constructor() {
    super({})
    this.state = {
      ns: []
    }
  }

  push( n: notification ) {
    console.log("NOTIFICATION: " + n.type + " : " + n.msg)
    this.state.ns.push(n)
    this.setState({ ns: this.state.ns } )
    setTimeout(() => {
      this.state.ns.shift()
      this.setState({ ns: this.state.ns } )        
    }, ShowTimeout)
  }

  render() {
    return(<div style={{
      display: "flex",
      flexDirection: "column",
    }}>
      {this.state.ns.map(n => {
        let color = "black"
        switch (n.type) {
          case "error":
            color = "#b808089c"
            break;
          case "ok":
            color = "#08770896"
            break;
          case "status":
            color = "#0a6aa99c"
            break;
        }

        return (<>
          <div style={{
            backgroundColor: color,
            border: "1px solid var(--main-color)",
            borderRadius: "0.7em",
            paddingBlock: "0.6em",
            paddingInline: "1em",
            width: "20em",
            margin: "0.3em",
          }}>
            {n.msg}
          </div>
        </>);
      })}    
    </div>);
  }
}

export let box: React.RefObject<NBox> = React.createRef();

export let Push = ( n: notification ) => {
  if (box.current != null) {
    box.current.push(n);
  }
}

// Setup
// Since it was included it is supposed to be used
const root = ReactDOM.createRoot(document.getElementById("notificationBox")!);
root.render(<NBox ref={box}/>);
console.log("NOTIFY FYCK")