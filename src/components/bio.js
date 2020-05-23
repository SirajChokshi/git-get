import React, {useState, useEffect} from "react";

const Bio = (props) => {

    return (
        <div className="App">
            <header className="App-header">
                <img src={props.avatar} />
                <h2>{props.username}</h2>
                <p>{props.bio}</p>
            </header>
        </div>
    );
}

export default Bio;