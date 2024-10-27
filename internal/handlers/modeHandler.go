package handlers

import (
	"Game_Mode_Usage_Web_service/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type ModeHandler struct {
	service *services.ModeService
}

func NewModeHandler(service *services.ModeService) *ModeHandler {
	return &ModeHandler{service: service}
}

func (h *ModeHandler) GetModeCounts(w http.ResponseWriter, r *http.Request) {
	areaCode := r.URL.Query().Get("area_code")

	if areaCode == "" {
		http.Error(w, "Area code is required", http.StatusBadRequest)
		return
	}

	counts, err := h.service.GetModeCounts(r.Context(), areaCode)
	if err != nil {
		log.Printf("Error retrieving counts: %v", err) 
		http.Error(w, "Failed to retrieve counts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(counts); err != nil {
		log.Printf("Error encoding response: %v", err) 
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ModeHandler) JoinMode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AreaCode string `json:"area_code"`
		Mode     string `json:"mode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.AreaCode == "" || req.Mode == "" {
		http.Error(w, "AreaCode and Mode must not be empty", http.StatusBadRequest)
		return
	}

	if err := h.service.JoinMode(r.Context(), req.AreaCode, req.Mode); err != nil {
		log.Printf("Failed to join mode: %v", err) 
		http.Error(w, "Failed to join mode: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User joined mode: %s in area code: %s", req.Mode, req.AreaCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Joined mode successfully"})
}

func (h *ModeHandler) LeaveMode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AreaCode string `json:"area_code"`
		Mode     string `json:"mode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.AreaCode == "" || req.Mode == "" {
		http.Error(w, "AreaCode and Mode must not be empty", http.StatusBadRequest)
		return
	}

	if err := h.service.LeaveMode(r.Context(), req.AreaCode, req.Mode); err != nil {
		log.Printf("Failed to leave mode: %v", err) 
		http.Error(w, "Failed to leave mode: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Left mode successfully"})
}

func (h *ModeHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AreaCode string `json:"area_code"`
		URL      string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.service.Subscribe(req.AreaCode, req.URL) 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Subscribed successfully"})
}

func (h *ModeHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AreaCode string `json:"area_code"`
		URL      string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.service.Unsubscribe(req.AreaCode, req.URL) 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Unsubscribed successfully"})
}
