package model

type VinylRecord struct {
	ID        uint    `json:"ID"`
	AlbumName string  `json:"albumName"`
	Artist    string  `json:"artist"`
	Year      uint16  `json:"year"`
	Price     float64 `json:"price"`
}
