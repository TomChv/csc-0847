package image

type Thumbnail struct {
	*image
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		image: newImage(),
	}
}

func (t *Thumbnail) Resize(h, w uint) error {
	return t.Mw.ThumbnailImage(h, w)
}
