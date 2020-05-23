export const getUser = (username) => {
    let url = "https://api.github.com/users/" + username;
    const info = fetch(url)
        .then((data) => {
            return data.json();
        }).then((data) => {
            console.log(data);
        });
};