import React from 'react';
import { Route, Routes } from 'react-router-dom';
import LoginForm from './login_form';
import RootLayout from './layouts/root_layout';
import HomeLayout from './layouts/home_layout';
import ExercisesLayout from './layouts/exercises_layout';
import ProgramsLayout from './layouts/programs_layout';

// The application main entry point which builds the whole UI structure.
export default function App() {
  return(
    <Routes>
      <Route path="/" element={<RootLayout/>}>
        <Route index element={<HomeLayout/>}/>
        <Route path="exercises" element={<ExercisesLayout/>}/>
        <Route path="programs" element={<ProgramsLayout/>}/>
        <Route path="login" element={<LoginForm/>}/>
      </Route>
    </Routes>
  );
}