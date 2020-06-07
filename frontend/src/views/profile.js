import React, {useEffect, useState} from 'react';
import { Link, Redirect } from '@reach/router';
import {FaSpinner} from 'react-icons/fa'
import Bio from "../components/bio";
import QuickStats from "../components/quickStats"

const client_id = `Iv1.82a68fd3b00d5cab`
const client_secret = `b6b9b2b580c0dfc25499d34d4e4fc9d4f33f0ce2`

const Profile = (props) => {

    const [user, setUser] = useState({loading: true})

    const fetchUser = (username) => {
        fetch(`http://localhost:8080/get/${username}`,
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

    useEffect(() => {
        setUser({loading: true})
        fetchUser(props.username)
    }, [props.username])
    useEffect(() => console.log(user), [user])

    return (
        <div className="App">
            <header className="App-header">
                {user.loading
                    ?
                    <FaSpinner size={"1.3em"} className={"spin-icon"} />
                    :
                    (
                        user.message
                        ?
                            <h1>Error -- Not Found</h1>
                        :
                        <>
                            <Bio
                                user={user}
                            />
                            <QuickStats
                                user={user}
                            />
                        </>
                    )
                }
            </header>
        </div>
    );
}

export default Profile;