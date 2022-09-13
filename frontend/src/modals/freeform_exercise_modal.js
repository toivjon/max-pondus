import React from 'react';
import Modal from 'react-bootstrap/Modal';

export default function FreeformExerciseModal(props) {
  return (
    <Modal {...props} fullscreen>
      <Modal.Header>
        <Modal.Title>Freeform Exercise</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        Hello Freeform!
      </Modal.Body>
    </Modal>
  );
}