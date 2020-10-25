import React from 'react';
import { Redirect, Link } from '@reach/router';

const Error = () => {
    return (
        <div className="App">
            <header className="App-header">
                {/*<Redirect to={"/"} />*/}
                <h1>
                    404
                </h1>
                <p style={{textAlign: "center"}}>
                    User not found. <Link to={"/"} style={{color: "white"}}>Return home.</Link>
                </p>
            </header>
        </div>
    );
}

export default Error;