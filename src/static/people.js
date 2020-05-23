export const getUser = (async (username) => {
    let url = "https://api.github.com/users/" + username;
    const info = await fetch(url)
        .then((data) => {
            console.log(data)
            return data.json()
        })
    return info['login'];
});