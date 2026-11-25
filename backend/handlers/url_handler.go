package handlers

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/Sahan619/url_shortner/backend/models"
	"github.com/go-chi/chi/v5"
)

var urls []models.URL

func isValidAlias(alias string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]{3,20}$`, alias)
	return matched
}

func CreateURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		OriginalURL string     `json:"originalUrl"`
		CustomAlias string     `json:"customAlias"`
		ExpiresAt   *time.Time `json:"expiresAt"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	parsedURL, err := url.ParseRequestURI(request.OriginalURL)

	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		http.Error(w, "Invalid URL. Please provide a valid HTTP or HTTPS URL.", http.StatusBadRequest)
		return

	}
	if request.ExpiresAt != nil && !request.ExpiresAt.After(time.Now()) {
		http.Error(w, "Expiration date must be in the future", http.StatusBadRequest)
		return
	}

	var shortCode string

	if request.CustomAlias != "" {

		if !isValidAlias(request.CustomAlias) {
			http.Error(w, "Invalid custom alias. Use 3-20 letters, numbers, hyphens or underscores.", http.StatusBadRequest)
			return
		}

		exists := false

		for _, url := range urls {
			if url.ShortCode == request.CustomAlias {
				exists = true
				break
			}
		}

		if exists {
			http.Error(w, "Custom alias is already taken", http.StatusConflict)
			return
		}

		shortCode = request.CustomAlias

	} else {
		shortCode, err = generateShortCode()

		if err != nil {
			http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
			return
		}
	}
	managementToken, err := generateManagementToken()
	if err != nil {
		http.Error(w, "Failed to generate management token", http.StatusInternalServerError)
		return
	}

	newURL := models.URL{
		ID:              len(urls) + 1,
		OriginalURL:     request.OriginalURL,
		ShortCode:       shortCode,
		ShortURL:        "http://localhost:8080/" + shortCode,
		Clicks:          0,
		CreatedAt:       time.Now(),
		ManagementToken: managementToken,
		ExpiresAt:       request.ExpiresAt,
	}

	urls = append(urls, newURL)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"url":             newURL,
		"managementToken": managementToken,
	})

}

func generateShortCode() (string, error) {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for {
		code := make([]byte, 6)

		for i := range code {
			randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
			if err != nil {
				return "", err
			}

			code[i] = characters[randomIndex.Int64()]
		}

		shortCode := string(code)

		exists := false

		for _, url := range urls {
			if url.ShortCode == shortCode {
				exists = true
				break
			}
		}

		if !exists {
			return shortCode, nil
		}
	}
}
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]

	for i := range urls {
		if urls[i].ShortCode == shortCode {
			if urls[i].ExpiresAt != nil && !urls[i].ExpiresAt.After(time.Now()) {
				http.Error(w, "This short URL has expired", http.StatusGone)
				return
			}

			urls[i].Clicks++

			http.Redirect(w, r, urls[i].OriginalURL, http.StatusFound)
			return
		}
	}
	http.Error(w, "Short URL not found", http.StatusNotFound)
}

func GetURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(urls)
}

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	for i, url := range urls {
		if url.ShortCode == shortCode {
			token := r.Header.Get("X-Management-Token")

			if subtle.ConstantTimeCompare(
				[]byte(token),
				[]byte(url.ManagementToken),
			) != 1 {
				http.Error(w, "Invalid management token", http.StatusForbidden)
				return
			}

			urls = append(urls[:i], urls[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Short URL not found", http.StatusNotFound)
}

func generateManagementToken() (string, error) {
	tokenBytes := make([]byte, 32)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(tokenBytes), nil
}
