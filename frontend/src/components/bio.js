import React, {useState, useEffect} from "react";
import reactStringReplace from 'react-string-replace'
import {Link} from '@reach/router'
import './bio.css'

const Bio = (props) => {

    const getAccountAge = () => {
        return Math.floor(((new Date().getTime()) - Date.parse(props.user.CreatedAt)) / (1000*60*60*24 * 365))
    }

    return (
        <header id={"bio"}>
                <img src={props.user.AvatarURL} id={"avatar"} />
                <div id={"bio-info"}>
                    <h2>@{props.user.Login} ({props.user.Name})</h2>
                    <h3>Created {getAccountAge()} years ago</h3>
                    <p>
                        {
                            reactStringReplace(props.user.Bio, /\B@([\w-]+)/gm, (match, i) => (
                                <Link key={i} to={`/${match}`} >@{match}</Link>
                            ))
                        }
                    </p>
                    <table>
                        {
                            props.user.Location && (
                                <tr>
                                    <td>Location</td>
                                    <td>{props.user.Location}</td>
                                </tr>
                            )
                        }
                        {
                            props.user.Company && (
                                <tr>
                                    <td>Company</td>
                                    <td>
                                        {
                                            reactStringReplace(props.user.Company, /\B@([\w-]+)/gm, (match, i) => (
                                                <Link key={i} to={`/${match}`} >@{match}</Link>
                                            ))
                                        }
                                    </td>
                                </tr>
                            )
                        }
                        {
                            props.user.WebsiteURL && (
                                <tr>
                                    <td>Website</td>
                                    <td>
                                        {
                                            (props.user.WebsiteURL).match(/^(http|https):\/\//)
                                                ?
                                                <a href={props.user.WebsiteURL} target={"_blank"} rel={"noopener"}>{props.user.WebsiteURL}</a>
                                                :
                                                <a href={"http://" + props.user.WebsiteURL} target={"_blank"} rel={"noopener"}>{props.user.WebsiteURL}</a>
                                        }
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
                            <div className="num">{props.user.Repositories.length >= 50 ? "50+" : props.user.Repositories.length}</div>
                                Repos
                            </div>
                            <div className="badge">
                                <div className="num">{props.user.Followers}</div>
                                Followers
                            </div>
                            <div className="badge">
                                <div className="num">{props.user.Following}</div>
                                Following
                            </div>
                        </>
                    }
                </div>
        </header>
    );
}

export default Bio;