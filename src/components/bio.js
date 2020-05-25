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
                    <table>
                        {
                            props.user.location && (
                                <tr>
                                    <td>Location</td>
                                    <td>{props.user.location}</td>
                                </tr>
                            )
                        }
                        {
                            props.user.company && (
                                <tr>
                                    <td>Company</td>
                                    <td>
                                        {
                                            reactStringReplace(props.user.company, /\B@([\w-]+)/gm, (match, i) => (
                                                <Link key={i} to={`/${match}`} >@{match}</Link>
                                            ))
                                        }
                                    </td>
                                </tr>
                            )
                        }
                        {
                            props.user.blog && (
                                <tr>
                                    <td>Website</td>
                                    <td>
                                        <a href={props.user.blog} target={"_blank"} rel={"noopener"}>{props.user.blog}</a>
                                    </td>
                                </tr>
                            )
                        }

                    </table>
                </div>
                <div id="extra-bio-info">
                    { props.user.type === "Organization"
                        ?
                           <></>
                        :
                        <>
                            <div className="badge">
                                <div className="num">{props.user.followers}</div>
                                Following
                            </div>
                            <div className="badge">
                                <div className="num">{props.user.following}</div>
                                Following
                            </div>
                        </>
                    }
                </div>
        </header>
    );
}

export default Bio;