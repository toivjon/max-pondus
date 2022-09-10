import React from 'react';
import Table from 'react-bootstrap/Table';
import WorkoutTableItem from './workout_table_item';

export default function WorkoutTable() {
  const obj = {
    date: "10.09.2022",
    time: "09:30",
    duration: "1 hour 30 min",
    notes: "Foo seems to be bar here?"
  };
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