package main

import "html/template"

var rootTemplate = template.Must(template.New("root").Parse(`<!-- FROM: https://www.websocket.org/echo.html -->
<!DOCTYPE html>
<htmll>
<head>
<meta charset="utf-8" />
<script language="javascript" type="text/javascript">
//var wsUri = "ws://echo.websocket.org/"; 
var wsUri = "ws://{{.}}/socket"; 
var output;
var input;
var send;
function init() {
  output = document.getElementById("output");
  input = document.getElementById("input");
  send = document.getElementById("send");
  send.onclick = sendClickHandler;
  input.onkeydown = function(event) { if (event.keyCode == 13) send.click(); };
  testWebSocket();
}
function sendClickHandler() {
  doSend(input.value);
  input.value = '';
}
function testWebSocket() {
  websocket = new WebSocket(wsUri);
  websocket.onopen = function(evt) { onOpen(evt) };
  websocket.onclose = function(evt) { onClose(evt) };
  websocket.onmessage = function(evt) { onMessage(evt) };
  websocket.onerror = function(evt) { onError(evt) }; 
}
function onOpen(evt) {
  writeToScreen("CONNECTED");
  //doSend("WebSocket rocks");
}
function onClose(evt) {
  writeToScreen("DISCONNECTED");
}
function onMessage(evt) {
  writeToScreen('<span style="color: blue;">RESPONSE: ' + evt.data+'</span>');
  //websocket.close();
}
function onError(evt) {
  writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data); 
}
function doSend(message) {
  writeToScreen("SENT: " + message);
  websocket.send(message);
}
function writeToScreen(message) {
  var pre = document.createElement("p");
  pre.style.wordWrap = "break-word";
  pre.innerHTML = message;
  output.appendChild(pre);
}
window.addEventListener("load", init, false);
</script>
<h2>WebSocket Test</h2>
<input type="text" id="input" /><input type="button" id="send" value="Send" />
<div id="output"></div>
</html>
`))
