function submitLoginForm() {
    document.getElementById("loginedUsername").submit();
}

function searchKeyboard() {
    document.getElementById("main-header-search-form").submit();
}

function keyboardDetail(ranking) {
    document.getElementById("rank" + ranking).submit();
}

window.onload = function() {
    // login/username and rendering
    const loginDiv = document.getElementById('login-switch');
    const favLink  = document.getElementById('favoriteLink');
    const account  = document.getElementById('account');

    const username = getUsernameCookie();
    if (username === "") { // show Login
        loginDiv.innerHTML = "<form class=\"main-header-controller-login\" id=\"loginedUsername\" action=\"/login\" method=\"get\" onclick=\"submitLoginForm()\">\n" +
                                "<div class=\"main-header-controller-login-logo\"></div>\n" +
                             "</form>";
    } else { // show username instead of login
        loginDiv.innerHTML = "<p class=\"logined-username\">" + username + "</p>";
        favLink.innerHTML  = "<div class=\"main-header-favorite\" href=\"#\" aria-label=\"lookAtFavorites\"></div>";
        account.innerHTML  = "<div class=\"main-header-controller-user-account\" >\n" +
                                "<a class=\"account\" href=\"/account\">Account</a>\n" +
                                "<div id=\"dropdown-content\">\n" +
                                    "<a class=\"logout\" href=\"/logout\">Log out</a>\n" +
                                "</div>\n" +
                             "</div>";
    }
}

function getUsernameCookie(){
    let cookieArr = document.cookie.split(";");
    for(let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if(cookiePair[0].trim() === 'username') {
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