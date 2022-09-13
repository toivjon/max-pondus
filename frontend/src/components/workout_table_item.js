import React from 'react';
import Button from 'react-bootstrap/Button';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import PropTypes from 'prop-types';

export default function WorkoutTableItem({workout}) {
  const dateOptions = { year: "numeric", month: "2-digit", day: "2-digit", weekday: "short"};
  const timeOptions = { hour: "2-digit", minute: "2-digit"};
  return (
    <tr>
      <td>{workout.start.toLocaleDateString(undefined, dateOptions)}</td>
      <td>{workout.start.toLocaleTimeString(undefined, timeOptions)}</td>
      <td>{workout.end - workout.start}</td>
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