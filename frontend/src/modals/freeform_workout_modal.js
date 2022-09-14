import React from 'react';
import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';
import Modal from 'react-bootstrap/Modal';

// TODO Build a new workout instance and assign it as a state.
// TODO Replace placeholders with the values from the workout state.

export default function FreeformWorkoutModal(props) {
  return (
    <Modal {...props} fullscreen>
      <Modal.Header>
        <Modal.Title>Freeform Workout</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form>
          <InputGroup className='p-2'>
            <InputGroup.Text>Date</InputGroup.Text>
            <Form.Control type="date" value={"2022-10-10"}/>
          </InputGroup>
          <InputGroup className='p-2'>
            <InputGroup.Text>Start</InputGroup.Text>
            <Form.Control type="time" value={"09:10"}/>
            <InputGroup.Text>End</InputGroup.Text>
            <Form.Control type="time" value={"10:10"}/>
          </InputGroup>
        </Form>
      </Modal.Body>
    </Modal>
  );
}