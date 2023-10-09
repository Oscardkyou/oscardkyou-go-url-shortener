
package models

// URL represents a shortened URL model
type URL struct {
    ID          string `json:"id"`
    OriginalURL string `json:"original_url"`
    ShortURL    string `json:"short_url"`
}
