import React from 'react';
import Modal from 'react-bootstrap/Modal';

export default function ProgramExerciseModal(props) {
  return (
    <Modal {...props} fullscreen>
      <Modal.Header>
        <Modal.Title>Program Exercise</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        Hello Program!
      </Modal.Body>
    </Modal>
  );
}