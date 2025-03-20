package storage

import "github.com/lat1992/thething/structs"

type MusicService interface {
	InitPlayList()
	GetPlayList() []structs.Music
}

type Music struct {
	playList []structs.Music
}

func New() *Music {
	return &Music{}
}

func (m *Music) InitPlayList() {
	m.playList = []structs.Music{}
}

func (m *Music) GetPlayList() []structs.Music {
	return m.playList
}
