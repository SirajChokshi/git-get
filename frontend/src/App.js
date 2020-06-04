import React from 'react';
import { Router } from '@reach/router'

import Index from './views/index'
import Profile from './views/profile'
import Error from './views/error'
import RateLimit from './components/rateLimit'
import './App.css';

function App() {
  return (
      <div className="container">
        <Router>
          <Index path="/" />
          <Profile path="/:username" />
          <Error path="/404" default />
        </Router>
        <RateLimit />
      </div>
  );
}

export default App;