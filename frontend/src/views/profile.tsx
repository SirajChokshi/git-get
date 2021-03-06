import React, {useEffect, useState} from 'react';
import { Link, Redirect } from '@reach/router';
import Error from './error'
import { User } from '../types'
import Bio from "../components/bio";
import QuickStats from "../components/quickStats"
import LoadingIndicator from "../components/LoadingIndicator"

interface ProfileProps {
    username: string;
}

const Profile = (props: ProfileProps) => {

    const BASE_URL = process.env.REACT_APP_API_URL;

    const [user, setUser] = useState<User | -1 | null>(null);

    const fetchUser = (username: string) : void => {
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
            setUser(-1);
        })
    }

    useEffect(() => {
        setUser(null)
        fetchUser(props.username)
    }, [props.username])
    useEffect(() => console.log(user), [user])

    return (
        <div className="App">
            <header>
                {user === null
                    ?
                    <LoadingIndicator size={"2.5em"} />
                    :
                    user === -1
                    ?
                    <Error />
                    :
                    (
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