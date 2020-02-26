//var socket = new WebSocket("ws://localhost:8080/ws");
var socket = new WebSocket("ws://sezzle-interview.azurewebsites.net/ws");

let connect = (cb) => {
    console.log("Attempting connection....");

    socket.onopen = () => {
        console.log("successfully connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("socket closed connection:  ", event);
    };

    socket.onerror = error => {
        console.log("socket error:  ", error);
    };
};

export { connect };