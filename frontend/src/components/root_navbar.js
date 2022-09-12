import React from 'react';
import { Navbar, Nav } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';

// The root navigation bar for the application.
export default function RootNavbar() {
  return (
    <Navbar bg="dark" variant="dark" className="p-2">
      <Navbar.Brand>MaX Pondus</Navbar.Brand>
      <Nav defaultActiveKey="/">
        <LinkContainer to="/">
          <Nav.Link>Home</Nav.Link>
        </LinkContainer>
        <LinkContainer to="/workouts">
          <Nav.Link>Workouts</Nav.Link>
        </LinkContainer>
        <LinkContainer to="/programs">
          <Nav.Link>Programs</Nav.Link>
        </LinkContainer>
      </Nav>
    </Navbar>
  );
}