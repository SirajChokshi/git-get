import React, {useState, useEffect, KeyboardEvent} from "react";
import {Link, navigate} from '@reach/router'

const Index = () => {

    const [query, updateQuery] = useState<string>("")

    const _handleKeyDown = (e: KeyboardEvent) => {
        if (e.key === 'Enter') {
            navigate(`/${query}`)
        }
      }

    return (
        <div className="App">
            <header className="search-header">
                <h1>Search for a GitHub user or organization:</h1>
                <div id="search-wrapper">
                    <input onChange={(e) => updateQuery(e.target.value)} onKeyDown={_handleKeyDown} />
                    <Link to={`/${query}`}>Submit &rarr;</Link>
                </div>
            </header>
        </div>
    );
}

export default Index;