import React from 'react';
import logo from '../logo.svg';
import {getUser} from "../static/people";

const Index = () => {
    console.log(getUser('daviskeene'));
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>
                    Davis is cool B)
                </p>
                <a
                    className="App-link"
                    href="https://reactjs.org"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Learn React
                </a>
            </header>
        </div>
    );
}

export default Index;