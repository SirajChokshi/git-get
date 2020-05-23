import React, {useEffect, useState} from 'react';
import { Link, Redirect } from '@reach/router';

const client_id = `Iv1.82a68fd3b00d5cab`
const client_secret = `b6b9b2b580c0dfc25499d34d4e4fc9d4f33f0ce2`

const Profile = (props) => {

    const [user, setUser] = useState()

    const fetchUser = (username) => {
        fetch(`https://api.github.com/users/${username}?client_id=${client_id}?client_secret=${client_secret}`,
            {method: "GET", headers: {'Content-Type': 'application/json'}}
        ).then(
            (userData => userData.json())
        ).then (
            json => {
                setUser(json)
            }
        ).catch((e) => {
            console.error(e);
        })
    }

    useEffect(() => fetchUser(props.username), [props.username])
    useEffect(() => console.log(user), [user])

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