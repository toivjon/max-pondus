import React from 'react';
import Table from 'react-bootstrap/Table';
import WorkoutTableItem from './workout_table_item';
import Workout from '../models/workout';

export default function WorkoutTable() {
  const [workouts, setWorkouts] = React.useState([]);
  React.useEffect(() =>{
    const workouts = [];
    for (let i = 0; i < 100; i++) {
      const workout = new Workout();
      workout.id = i;
      workouts.push(workout);
    }
    setWorkouts(workouts);
  }, []);
  return (
    <Table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Date</th>
          <th>Time</th>
          <th>Duration</th>
          <th>Actions</th>
          <th>Notes</th>
        </tr>
      </thead>
      <tbody>
        {workouts.map((workout) => {
          return (
            <WorkoutTableItem key={workout.id} workout={workout}/>
          )
        })}
      </tbody>
    </Table>
  )
}