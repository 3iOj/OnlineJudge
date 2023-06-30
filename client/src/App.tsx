import React, { useState } from 'react';
import { BrowserRouter as Routes, Route} from 'react-router-dom';

import Contests from './Components/Contests/Contests.tsx';
import Home from './Pages/Home/Home.tsx';
import Problems from './Components/Problems/Problems.tsx';


const App: React.FC = () => {
  // const [user, setUser] = useState<User | null>(null);

  // return (
  //   <UserContext.Provider value={{ user, setUser }}>
     return (
     <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/contests" element={<Contests/>} />
        <Route path="/problems" element={<Problems/>} />
      </Routes>
     )
      
    // </UserContext.Provider>
  // );
};

export default App;
