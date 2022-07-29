import React, { Component} from "react";
import "./app.css";

class App extends Component{
  render(){
    return(
      <div className="App">
        <form className="form">
          <div className="inputGroup">
            <label htmlFor="email">Email</label>
            <input type="email" name="email" placeholder="name@email.com"></input>
          </div>
          <div className="inputGroup">
            <label htmlFor="password">Password</label>
            <input type="password" name="password"></input>
          </div>
          <button className="primary">Login</button>
        </form>
      </div>
    );
  }
}

export default App;