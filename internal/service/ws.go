package service

import (
	"encoding/json"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WSService struct {
	Clients   []*WSClient
	Subscribe chan *WSClient
	Notify    chan *WSNotify
	mutex     sync.Mutex
}

type WSClient struct {
	Context *websocket.Conn
	UserID  int64
}

type WSNotify struct {
	ToUserID int64
	Data     interface{}
}

func NewWSService() *WSService {
	service := &WSService{
		Clients:   []*WSClient{},
		Subscribe: make(chan *WSClient),
		Notify:    make(chan *WSNotify),
	}
	go service.Run()
	return service
}

func (service *WSService) Run() {
	for {
		select {
		case client := <-service.Subscribe:
			service.addClient(client)
			go service.handleConnection(client)
		case notify := <-service.Notify:
			service.notify(notify)
		}
	}
}

func (service *WSService) addClient(client *WSClient) {
	service.mutex.Lock()
	defer service.mutex.Unlock()
	service.Clients = append(service.Clients, client)
}

func (service *WSService) handleConnection(client *WSClient) {
	defer func() {
		service.removeClient(client.Context)
	}()

	for {
		_, _, err := client.Context.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (service *WSService) notify(notify *WSNotify) {
	for _, client := range service.Clients {
		if client.UserID == notify.ToUserID {
			message, err := json.Marshal(notify.Data)
			if err == nil {
				client.Context.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
}

func (service *WSService) removeClient(conn *websocket.Conn) {
	service.mutex.Lock()
	defer service.mutex.Unlock()
	for i, client := range service.Clients {
		if client.Context == conn {
			service.Clients = append(service.Clients[:i], service.Clients[i+1:]...)
			return
		}
	}
}
