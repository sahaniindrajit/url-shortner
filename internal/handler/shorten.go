package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortner/internal/service"
)

type ShortenHandler struct {
	shortener *service.Shortner
	baseURL   string
}

func NewShortenHandler(s *service.Shortner, baseURL string) *ShortenHandler {

	return &ShortenHandler{
		shortener: s,
		baseURL:   baseURL,
	}
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "" &&
		r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "content type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var req shortenRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
	}

	code, err := h.shortener.Shortner(req.URL)

	if err != nil {

		if errors.Is(err, service.ErrInvalidURL) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return

	}

	resp := shortenResponse{
		ShortURL: h.baseURL + "/" + code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
