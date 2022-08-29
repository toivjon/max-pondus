import React, { Component} from "react";
import "./login_form"
import "./app.css";
import LoginForm from "./login_form";

class App extends Component{
  render(){
    return(
      <div className="App">
        <h1>Sign In</h1>
        <LoginForm />
      </div>
    );
  }
}

export default App;