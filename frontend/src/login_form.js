import React from "react";

class LoginForm extends React.Component {

  handleSubmit(event) {
    const username = event.target.username.value;
    const password = event.target.password.value;
    console.log("Login username=" + username + " password=" + password);
  }

  render() {
    return(
      <form className="login-form" onSubmit={this.handleSubmit}>
          <input name="username" type="email"  placeholder="User Email" />
          <br />
          <input name="password" type="password" placeholder="Password" autoComplete="on" />
          <br />
          <button className="primary">Login</button>
      </form>
    )
  }
}

export default LoginForm;