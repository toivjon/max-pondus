import React, { Component } from 'react';
import LoginForm from './login_form';
import './app.css';

export default class App extends Component {
  render(){
    return(
      <div className="App">
        <LoginForm />
      </div>
    );
  }
}