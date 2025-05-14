document.getElementById("RegistrationButton").addEventListener('click', function() {
  login = document.getElementById("loginreg").value;
  pass = document.getElementById("passreg").value;
  if (login == "" || pass == ""){
     alert("Write something...")
     return
  } 
  fetchregistration(login, pass)
});

document.getElementById("LoginButton").addEventListener('click', function() {
  login = document.getElementById("loginlog").value;
  pass = document.getElementById("passlog").value;
  if (login == "" || pass == ""){
     alert("Write something...")
     return
  } 
  fetchlogin(login, pass)
});

document.getElementById("ToLogin").addEventListener('click', function() {
     document.getElementById("Login").style.display = "block";
     document.getElementById("Registration").style.display = "none";
     document.getElementById("loginreg").value = '';
     document.getElementById("passreg").value = '';
});

document.getElementById("ToRegistration").addEventListener('click', function() {
     document.getElementById("Login").style.display = "none";
     document.getElementById("Registration").style.display = "block";
     document.getElementById("loginlog").value = '';
     document.getElementById("passlog").value = '';
});

function fetchregistration(login, pass){
   fetch("http://localhost:8008/registration", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
        "login": login,
        "password": pass,
    })
   })
   .then(response =>{
        if (!response.ok) {
            return response.json().then(error => { throw new Error(`Error with status ${response.status}: ${error.error}`) });
        }
        return response.json()
   })
   .then(data => {
        alert(data.message)
   })
   .catch(error => {
        console.error(`Error ${error.message}`)
        alert(error.message)
   })
};

function fetchlogin(login, pass){
   fetch("http://localhost:8008/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
        "login": login,
        "password": pass,
    })
   })
   .then(response =>{
        if (!response.ok) {
            return response.json().then(error => { throw new Error(`Error with status ${response.status}: ${error.error}`) });
        }
        return response.json()
   })
   .then(data => {
        alert(data.message)
   })
   .catch(error => {
        console.error(`Error ${error.message}`)
        alert(error.message)
   })
};
