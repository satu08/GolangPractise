package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"time"
)

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

//func websocketHandler(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer conn.Close()
//
//	for {
//		_, message, err := conn.ReadMessage()
//		if err != nil {
//			fmt.Println(err)
//		}
//		log.Printf("recv: %s", message)
//
//		err = conn.WriteMessage(websocket.TextMessage, message)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//}

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{conns: make(map[*websocket.Conn]bool)}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client ", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) handleBook(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client ", ws.RemoteAddr())
	for {
		payload := fmt.Sprintf("orderbook data %s\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err != io.EOF { // connection from another side has been closed
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := string(buf[:n]) // read-only bytes till n
		//fmt.Println(string(msg))
		s.broadcast([]byte(msg))
		//_, err = ws.Write([]byte("thank you for the msg"))
		//if err != nil {
		//	fmt.Println("write error:", err)
		//}
	}
}

func (s *Server) broadcast(msg []byte) {
	for conn := range s.conns {
		go func(conn *websocket.Conn) {
			if _, err := conn.Write(msg); err != nil {
				fmt.Println("write error:", err)
			}
		}(conn)
	}
}
func main() {

	//http.HandleFunc("/ws", websocketHandler)
	//err := http.ListenAndServe(":8082", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/book", websocket.Handler(server.handleBook))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("listen and server error---", err)
	}
}
