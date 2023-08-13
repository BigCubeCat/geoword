package server

type WordDTO struct {
	Word string `json:"word"`
}

type AddressDTO struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Address   string `json:"address"`
}
