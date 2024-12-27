package render

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"sort"
	"sync"
)

func GIF(urls []string, delay int, low bool) (*bytes.Buffer, error) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var images map[int]*image.Paletted = make(map[int]*image.Paletted)
	sem := make(chan struct{}, 5)
	for i, url := range urls {
		if low && (i+1)%3 != 0 { // | 15min
			continue
		}
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, url string) {
			defer wg.Done()
			defer func() { <-sem }()
			img, err := LoadImageFromURL(url)
			if err != nil {
				return
			}
			palette := Palette(img)
			palettedImage := image.NewPaletted(img.Bounds(), palette)
			for y := 0; y < palettedImage.Bounds().Dy(); y++ {
				for x := 0; x < palettedImage.Bounds().Dx(); x++ {
					palettedImage.Set(x, y, img.At(x, y))
				}
			}
			mu.Lock()
			images[i] = palettedImage
			mu.Unlock()
		}(i, url)
	}
	wg.Wait()
	var keys []int
	var delays []int
	var result []*image.Paletted
	for key := range images {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		result = append(result, images[key])
		delays = append(delays, delay)
	}
	var buf bytes.Buffer
	err := gif.EncodeAll(&buf, &gif.GIF{
		Image: result,
		Delay: delays,
	})
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func Palette(frame image.Image) color.Palette {
	colorSet := make(map[color.Color]struct{})
	for y := 0; y < frame.Bounds().Dy(); y++ {
		for x := 0; x < frame.Bounds().Dx(); x++ {
			colorSet[frame.At(x, y)] = struct{}{}
		}
	}
	var colors []color.Color
	for c := range colorSet {
		colors = append(colors, c)
	}
	if len(colors) > 256 {
		colors = colors[:256]
	}
	return colors
}
