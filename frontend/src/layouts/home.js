import React from 'react';
import Button from 'react-bootstrap/Button';
import FreeformWorkoutModal from '../modals/freeform_workout_modal';
import ProgramWorkoutModal from '../modals/program_workout_modal';

export default function HomeLayout() {
  const [fwVisible, setFwVisible] = React.useState(false);
  const [pwVisible, setPwVisible] = React.useState(false);
  return (
    <div>
      <Button onClick={() => setFwVisible(true)}>Freeform Workout</Button>
      <Button onClick={() => setPwVisible(true)}>Program Workout</Button>
      <FreeformWorkoutModal show={fwVisible} onHide={() => setFwVisible(false)} />
      <ProgramWorkoutModal show={pwVisible} onHide={() => setPwVisible(false)} />
    </div>
  )
}