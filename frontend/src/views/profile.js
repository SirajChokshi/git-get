import React, {useEffect, useState} from 'react';
import { Link, Redirect } from '@reach/router';
import Bio from "../components/bio";
import QuickStats from "../components/quickStats"
import LoadingIndicator from "../components/LoadingIndicator"

const Profile = (props) => {

    // const BASE_URL = "https://arcane-ocean-76968.herokuapp.com/"
    const BASE_URL = "http://localhost:8080/"

    const [user, setUser] = useState({loading: true})

    const fetchUser = (username) => {
        fetch(`${BASE_URL}get/${username}`,
            {method: "GET", headers: {'Content-Type': 'application/json'}}
        ).then(
            (userData => userData.json())
        ).then (
            json => {
                setUser(json)
                console.log(json)
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
                    <LoadingIndicator />
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