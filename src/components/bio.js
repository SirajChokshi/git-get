import React, {useState, useEffect} from "react";
import reactStringReplace from 'react-string-replace'
import {Link} from '@reach/router'
import './bio.css'

const Bio = (props) => {

    return (
        <header id={"bio"}>
                <img src={props.user.avatar_url} id={"avatar"} />
                <div id={"bio-info"}>
                    <h2>{props.user.name}</h2>
                    <h3>@{props.user.login}</h3>
                    <p>
                        {
                            reactStringReplace(props.user.bio, /\B@([\w-]+)/gm, (match, i) => (
                                <Link key={i} to={`/${match}`} >@{match}</Link>
                            ))
                        }
                    </p>
                </div>
                <div id="extra-bio-info">
                    X Followers
                </div>
        </header>
    );
}

export default Bio;