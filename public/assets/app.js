var app = new Vue({
  el: "#app",
  data: {
    ws: null,
    serverUrl: "ws://localhost:8090/ws",
    messages: [],
    newMessage: "jkhjkh",
    roomInput: null,
    rooms: [],
    user: {
      name: "jatin",
    },
  },
  mounted: function () {
    this.connectToWebsocket();
  },
  methods: {
    connect() {
      this.connectToWebsocket();
    },
    connectToWebsocket() {
      this.ws = new WebSocket(this.serverUrl + "?name=" + this.user.name);
      this.ws.addEventListener("open", (event) => {
        this.onWebsocketOpen(event);
      });
      this.ws.addEventListener("message", (event) => {
        this.handleNewMessage(event);
      });
    },
    onWebsocketOpen() {
      console.log("connected to WS!");
    },
    handleNewMessage(event) {
      let data = event.data;
      data = data.split(/\r?\n/);

      for (let i = 0; i < data.length; i++) {
        let msg = JSON.parse(data[i]);
        // display the message in the correct room.
        const room = this.findRoom(msg.target);
        if (typeof room !== "undefined") {
          room.messages.push(msg);
        }
      }
    },
    sendMessage(room) {
      console.log("send message",room)
      // send message to correct room.
      if (room.newMessage !== "") {
        this.ws.send(
          JSON.stringify({
            action: "send-message",
            message: room.newMessage,
            target: room.name,
          })
        );
        room.newMessage = "";
      }
    },
    findRoom(roomName) {
      for (let i = 0; i < this.rooms.length; i++) {
        if (this.rooms[i].name === roomName) {
          return this.rooms[i];
        }
      }
    },
    joinRoom() {
      console.log('join room ', this.roomInput)
      this.ws.send(
        JSON.stringify({ action: "join-room", message: this.roomInput })
      );
      this.messages = [];
      this.rooms.push({ name: this.roomInput, messages: [] });
      this.roomInput = "";
    },
    leaveRoom(room) {
      this.ws.send(
        JSON.stringify({ action: "leave-room", message: room.name })
      );

      for (let i = 0; i < this.rooms.length; i++) {
        if (this.rooms[i].name === room.name) {
          this.rooms.splice(i, 1);
          break;
        }
      }
    },
  },
});
