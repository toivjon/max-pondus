import React from "react";

class LoginForm extends React.Component {
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(event) {
    const username = event.target.username.value;
    const password = event.target.password.value;
    console.log("Login username=" + username + " password=" + password);
    event.preventDefault();
  }

  render() {
    return(
      <form className="login-form" onSubmit={this.handleSubmit}>
        <label htmlFor="username">Username</label>
        <input id="username" name="username" type="email"  placeholder="User Email" />
        <label htmlFor="password">Password</label>
        <input id="password" name="password" type="password" placeholder="Password" autoComplete="on" />
        <input name="submit" type="submit" value="Login" />
      </form>
    )
  }
}

export default LoginForm;