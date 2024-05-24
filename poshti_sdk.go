package poshti_sdk

import (
	"encoding/json"
    "fmt"
    "log"
	"sync"
	"github.com/gorilla/websocket"
)

type Message struct {
	JoinRef string			`json:"join_ref"`
	MsgRef  string          `json:"msg_ref"`
	Topic   string          `json:"topic"`
	Event   string          `json:"event"`
	Payload json.RawMessage `json:"payload"`
}

type Subscription struct {
	Topic string
	Callback func(msg Message)
}

type Client struct {
	conn *websocket.Conn
	url string
	mu sync.Mutex 
	messageChan chan Message
	subscriptions map[string]Subscription
}

func NewClient(url string) *Client {
	return &Client{
		url:           url,
		messageChan:   make(chan Message),
		subscriptions: make(map[string]Subscription),
	}
}

