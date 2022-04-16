package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/", homeHandler)
	r.GET("/ws", websocketHandler)

	r.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func websocketHandler(c *gin.Context) {
	wsHandler(c.Writer, c.Request)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	fmt.Println(r.URL.RequestURI())
	for {
		// t, msg, err := conn.ReadMessage()
		// if err != nil {
		// 	break
		// }
		// print(t)
		msg := "hello"
		t := 1
		conn.WriteMessage(t, []byte(fmt.Sprintf("SUP %s", msg)))
		time.Sleep(time.Second)
	}
}
