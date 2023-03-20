function submitLoginForm() {
    document.getElementById("loginedUsername").submit();
}

function submitHomeForm() {
    document.getElementById("home").submit();
}

function searchKeyboard() {
    document.getElementById("main-header-search-form").submit();
}

window.onload = function() {
    // login/username and rendering
    const loginDiv = document.getElementById('login-switch');
    const dropDown = document.getElementById('dropdown-content');
    const username = getUsernameCookie();
    if (username === "") { // show Login
        loginDiv.innerHTML = "<form class=\"main-header-controller-login\" id=\"loginedUsername\" action=\"/login\" method=\"get\" onclick=\"submitLoginForm()\">\n" +
                                "<div class=\"main-header-controller-login-logo\"></div>\n" +
                             "</form>";
    } else { // show username instead of login
        loginDiv.innerHTML = "<p class=\"logined-username\">" + username + "</p>"
        dropDown.innerHTML = "<a class=\"logout\" href=\"/logout\">Log out</a>"
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

function newKeyboardRequest() {
    const form = document.createElement('form');
    form.action = '/newKeyboardRequest';
    form.method = 'GET';
    document.body.appendChild(form);
    form.submit();
}