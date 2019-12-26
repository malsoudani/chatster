$(function() {
    var socket = null;
    var msgBox = $("#chatbox textarea");
    var messages = $("#messages");

    $("#chatbox").submit(function() {
        if (!msgBox.val()) return false;
        if (!socket) {
            alert("there is no socket connection.");
            return false;
        }
        socket.send(msgBox.val());
        msgBox.val("");
        return false;
    });

    if (!window["WebSocket"]) {
        alert("Error: your browser does not supprot websockets");
    } else {
        socket = new WebSocket("ws://localhost:8080/room");
        socket.onclose = function() {
            alert("connection has been closed");
        };
        socket.onmessage = function(e) {
            messages.append($("<i>").text(e.data));
        };
    }
});
