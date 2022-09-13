import React from 'react';
import Table from 'react-bootstrap/Table';
import WorkoutTableItem from './workout_table_item';
import Workout from '../models/workout';

export default function WorkoutTable() {
  const obj = new Workout();
  return (
    <Table>
      <thead>
        <tr>
          <th>Date</th>
          <th>Time</th>
          <th>Duration</th>
          <th>Actions</th>
          <th>Notes</th>
        </tr>
      </thead>
      <tbody>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
        <WorkoutTableItem workout={obj}/>
      </tbody>
    </Table>
  )
}