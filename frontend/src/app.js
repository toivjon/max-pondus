import React from 'react';
import { Route, Routes } from 'react-router-dom';
import LoginForm from './login_form';
import RootLayout from './layouts/root_layout';
import HomeLayout from './layouts/home_layout';

// The application main entry point which builds the whole UI structure.
export default function App() {
  return(
    <Routes>
      <Route path="/" element={<RootLayout />}>
        <Route path="home" element={<HomeLayout />}/>
        <Route path="login" element={<LoginForm />}/>
        <Route path="*" element={<p>nothing!</p>}/>
      </Route>
    </Routes>
  );
}