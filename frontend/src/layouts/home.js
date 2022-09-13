import React from 'react';
import Button from 'react-bootstrap/Button';
import FreeformExerciseModal from '../modals/freeform_exercise_modal';
import ProgramExerciseModal from '../modals/program_exercise_modal';

export default function HomeLayout() {
  const [feVisible, setFeVisible] = React.useState(false);
  const [peVisible, setPeVisible] = React.useState(false);
  return (
    <div>
      <Button onClick={() => setFeVisible(true)}>Freeform Exercise</Button>
      <Button onClick={() => setPeVisible(true)}>Program Exercise</Button>
      <FreeformExerciseModal show={feVisible} onHide={() => setFeVisible(false)} />
      <ProgramExerciseModal show={peVisible} onHide={() => setPeVisible(false)} />
    </div>
  )
}