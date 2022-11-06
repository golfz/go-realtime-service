const wsURL = "ws://localhost/realtime/subscribe"
let websocket;
let selectedDevices = []
let selectedEmployees = []
let transactions = []

const MESSAGE_TEMPLATE = {
    "scope": "test",
    "topic": "transaction",
    "args": {
        "hardware_uuid": [], // array of strings
        "employee_uuid": [], // array of strings
    }
}

setIdleState();

// clear websocket, close websocket, and set to null for garbage collection
window.onbeforeunload = function () {
    websocket.onclose = function () {
    };
    websocket.close()
    websocket = null
}

// click "SUBSCRIBE" btn
document.getElementById("subscribe_btn").onclick = function (event) {
    console.log("click SUBSCRIBE button.")

    selectedDevices = []
    selectedEmployees = []

    let cbDevices = document.getElementsByClassName("device_checkbox")
    for (let i = 0; cbDevices[i]; ++i) {
        if (cbDevices[i].checked) {
            selectedDevices = [...selectedDevices, cbDevices[i].value]
        }
    }

    let cbEmployees = document.getElementsByClassName("employee_checkbox")
    for (let i = 0; cbEmployees[i]; ++i) {
        if (cbEmployees[i].checked) {
            selectedEmployees = [...selectedEmployees, cbEmployees[i].value]
        }
    }

    console.log(selectedDevices)
    if (selectedDevices.length === 0 && selectedEmployees.length === 0) {
        alert("Please select some device(s) or employee(s)")
        return false
    }

    setMonitoringState()

    openWebSocket()

    return false
}

// click "UNSUBSCRIBE" btn
document.getElementById("unsubscribe_btn").onclick = function (event) {
    closeWebSocket()
    setIdleState()
    return false
}

// open websocket
function openWebSocket() {
    console.log("try to open WebSocket")

    if (websocket) {
        console.log("websocket is already opened.")
        return false;
    }
    websocket = new WebSocket(wsURL);

    clearAllDataBeforeSubscribe()

    websocket.onopen = onWebSocketOpened
    websocket.onclose = onWebSocketClosed
    websocket.onmessage = onWebSocketMessage
    websocket.onerror = onWebSocketError

    return false;
}

// close websocket
function closeWebSocket() {
    console.log("try to close WebSocket")

    if (!websocket) {
        console.log("WebSocket is null, no closing")
        return false;
    }

    clearWebSocketListener()

    websocket.close();
    websocket = null

    console.log("websocket closed.")

    return false
}

// send message to websocket, after websocket is opened
function sendWebSocketMessage(msg) {
    if (!websocket) {
        return false;
    }
    console.log("SEND: " + msg);
    websocket.send(msg);
    return true;
}

// Event: Websocket is opened
function onWebSocketOpened(event) {
    console.log("websocket is opened.");
    if (selectedDevices.length === 0 && selectedEmployees.length === 0) {
        console.log("no selected device, no send subscribe message")
        return false

    } else {
        let subscribeObj = MESSAGE_TEMPLATE
        subscribeObj["args"]["hardware_uuid"] = selectedDevices
        subscribeObj["args"]["employee_uuid"] = selectedEmployees
        sendWebSocketMessage(JSON.stringify(subscribeObj))
    }
}

// Event: Websocket is closed
// if server is closed, websocket will be closed too.
function onWebSocketClosed(event) {
    console.log("websocket is closed.")

    setIdleState()

    clearWebSocketListener()
    closeWebSocket()

    alert("WebSocket is closed.")

    return false
}

// Event: Websocket received message
function onWebSocketMessage(event) {
    console.log("websocket on message: " + event.data)
    let msg = JSON.parse(event.data)

    let sEmployee = msg["payload"]
    let emp = JSON.parse(sEmployee)

    // console.log(emp)

    transactions = [emp, ...transactions]
    if (transactions.length > 5) {
        transactions = transactions.slice(0, 5)
    }
    renderTable()
}

// Event: Websocket error
function onWebSocketError(event) {
    console.log("websocket on error: " + event.type)
    console.log(event);
}

// clear all websocket listener
function clearWebSocketListener() {
    websocket.onopen = function () {
    };
    websocket.onclose = function () {
    };
    websocket.onmessage = function () {
    };
    websocket.onerror = function () {
    };
}

// clear all data before subscribe, including table and global variables
function clearAllDataBeforeSubscribe() {
    transactions = []
    renderTable()
}

// render table from transactions (global variable)
function renderTable() {
    let rows = ""
    transactions.forEach(function (item, index, arr) {
        const date = new Date(item["time_stamp"])
        const day = moment(date).format("DD MMMM YYYY")
        const time = moment(date).format("HH:mm:ss (Z)")
        rows += `<tr>
                        <td>${day}</td>
                        <td>${time}</td>
                        <td>${item["hardware"]["title_th"]}</td>
                        <td><img src="${item["picture_url"]}"/></td>
                        <td>${item["employee"]["employee_code"]}</td>
                        <td>${item["employee"]["name_th"]}</td>
                    </tr>`
    })
    let tableBody = document.getElementById("table_body")
    tableBody.innerHTML = rows
}

// set UI state to idle, before click "SUBSCRIBE" button
function setIdleState() {
    // enable checkbox
    document.getElementById("device1").disabled = false;
    document.getElementById("device2").disabled = false;
    //
    document.getElementById("employee1").disabled = false;
    document.getElementById("employee2").disabled = false;
    document.getElementById("employee3").disabled = false;
    document.getElementById("employee4").disabled = false;
    document.getElementById("employee5").disabled = false;
    // show SUBSCRIBE btn
    document.getElementById("subscribe_btn").style.display = "inline-block"
    document.getElementById("unsubscribe_btn").style.display = "none"
}

// set UI state to monitoring, after click SUBSCRIBE button
function setMonitoringState() {
    // disable checkbox
    document.getElementById("device1").disabled = true;
    document.getElementById("device2").disabled = true;
    //
    document.getElementById("employee1").disabled = true;
    document.getElementById("employee2").disabled = true;
    document.getElementById("employee3").disabled = true;
    document.getElementById("employee4").disabled = true;
    document.getElementById("employee5").disabled = true;
    // show UNSUBSCRIBE btn
    document.getElementById("subscribe_btn").style.display = "none"
    document.getElementById("unsubscribe_btn").style.display = "inline-block"
}