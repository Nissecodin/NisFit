function logIn() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `username=${username}&password=${password}`,  
    })
    .then(response => response.text())
    .then(data => alert(data))  
    .catch(error => console.error('Error:', error)); 
}

