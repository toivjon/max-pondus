import React from 'react';
import Button from 'react-bootstrap/Button';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import PropTypes from 'prop-types';

export default function WorkoutTableItem({workout}) {
  return (
    <tr>
      <td>{workout.date}</td>
      <td>{workout.time}</td>
      <td>{workout.duration}</td>
      <td className='p-1'>
        <ButtonGroup size="sm">
          <Button variant="primary">Edit</Button>
          <Button variant="danger">Delete</Button>
        </ButtonGroup>
      </td>
      <td>{workout.notes}</td>
    </tr>
  )
}

WorkoutTableItem.propTypes = {
  workout: PropTypes.object.isRequired
}