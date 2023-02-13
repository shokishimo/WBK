async function sendGetRequest(query) {
    try {
        const response = await fetch(`https://bestkeyboard.onrender.com/getRanking?number=${query}`);
        return await response.json();
    } catch (error) {
        console.log(error);
    }
}

sendGetRequest("3").then(data => {
    let url = data[0].url;
    let pic = document.getElementById("picture1");
    pic.style.backgroundImage = `url(${url})`;
    let rank = data[0].ranking;
    let rankLetter = document.getElementById("ranking-letter1");
    rankLetter.textContent = rank
    rankLetter.setAttribute("style", "font-size: 18px; padding-left: 11px; padding-top: 4px; color: #fff;")

    url = data[1].url;
    pic = document.getElementById("picture2");
    pic.style.backgroundImage = `url(${url})`;
    rank = data[1].ranking;
    rankLetter = document.getElementById("ranking-letter2");
    rankLetter.textContent = rank
    rankLetter.setAttribute("style", "font-size: 18px; padding-left: 11px; padding-top: 4px; color: #fff;")

    url = data[2].url;
    pic = document.getElementById("picture3");
    pic.style.backgroundImage = `url(${url})`;
    rank = data[2].ranking;
    rankLetter = document.getElementById("ranking-letter3");
    rankLetter.textContent = rank
    rankLetter.setAttribute("style", "font-size: 18px; padding-left: 11px; padding-top: 4px; color: #fff;")
});