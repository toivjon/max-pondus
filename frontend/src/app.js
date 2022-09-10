import React from 'react';
import { Route, Routes } from 'react-router-dom';
import LoginForm from './login_form';
import RootLayout from './layouts/root_layout';
import HomeLayout from './layouts/home_layout';
import WorkoutsLayout from './layouts/workouts_layout';
import ProgramsLayout from './layouts/programs_layout';

// The application main entry point which builds the whole UI structure.
export default function App() {
  return(
    <Routes>
      <Route path="/" element={<RootLayout/>}>
        <Route index element={<HomeLayout/>}/>
        <Route path="workouts" element={<WorkoutsLayout/>}/>
        <Route path="programs" element={<ProgramsLayout/>}/>
        <Route path="login" element={<LoginForm/>}/>
      </Route>
    </Routes>
  );
}