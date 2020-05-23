import React from 'react';
import { Router } from '@reach/router'

import Index from './views/index'
import Profile from './views/profile'
import Error from './views/error'
import './App.css';

function App() {
  return (
    <Router>
      <Index path="/" />
      <Profile path="/:username" />
      <Error path="/404" default />
    </Router>
  );
}

export default App;