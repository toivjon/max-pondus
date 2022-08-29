import React from 'react';
import { createRoot } from 'react-dom/client'

import App from "./app.js";

// Build and render the React application root structure.
const rootElement = document.getElementById("root");
const root = createRoot(rootElement);
root.render(<React.StrictMode><App /></React.StrictMode>);