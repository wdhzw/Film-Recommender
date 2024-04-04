package movie_server

import (
	"encoding/json"
)

const (
	ImagePrefix = "https://image.tmdb.org/t/p/original"
)

func FillURI(uri string) string {
	return ImagePrefix + uri
}

type Actor struct {
	Name       string `json:"name"`
	Character  string `json:"character"`
	ProfileUri string `json:"profile_uri"`
}

func FillActorProfileImage(actorsBytes []byte) ([]byte, error) {
	actors := make([]*Actor, 0)
	if err := json.Unmarshal(actorsBytes, &actors); err != nil {
		return nil, err
	}

	for _, actor := range actors {
		if actor.ProfileUri == "" {
			continue
		}
		actor.ProfileUri = FillURI(actor.ProfileUri)
	}

	res, err := json.Marshal(actors)
	if err != nil {
		return nil, err
	}
	return res, nil
}
