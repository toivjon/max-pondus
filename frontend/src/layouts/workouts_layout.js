import React from 'react';
import Table from 'react-bootstrap/Table';
import Button from 'react-bootstrap/Button';

export default function WorkoutsLayout() {
  return (
    <Table>
    <thead>
      <tr>
        <th></th>
        <th>Date</th>
        <th>Time</th>
        <th>Duration</th>
        <th>Notes</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td className='p-1'>
          <Button variant="primary" size="sm">Edit</Button>
          <Button variant="danger" size="sm">Delete</Button>
        </td>
        <td>10.09.2022</td>
        <td>09:30</td>
        <td>1 hour 30 min</td>
        <td>Foo seems to be bar here?</td>
      </tr>
      <tr>
       <td className='p-1'>
          <Button variant="primary" size="sm">Edit</Button>
          <Button variant="danger" size="sm">Delete</Button>
        </td>
        <td>01.09.2022</td>
        <td>12:30</td>
        <td>1 hour</td>
        <td></td>
      </tr>
      <tr>
        <td className='p-1'>
          <Button variant="primary" size="sm">Edit</Button>
          <Button variant="danger" size="sm">Delete</Button>
        </td>
        <td>30.08.2022</td>
        <td>10:30</td>
        <td>1 hour 15 min</td>
        <td>Foo seems to be bar here?</td>
      </tr>
      <tr>
        <td className='p-1'>
          <Button variant="primary" size="sm">Edit</Button>
          <Button variant="danger" size="sm">Delete</Button>
        </td>
        <td>20.08.2022</td>
        <td>09:00</td>
        <td>45 min</td>
        <td>Foo goo foo here?</td>
      </tr>
      <tr>
        <td className='p-1'>
          <Button variant="primary" size="sm">Edit</Button>
          <Button variant="danger" size="sm">Delete</Button>
        </td>
        <td>13.05.2022</td>
        <td>10:45</td>
        <td>1 hour</td>
        <td>Foo bar bar foo?</td>
      </tr>
    </tbody>
    </Table>
  )
}