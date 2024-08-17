window.addEventListener("load", (event) => {
    fetch('http://localhost:8080/isTokenValid', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        credentials: 'include'
    })
    .then((response) => {
        if (!response.ok) {
            throw new Error("Not logged in");
        }
    })
    .then(() => {
        window.location.replace("/dashboard");
    })
    .catch((error) => {
        console.error("Error:", error);
    });
    
});