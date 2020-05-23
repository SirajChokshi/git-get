import React, {useEffect, useState} from 'react';
import { Link, Redirect } from '@reach/router';

const Profile = (props) => {

    const [user, setUser] = useState()

    const fetchUser = (username) => {
        fetch(`https://api.github.com/user/${username}`,
            {method: "GET", headers: {'Content-Type': 'application/json'}}
        ).then(
            (userData => userData.json())
        ).then (
            json => {
                setUser(json)
                console.log(user)
            }
        ).catch((e) => {
            console.error(e);
        })
    }

    // const fetchRateLimit = () => {
    //     fetch(`https://api.github.com/rate_limit`, {method: "GET", headers: {'Content-Type': 'application/json'}}
    //     ).then(
    //         (lim => lim.json())
    //     ).then(
    //         json => console.log(json)
    //     ).catch((e) => {
    //         console.error(e);
    //     })
    // }

    useEffect(() => fetchUser(props.username))

    return (
        <div className="App">
            <header className="App-header">
                <h1>
                    PROFILE OF {props.username}
                </h1>
            </header>
        </div>
    );
}

export default Profile;