import React, {useState, useEffect} from "react";
import {Link} from '@reach/router'
import logo from '../logo.svg';

const Index = () => {

    const [query, updateQuery] = useState("")

    return (
        <div className="App">
            <header className="App-header">
                <h1>Search for a GitHub user or organization:</h1>
                <input style={{fontSize: "24px"}} onChange={(e) => updateQuery(e.target.value)} />
                <Link to={`/${query}`}>Submit</Link>
            </header>
        </div>
    );
}

export default Index;