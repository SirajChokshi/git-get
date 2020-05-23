import React, {useState} from "react";
const [user, setUser] = useState()

export const getUser = (username) => {
    let url = "https://api.github.com/users/" + username;
    fetch(url)
        .then((data) => {
            return data.json();
        }).then((data) => {
            setUser(data['info']);
            console.log(user);
        });
};