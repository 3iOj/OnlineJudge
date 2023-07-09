import React, { useState } from 'react';
import Axios from 'axios';
import { Navigate } from 'react-router-dom';
interface User {
  id: string;
  email: string;
  name: string;
  dob: string;
  moto: string;
}
export interface UserInfo {
  token: string,
  user: User
}

const Login: React.FC = () => {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [loginStatus, setLoginStatus] = useState<boolean>(false);

  Axios.defaults.withCredentials = true;
  const handleLogin = (e :React.FormEvent) => {
    e.preventDefault();
    const data = {
      username,
      password
    };
    console.log(data);
    Axios.post('http://localhost:8080/users/login', data)
      .then(response => {
        console.log(response.data);
        const userData: UserInfo = response.data;
        localStorage.setItem("info", JSON.stringify(userData));
        setLoginStatus(true);
      })
      .catch(error => {
        console.log(error);
      });
  }

  if (loginStatus) {
    return <Navigate to="/" />;
  }
  return (
    <>
      <section className='section-form section-form-small'>
        <form className='form'>
          <h2>Login</h2>
          <div className='form-control'>
            <label htmlFor='username'>Username</label>
            <input
              required
              type='username'
              name='username'
              id='username'
              value = {username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className='form-control'>
            <label htmlFor='password'>Password</label>
            <input
              required
              type='password'
              name='password'
              id='password'
              value = {password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button type='submit' onClick={handleLogin} className='submit-btn'>
            Login
          </button>
        </form>
      </section>
    </>
  )

};

export default Login;
