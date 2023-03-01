function submitLoginForm() {
    document.getElementById("loginedUsername").submit();
}

function submitHomeForm() {
    document.getElementById("home").submit();
}

const loginDiv = document.getElementById('login-switch');

window.onload = function() {
    const username = getUsernameCookie();
    if (username === "") { // show Login
        loginDiv.innerHTML = "<form class=\"main-header-controller-login\" id=\"loginedUsername\" action=\"/login\" method=\"get\" onclick=\"submitLoginForm()\">\n" +
                                "<div class=\"main-header-controller-login-logo\"></div>\n" +
                             "</form>";
    } else { // show username instead of login
        loginDiv.innerHTML = "<p class=\"logined-username\">" + username + "</p>"
    }
}

function getUsernameCookie(){
    let cookieArr = document.cookie.split(";");
    for(let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if(cookiePair[0].trim() === 'username') {
            console.log(cookiePair[1]);
            return cookiePair[1];
        }
    }
    return ""
}