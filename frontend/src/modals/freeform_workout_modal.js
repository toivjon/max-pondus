import React from 'react';
import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import PropTypes from 'prop-types';

// TODO Build a new workout instance and assign it as a state.
// TODO Replace placeholders with the values from the workout state.

export default function FreeformWorkoutModal(props) {
  const onSave = () => {
    // TODO Execute the actual saving of the form contents.
    props.onHide();
  }
  const onAddExercise = () => {
    // TODO Add the selected exercise to the exercises table.
    // TODO Activate the selected exercise.
  }
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
          <InputGroup className='p-2'>
            <Form.Select>
              <option>Select an exercise</option>
              <option>Foo</option>
              <option>Bar</option>
              <option>Baz</option>
            </Form.Select>
            <Button onClick={onAddExercise}>Add Exercise</Button>
          </InputGroup>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button onClick={onSave}>Save</Button>
      </Modal.Footer>
    </Modal>
  );
}

FreeformWorkoutModal.propTypes = {
  onHide: PropTypes.func.isRequired
}