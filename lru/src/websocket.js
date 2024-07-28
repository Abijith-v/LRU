import { w3cwebsocket as webSocket } from 'websocket';

const uiConnection = new webSocket('ws://localhost:8080');

uiConnection.onopen = () => {
    console.log('UI Connection established');
};

uiConnection.onclose = () => {
    console.log('UI Disconnected');
};

export default uiConnection;
