package handler

import (
	"net/http"
	"strings"
	"url-shortner/internal/store"
)

type RedirectHandler struct {
	store store.Store
}

func NewRedirectHandler(s store.Store) *RedirectHandler {
	return &RedirectHandler{
		store: s,
	}
}

func (h *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2️⃣ Extract code from path
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.NotFound(w, r)
		return
	}

	originalURL, ok := h.store.Get(code)
	if !ok {
		http.NotFound(w, r)
		return
	}

	// 4️⃣ Redirect
	http.Redirect(w, r, originalURL, http.StatusFound)
}
