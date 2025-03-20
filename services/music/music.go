package storage

import "github.com/lat1992/thething/structs"

type MusicService interface {
	GetPlayList() []structs.Music
}

type Music struct {
}

func New() *Music {
	return &Music{}
}

func (m *Music) GetPlayList() []structs.Music {
	return nil
}
