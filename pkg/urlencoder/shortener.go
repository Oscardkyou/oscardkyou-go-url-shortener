package urlencoder

import (
	"crypto/sha256"
	"encoding/hex"
)


func (us *UrlShortener) saveToStorage(shortUrl string, longUrl string) {
    file, err := os.OpenFile(us.store, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if _, err := file.WriteString(shortUrl + "::::" + longUrl + "\n"); err != nil {
        log.Fatal(err)
    }
}

func (us *UrlShortener) loadFromStorage(shortUrl string) string {
    file, err := os.Open(us.store)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "::::")
        if parts[0] == shortUrl {
            return parts[1]
        }
    }
    return ""
}

type UrlShortener struct {
	store map[string]string
}

func New() *UrlShortener {
	return &UrlShortener{store: "/mnt/data/go-url-shortener-reuploaded/go-url-shortener-dev/data/url_storage.txt"}
}

func (us *UrlShortener) Encode(longUrl string) string {
	hash := sha256.Sum256([]byte(longUrl))
	shortUrl := hex.EncodeToString(hash[:6])  // берем первые 3 байта для короткого URL
	us.store[shortUrl] = longUrl
	return shortUrl
}

func (us *UrlShortener) Decode(shortUrl string) (string, bool) {
	longUrl, exists := us.store[shortUrl]
	return longUrl, exists
}
