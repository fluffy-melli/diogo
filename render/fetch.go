package render

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"
	"sync"
)

var FileCache = make(map[string]image.Image)
var mu sync.RWMutex

func LoadImageFromURL(ImageURL string) (image.Image, error) {
	ImageURL = strings.TrimSpace(ImageURL)
	mu.RLock()
	img, exists := FileCache[ImageURL]
	mu.RUnlock()
	if exists {
		return img, nil
	}
	resp, err := http.Get(ImageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image from URL %s: %v", ImageURL, err)
	}
	defer resp.Body.Close()
	img, _, err = image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image from URL %s: %v", ImageURL, err)
	}
	mu.Lock()
	FileCache[ImageURL] = img
	mu.Unlock()
	return img, nil
}
