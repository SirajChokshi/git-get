import React from 'react';
import { Link } from '@reach/router';

const Profile = (props) => {

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