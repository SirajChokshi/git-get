export const getUser = (username) => fetch("https://api.github.com/users/" + username, {headers: {'Content-Type': 'application/json'}})
    .then (res => res.json())