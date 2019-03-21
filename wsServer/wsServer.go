package wsServer

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// работающий эххо сервер
// address
// IP and PORT
var address string = "192.168.0.101:8085"

var addr = flag.String("addr", address, "http service address")

var upgrader = websocket.Upgrader{} // use default options

// эхо сервер.
func echo(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	} else {
		fmt.Println("New conn")
		err = c.WriteMessage(1, []byte("Hello User, your connection accept"))
		if err != nil {
			log.Println("write:", err)
		}
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

// Start - start server WS
func Start() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
