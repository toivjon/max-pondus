import React from 'react';
import Modal from 'react-bootstrap/Modal';

export default function ProgramWorkoutModal(props) {
  return (
    <Modal {...props} fullscreen>
      <Modal.Header>
        <Modal.Title>Program Workout</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        Hello Program!
      </Modal.Body>
    </Modal>
  );
}