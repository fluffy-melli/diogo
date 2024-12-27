package render

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
)

func GIF(urls []string, delay int) (*bytes.Buffer, error) {
	var images []*image.Paletted
	var delays []int
	for _, url := range urls {
		img, err := LoadImageFromURL(url)
		if err != nil {
			return nil, err
		}
		palette := Palette(img)
		palettedImage := image.NewPaletted(img.Bounds(), palette)
		for y := 0; y < palettedImage.Bounds().Dy(); y++ {
			for x := 0; x < palettedImage.Bounds().Dx(); x++ {
				palettedImage.Set(x, y, img.At(x, y))
			}
		}
		images = append(images, palettedImage)
		delays = append(delays, delay)
	}
	var buf bytes.Buffer
	err := gif.EncodeAll(&buf, &gif.GIF{
		Image: images,
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
