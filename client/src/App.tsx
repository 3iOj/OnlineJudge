import React, { useState, useRef, useEffect } from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import { FaBars } from 'react-icons/fa'
import logo from "./3iojlongo.png";
import Contests from './Components/Contests/Contests.tsx';
import Home from './Pages/Home/Home.tsx';
import Problems from './Components/Problems/Problems.tsx';
import Login from './Components/Authentication/Login.tsx';
import Register from './Components/Authentication/Register.tsx';
import Profile from './Components/User/Profile.tsx';

export interface IApplicationProps { }

const App: React.FC<IApplicationProps> = (props) => {
  // const [user, setUser] = useState<User | null>(null);
  const [showLinks, setShowLinks] = useState(false);
  const [loginStatus, setLoginStatus] = useState(false);
  const linksContainerRef: any = useRef(null);
  const linksRef: any = useRef(null);
  useEffect(()=>{
    if(localStorage.getItem('info')){
      setLoginStatus(true);
    }
  },[]);
  
  useEffect(() => {
    const linksHeight: any = linksRef.current.getBoundingClientRect().height;
    if (showLinks) {
      linksContainerRef.current.style.height = `${linksHeight}px`;
    } else {
      linksContainerRef.current.style.height = "0px";
    }
  })
  // return (
  //   <UserContext.Provider value={{ user, setUser }}>
  return (
    <div>
      <nav>
        <div className="nav-center">
          <div className="nav-header">
            <img src={logo} className="logo" alt="logo" />
            <button
              className="nav-toggle"
              onClick={() => setShowLinks(!showLinks)}
            >
              <FaBars />
            </button>
          </div>
          <div className='links-container' ref={linksContainerRef}>
            <ul className="links" ref={linksRef}>
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="contests">Contests</Link>
              </li>
              <li>
                <Link to="problems">Problems</Link>
              </li>
              {loginStatus &&
                <li>
                  <Link to="profile">Profile</Link>
                </li>
              }
              {!loginStatus && 
              ( <>
                  <li className='authentication'>
                    <Link to="login">Login</Link>
                  </li>
                  <li className='authentication'>
                    <Link to="register">Register</Link>
                  </li>
              </>)}
            </ul>
          </div>
        </div>

      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="contests" element={<Contests />} />
        <Route path="problems" element={<Problems />} />
        <Route path="login" element={<Login />} />
        <Route path="register" element={<Register />} />
      </Routes>
    </div>
  )

  // </UserContext.Provider>
  // );
};

export default App;
