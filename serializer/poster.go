package serializer

import (
	"Waymon_api/model"
)

type Poster struct {
	Img string `json:"img"`
}

func BuildPoster(poster model.Poster) Poster {
	return Poster{
		Img: poster.Img,
	}
}

func BuildPosters(items []model.Poster) (posters []Poster) {
	for _, item := range items {
		poster := BuildPoster(item)
		posters = append(posters, poster)
	}
	return
}
