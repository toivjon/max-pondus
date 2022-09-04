import React from 'react';
import { Navbar, Nav } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';

// The root navigation bar for the application.
export default function RootNavbar() {
  return (
    <Navbar bg="light" className="p-2">
      <Navbar.Brand>MaX Pondus</Navbar.Brand>
      <Nav defaultActiveKey="/home">
        <LinkContainer to="/home">
          <Nav.Link>Home</Nav.Link>
        </LinkContainer>
        <LinkContainer to="/exercises">
          <Nav.Link>Exercises</Nav.Link>
        </LinkContainer>
        <LinkContainer to="/programs">
          <Nav.Link>Programs</Nav.Link>
        </LinkContainer>
      </Nav>
    </Navbar>
  );
}