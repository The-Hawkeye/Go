package main

import (
	"Game_Mode_Usage_Web_service/configs"
	"Game_Mode_Usage_Web_service/internal/handlers"
	"Game_Mode_Usage_Web_service/internal/services"
	"Game_Mode_Usage_Web_service/internal/storage"
	"log"
	"net/http"
	"time"
)

func main() {
	config := configs.LoadConfig()

	repo := storage.NewRedisRepository(config.RedisAddress, config.RedisPassword)
	webhookService := services.NewWebhookService()           // Create a new webhook service
	service := services.NewModeService(repo, webhookService) // Pass it to ModeService
	handler := handlers.NewModeHandler(service)

	http.HandleFunc("/modes", handler.GetModeCounts)
	http.HandleFunc("/mode/join", handler.JoinMode)
	http.HandleFunc("/mode/leave", handler.LeaveMode)
	http.HandleFunc("/subscribe", handler.Subscribe)     // Add subscribe endpoint
	http.HandleFunc("/unsubscribe", handler.Unsubscribe) // Add unsubscribe endpoint

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func logRedisContents(repo *storage.RedisRepository) {
	for {
		// Assuming you have a method in your repository to get all mode counts
		counts, err := repo.GetAllModeCounts() // Implement this method in your Redis repository
		if err != nil {
			log.Printf("Error fetching mode counts from Redis: %v", err)
		} else {
			log.Printf("Current mode counts in Redis: %+v", counts)
		}

		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}
}
