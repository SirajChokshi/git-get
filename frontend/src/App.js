import React from 'react';
import { Router } from '@reach/router'

import Index from './views/index'
import Profile from './views/profile'
import Error from './views/error'
import Nav from './components/Nav'
import RateLimit from './components/rateLimit'
import './App.scss';
import Footer from './components/footer';

function App() {
  return (
    <>
      <div className="container" style={{paddingTop: "5vh"}}>
        <Nav />
        <Router>
          <Index path="/" />
          <Index path="/about" />
          <Index path="/page" />
          <Profile path="/:username" />
          <Error path="/404" default />
        </Router>
        {/* <RateLimit /> */}
      </div>
      <Footer />
    </>
  );
}

export default App;