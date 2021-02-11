# Sprint 1 - 2/10-2/17

## Ticket #2 - C2 Research

### Socket Programming
- https://dev.to/alicewilliamstech/getting-started-with-sockets-in-golang-2j66
- https://golangr.com/socket-server/
- https://golang.org/pkg/net/
  
### Database Storage
- https://entgo.io/
- https://entgo.io/docs/getting-started/
- https://entgo.io/docs/migrate/

## Ticket #3
  
### Executing commands/scripts
- https://golang.org/pkg/os/exec/
- https://medium.com/rungo/executing-shell-commands-script-files-and-executables-in-go-894814f1c0f7
- https://zetcode.com/golang/exec-command/
- https://github.com/golang/go/issues/22278
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

### Socket Programming

*See Ticket #2*

### Interacting with the OS
- https://golang.org/pkg/os/

## Ticket #4

The project is now named *Titan*... way cooler than RITSEC Duckies.

## Ticket #5

I wrote a basic server and client in Go that uses tcp sockets to communicate. The server just acts as an echo server for the client.

**Server output**
```
Starting tcp server on localhost:1337
Client 127.0.0.1:51204connected.
2021/02/11 16:40:49 echo
```

**Client Output**
```
Message: echo
2021/02/11 16:40:49 Server echo: echo
Message: 
```