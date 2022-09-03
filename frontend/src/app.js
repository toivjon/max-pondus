import React, { Component } from 'react';
import LoginForm from './login_form';
import './app.css';

// The application main entry point which builds the whole UI structure.
export default class App extends Component {
  render(){
    return(
      <div className="App">
        <LoginForm />
      </div>
    );
  }
}