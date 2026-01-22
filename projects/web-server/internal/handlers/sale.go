package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web-server/internal/database"
	"web-server/internal/models"
)

type SaleHandler struct {
	db database.Service
}

func NewSaleHandler(db database.Service) *SaleHandler {
	return &SaleHandler{db: db}
}

func (h *SaleHandler) CreateSale(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Basic validation
	if len(req.Items) == 0 {
		http.Error(w, "sale must contain at least one item", http.StatusBadRequest)
		return
	}

	sale, err := h.db.CreateSale(r.Context(), &req)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create sale: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sale)
}

func (h *SaleHandler) GetSale(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid sale ID", http.StatusBadRequest)
		return
	}

	sale, err := h.db.GetSale(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sale)
}
