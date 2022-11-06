package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/golfz/go-realtime-service/k"
	"github.com/golfz/gorealtime"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	RealtimeTypeSubscription = iota
)

type RealtimeRequest struct {
	Scope string              `json:"scope"`
	Topic string              `json:"topic"`
	Args  map[string][]string `json:"args"`
}

var upgrader = websocket.Upgrader{} // use default options

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	http.HandleFunc("/realtime/subscribe", realtimeSubscribe)

	log.Println("realtime service is starting")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func realtimeSubscribe(w http.ResponseWriter, r *http.Request) {
	subscriberID := generateSubscribeID(6)
	defer func() {
		log.Printf("[subscriber: %s] closed & destroyed realtimeSubscribe(w,r) request handler func \n", subscriberID)
	}()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	log.Printf("[subscriber: %s] opened ws connection \n", subscriberID)

	_, message, err := c.ReadMessage()
	if err != nil {
		log.Printf("[subscriber: %s] ws read message err : %v \n", subscriberID, err)
		return
	}
	log.Printf("[subscriber: %s] recieved message: %s \n", subscriberID, message)

	var realtimeRequest RealtimeRequest
	err = json.Unmarshal(message, &realtimeRequest)
	if err != nil {
		log.Printf("[subscriber: %s] cannot unmarshal message to realtime-request \n", subscriberID)
		return
	}

	log.Printf("[subscriber: %s] marshal request to struct %#v \n", subscriberID, realtimeRequest)

	amqpEndpoint, err := k.GetAMQPEndpoint()
	if err != nil {
		log.Panicf("[subscriber: %s] cannot get amqp endpoint : %v \n", subscriberID, err)
	}

	ch := make(chan gorealtime.RealtimePayload, 2)
	defer close(ch)

	ctxWithCancel, cancel := context.WithCancel(context.Background())

	go gorealtime.Subscribe(ctxWithCancel, amqpEndpoint, realtimeRequest.Scope, realtimeRequest.Topic, subscriberID, realtimeRequest.Args, ch)

	stop := make(chan struct{})
	go func() {
		for {
			_, _, err = c.ReadMessage()
			if err != nil {
				log.Printf("[subscriber: %s] ws got err : %v \n", subscriberID, err)
				stop <- struct{}{}
				return
			}
		}
	}()

loop:
	for {
		select {
		case returnData := <-ch:
			log.Printf("[subscriber: %s] got data from subscriber : \n%v \n", subscriberID, returnData)
			_ = c.WriteJSON(returnData)

		case <-ticker.C:
			err = c.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Printf("[subscriber: %s] ws sent ping err : %v \n", subscriberID, err)
			}

		case <-stop:
			cancel()
			log.Printf("[subscriber: %s] call cancel() to stop subscribers \n", subscriberID)
			break loop
		}
	}

	log.Printf("[subscriber: %s] all subscribers was completely closed. \n", subscriberID)
}

func generateSubscribeID(n int) string {
	const charset = "abcdefhijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
