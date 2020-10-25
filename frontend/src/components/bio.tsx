import React from "react";
import reactStringReplace from 'react-string-replace'
import {User} from '../types'
import {Link} from '@reach/router'
import './bio.scss'

interface BioProps {
    user: User;
}

const Bio = (props: BioProps) => {

    const {user} = props;

    const getAccountAge = (date: string) : number => {
        return Math.floor(((new Date().getTime()) - Date.parse(date)) / (1000*60*60*24 * 365))
    }

    return (
        <header id="bio">
                <div id="avatar">
                    <img src={user.AvatarURL} />
                    { user.Name && <h2>{user.Name}</h2> }
                    <h2>@{user.Login}</h2>
                </div>
                <div id="bio-info">
                    <p>
                        {
                            reactStringReplace(user.Bio, /\B@([\w-]+)/gm, (match, i) => (
                                <Link key={i} to={`/${match}`} >@{match}</Link>
                            ))
                        }
                    </p>
                    <h3>Created {getAccountAge(user.CreatedAt)} years ago</h3>
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
                            user.Company && (
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
                            user.WebsiteURL && (
                                <tr>
                                    <td>Website</td>
                                    <td>
                                        {
                                            (user.WebsiteURL).match(/^(http|https):\/\//)
                                                ?
                                                <a href={user.WebsiteURL} target={"_blank"} rel={"noopener"}>{user.WebsiteURL}</a>
                                                :
                                                <a href={"http://" + user.WebsiteURL} target={"_blank"} rel={"noopener"}>{user.WebsiteURL}</a>
                                        }
                                    </td>
                                </tr>
                            )
                        }

                    </table>
                </div>
                <div id="extra-bio-info">
                    {/* { user.type !== "Organization" && */}
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
                    {/* } */}
                </div>
        </header>
    );
}

export default Bio;