async function sendGetRequest(query) {
    try {
        const response = await fetch(`/getRanking?number=${query}`);
        const data = await response.json();
        return data
    } catch (error) {
        console.log(error);
    }
}

// Example usage
sendGetRequest("3").then(data => console.log(data));