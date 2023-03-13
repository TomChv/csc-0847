package image

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

type image struct {
	Mw *imagick.MagickWand
}

func newImage() *image {
	imagick.Initialize()

	return &image{
		Mw: imagick.NewMagickWand(),
	}
}

func (i *image) Terminate() {
	i.Mw.Destroy()
	imagick.Terminate()
}

// Read loads new image in magickWand memory
func (i *image) Read(filename string) error {
	return i.Mw.ReadImage(filename)
}

// Read writes image in local file
func (i *image) Write(filename string) error {
	return i.Mw.WriteImage(filename)
}
