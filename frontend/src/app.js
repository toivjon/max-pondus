import React from 'react';
import LoginForm from './login_form';
import './app.css';

// The application main entry point which builds the whole UI structure.
export default function App() {
  return(
    <div className="App">
      <LoginForm />
    </div>
  );
}