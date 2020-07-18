import React, {useState, useEffect} from "react";
import {Link, navigate} from '@reach/router'
import logo from '../logo.svg';

const Index = () => {

    const [query, updateQuery] = useState("")

    const _handleKeyDown = (e) => {
        if (e.key === 'Enter') {
            navigate(`/${query}`)
        }
      }

    return (
        <div className="App">
            <header className="app-header">
                <h1>Search for a GitHub user or organization:</h1>
                <div id="search-wrapper">
                    <input style={{fontSize: "24px"}} onChange={(e) => updateQuery(e.target.value)} onKeyDown={_handleKeyDown} />
                    <Link to={`/${query}`}>Submit &rarr;</Link>
                </div>
            </header>
        </div>
    );
}

export default Index;