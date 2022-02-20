const configs = document.getElementById("config-row");
const logs = document.getElementById("logs");
const youtubeSearchButton = document.getElementById("youtube-search");
const youtubeSwitchButton = document.getElementById("youtube-switch");
const youtubeInput = document.getElementById("youtube-input");
const youtubeThumbnail = document.getElementById("youtube-thumb");
const yotubeURL = document.getElementById("youtube-url");

let direct = false;

const errorHandler = async (data, endpoint, errored = false) => {
    endpoint = endpoint.replace("./api", "").split('?')[0];
    const status = data.status;
    if (!errored) data = await data.json().catch(() => data);
    logs.innerText += (endpoint + "  " + status + "  " + JSON.stringify(data) + "\n");
    logs.scrollTop = logs.scrollHeight - logs.clientHeight;
    notify(`HTTP ${status} - ${endpoint}`);
    if (!errored) return data;
};
const niceFetch = (value, method = "GET", body, headers) => {
    return fetch(value, {method, body, headers})
        .then(async r => await errorHandler(r, value))
        .catch(async e => await errorHandler(e, value, true));
};
const refreshMusic = async () => {
    const musicList = await niceFetch("./api/music");
    const musicPicker = document.getElementById("musicpickerlist");
    musicPicker.innerHTML = '';
    musicList.forEach(elem => {
        const option = document.createElement('option');
        option.value = elem;
        musicPicker.appendChild(option);
    });
};
refreshMusic();
const getSelectedMusic = () => {
    return document.getElementById("musicpicker").value;
};
const getConfig = async () => {
    const config = await niceFetch("./api/config");
    for (let key in config) {
        configs.innerHTML += `<div class="col-md-6">
            <label for="${key}" class="text-white">${key}:</label>
            <input type="${typeof config[key]}" class="form-control setting bg-dark text-white"
                   id="${key}"
                   value="${config[key]}">
            </div>`;
    }
};
getConfig();
const saveConfig = () => {
    const inputs = document.getElementsByClassName("setting");
    const settings = {};
    for (let data of inputs) {
        let value = data.value;
        if (["true", "false"].includes(value)) value = (value === 'true');
        if (data.attributes.type.value === "number") value = Number(value);
        settings[data.id] = value;
    }
    niceFetch("./api/save", "POST", JSON.stringify(settings), {'Content-Type': 'application/json'});
};

function htmlDecode(input) {
    const doc = new DOMParser().parseFromString(input, "text/html");
    return doc.documentElement.textContent;
}

function notify(message) {
    const snackbar = document.getElementById("snackbar");
    snackbar.className = "show";
    snackbar.textContent = message;
    setTimeout(function () {
        snackbar.className = snackbar.className.replace("show", "");
    }, 1500);
}


youtubeSearchButton.addEventListener("click", async () => {
    const data = await niceFetch(`./api/yt?q=${encodeURIComponent(document.getElementById("youtube-input").value)}&search=true`);
    if (data?.error) return;
    ["title", "channelTitle", "description"].forEach(prop => {
        document.getElementById("youtube-" + prop).textContent = htmlDecode(data?.snippet?.[prop]);
    });
    youtubeThumbnail.src = data.snippet.thumbnails.high.url;
    yotubeURL.href = "https://youtu.be/" + data?.id?.videoId;
});

youtubeSwitchButton.addEventListener("click", async () => {
    if (youtubeSwitchButton.checked) {
        youtubeSearchButton.setAttribute("disabled", "disabled");
        youtubeInput.setAttribute("placeholder", "https://www.youtube.com/watch?v=Vhh_GeBPOhs");
        direct = true;
    } else {
        direct = false;
        youtubeSearchButton.removeAttribute("disabled");
        youtubeInput.setAttribute("placeholder", "Brick Hustley - Don't give up");
    }
});