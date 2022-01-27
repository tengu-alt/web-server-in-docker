
function doRequest(){
  const emailInput = document.querySelector('.emailInput');
const passwordInput = document.querySelector('.passwordInput');
const FnameInput = document.querySelector('.frNameInput');
const LnameInput = document.querySelector('.lstNameInput');

  const url = window.location.origin + "/signUp";
  
  const data = {
    FirstName: FnameInput.value,
    LastName: LnameInput.value,
    email: emailInput.value,
    password: passwordInput.value,
  }
  const response = fetch(url, {
    method: 'POST',
    mode: 'cors', 
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow', 
    referrerPolicy: 'no-referrer', 
    body: JSON.stringify(data) 
  }).then((response) => {
    return response.json();
  })
  .then((data) => {
    function errorMessage (object){
      let fnameErr = document.querySelector(".fNameErrorDiv")
      let lnameErr = document.querySelector(".lNamelErrorDiv")
      let emailErr = document.querySelector(".emailErrorDiv")
      let passwordErr = document.querySelector(".passwordErrorDiv")
      for (let i = 0; i < object.length; i++){
          if(!object[i].FieldValue){
            fnameErr = "";
            lnameErr = "";
            emailErr = "";
            passwordErr = "";
            console.log("success")
            window.location.href = 'submit.html'
            return
          }

        if(object[i].FieldValue === "FirstName"){
          fnameErr.innerHTML = object[i].ErrMassage
        }

        if(object[i].FieldValue === "LastName"){
          lnameErr.innerHTML = object[i].ErrMassage
        }

        if(object[i].FieldValue === "Email"){
          emailErr.innerHTML = object[i].ErrMassage
        }
        if(object[i].FieldValue === "Password"){
          passwordErr.innerHTML = object[i].ErrMassage
        }
      }

    }
    errorMessage(data)
  });
}
function doLogin(){
  const emailInput = document.querySelector('.emailInputLogin');
  const passwordInput = document.querySelector('.passwordInputLogin');

  const url = window.location.origin + "/login";

  const data = {
    email: emailInput.value,
    password: passwordInput.value,
  }
  const response = fetch(url, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify(data)
  }).then((response) => {
    return response.json();
  })
      .then((data) => {
        function login (object) {
          let loginErr = document.querySelector(".loginError");
          if (object.ErrMassage === "invalid data"){
            loginErr.innerHTML = object.ErrMassage;
          }
          else {
            loginErr.innerHTML = "";
            console.log("success");
            window.location.href = 'login.html';
            return;
          }
        }
        login(data)
      });




}

function registration() {
const sendBut = document.querySelector('.button')
sendBut.addEventListener('click', doRequest);
const loginBut = document.querySelector('.buttonLogin')
  loginBut.addEventListener('click', doLogin)
}
window.addEventListener('load', registration);