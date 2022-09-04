import React from 'react';
import { Outlet } from 'react-router-dom';
import RootNavbar from '../components/root_navbar';

// The root layout for the whole application.
export default function RootLayout() {
  return (
    <div>
      <RootNavbar />
      <main>
        <Outlet />
      </main>
    </div>
  );
}