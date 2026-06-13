class SocketService {
    constructor() {
        this.socket = null;
    }

    connect(
        roomId,
        username,
    ) {
        this.socket =
            new WebSocket(
                `ws://localhost:8080/ws/${roomId}?username=${encodeURIComponent(username)}`
            );
            
        this.socket.onopen = () => {
            console.log("Connected");
        };

        this.socket.onclose = () => {
            console.log("Disconnected");
        };

        this.socket.onerror = (err) => {
            console.error(err);
        };

        return this.socket;
    }

    send(data) {
        if (
            this.socket &&
            this.socket.readyState === WebSocket.OPEN
        ) {
            this.socket.send(
                JSON.stringify(data)
            );
        }
    }
}

export default new SocketService();