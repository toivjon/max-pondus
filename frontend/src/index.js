import React from 'react';
import ReactDOM from 'react-dom/client'
import App from './app';

// Import bootstrap styles to get UI components look pretty.
import 'bootstrap/dist/css/bootstrap.min.css';

// Build and render the React application root structure.
const rootElement = document.getElementById("root");
const root = ReactDOM.createRoot(rootElement);
root.render(<React.StrictMode><App /></React.StrictMode>);