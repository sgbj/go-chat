# Chat

A simple chat application built with Go and WebSocket.

* Uses github.com/gorilla/websocket for the WebSocket implementation
* Uses goroutines and channels to synchronize around clients

![image](https://github.com/sgbj/go-chat/assets/5178445/eef17aa8-2124-4a6b-9890-658a5eb80b05)

## Getting started

Clone the repo

```bash
git clone https://github.com/sgbj/go-chat.git
cd go-todos
```

Install packages

```bash
go get .
```

Run the application

```bash
go run .
```

Open a browser to http://localhost:8080

## Future

* Use JSON for messages
* Support different types of messages (e.g., connected, disconnected, message)
* Add a name
