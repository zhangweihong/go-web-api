package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func server(ws *websocket.Conn) {
	fmt.Printf("new connection\n")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
		if len(buf) > 0 {
			str := string(buf[:])
			fmt.Println(str)
		}
	}
	fmt.Printf(" => closing connection\n")
	ws.Close()
}

func main() {
	http.Handle("/websocket", websocket.Handler(server))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
