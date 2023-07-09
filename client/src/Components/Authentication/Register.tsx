import React, { useState } from 'react';
import DatePicker from "react-datepicker";
import 'react-datepicker/dist/react-datepicker.css';
import Axios from 'axios';
import { Navigate } from 'react-router-dom';

const Register:React.FC = () => {
  const [name, setName] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [dob, setDob] = useState<Date|null>();
  const [registrationSuccess, setRegistrationSuccess] = useState<boolean>(false);

  Axios.defaults.withCredentials = true;
  const handleRegistration = (e: React.FormEvent) => {
    e.preventDefault();
    const data = {
      name,
      username,
      password,
      email,
      dob,
    };
    Axios.post('http://localhost:8080/users/register', data)
      .then(response => {
        console.log(response.data);
        setRegistrationSuccess(true);
        // alert(`Succesfully registered`);
      })
      .catch(error => {
        console.log(error);
      });
  }
  
  if (registrationSuccess) {
    return <Navigate to="/login" />;
  }
  return (
    <>
      <section className='section-form section-form-small'>
        <form className='form'>
          <h2>Register</h2>
          <div className='form-control'>
            <label htmlFor='name'>Name</label>
            <input
              required
              type='name'
              name='name'
              id='name'
              value = {name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className='form-control'>
            <label htmlFor='username'>Username</label>
            <input
              required
              type='username'
              name='username'
              id='username'
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className='form-control'>
            <label htmlFor='email'>Email</label>
            <input
              required
              type='email'
              name='email'
              id='email'
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div className='form-control'>
            <label htmlFor='password'>Password</label>
            <input
              type='password'
              name='password'
              id='password'
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className='form-control'>
            <label htmlFor='dob'>Date of birth</label>
            <DatePicker 
              required
              selected={dob}
              onChange={(dob) => setDob(dob)}
            />
          </div>
          <button type='submit' onClick={handleRegistration} className='submit-btn'>
            Register
          </button>
        </form>
      </section>
    </>
  )

};

export default Register;
