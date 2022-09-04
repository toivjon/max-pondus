import React from 'react';
import Button from 'react-bootstrap/Button';
import FloatingLabel from 'react-bootstrap/FloatingLabel';
import Form from 'react-bootstrap/Form';

export default function LoginForm() {
  const [username, setUsername] = React.useState("");
  const [password, setPassword] = React.useState("");

  function submit(event) {
    // TODO Implement the logic required for the authentication.
    console.log("username: " + username + " password: " + password);
    event.preventDefault();
  }

  function validate() {
    return username.length > 0 && password.length > 0;
  }

  return(
    <Form onSubmit={submit} id="login-form">
      <Form.Group className="mb-3">
        <FloatingLabel className="mb-3" controlId="email" label="Username">
          <Form.Control
            autoFocus
            type="email"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            autoComplete="username"
          />
        </FloatingLabel>
      </Form.Group>
      <Form.Group className="mb-3">
        <FloatingLabel className="mb-3" controlId="password" label="Password">
          <Form.Control
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            autoComplete="current-password"
          />
        </FloatingLabel>
      </Form.Group>
      <Button variant="primary" type="submit" block="true" disabled={!validate()}>Login</Button>
    </Form>
  )
}