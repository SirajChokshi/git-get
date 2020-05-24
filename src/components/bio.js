import React, {useState, useEffect} from "react";
import reactStringReplace from 'react-string-replace'
import {Link} from '@reach/router'

const Bio = (props) => {

    return (
        <div className="App">
            <header className="App-header">
                <img src={props.avatar} />
                <h2>{props.username}</h2>
                <p>
                    {
                        reactStringReplace(props.bio, /\B@([\w-]+)/gm, (match, i) => (
                            <Link key={i} to={`/${match}`} >@{match}</Link>
                        ))
                    }
                </p>
            </header>
        </div>
    );
}

export default Bio;