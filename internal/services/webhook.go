package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type WebhookService struct {
	mu          sync.Mutex
	subscribers map[string][]string 
}

type WebhookPayload struct {
	AreaCode string `json:"area_code"`
	Mode     string `json:"mode"`
	Count    int    `json:"count"`
}


func NewWebhookService() *WebhookService {
	return &WebhookService{
		subscribers: make(map[string][]string),
	}
}


func (ws *WebhookService) Subscribe(areaCode, url string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.subscribers[areaCode] = append(ws.subscribers[areaCode], url)
}

func (ws *WebhookService) Unsubscribe(areaCode, url string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	subscribers := ws.subscribers[areaCode]
	for i, subscriber := range subscribers {
		if subscriber == url {
			ws.subscribers[areaCode] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
}

func (ws *WebhookService) Notify(areaCode string, mode string, count int) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	payload := WebhookPayload{
		AreaCode: areaCode,
		Mode:     mode,
		Count:    count,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		return
	}

	for _, url := range ws.subscribers[areaCode] {
		go func(subscriberURL string) {
			log.Printf("Sending notification to %s with data: %+v", subscriberURL, payload)
			resp, err := http.Post(subscriberURL, "application/json", bytes.NewBuffer(data))
			if err != nil {
				log.Printf("Failed to send notification to %s: %v", subscriberURL, err)
				return
			}
			defer resp.Body.Close()

	
			if resp.StatusCode != http.StatusOK {
				log.Printf("Notification to %s failed with status: %s", subscriberURL, resp.Status)
			} else {
				log.Printf("Notification sent successfully to %s", subscriberURL)
			}
		}(url)
	}
}
