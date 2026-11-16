package models

type AdvertisementsList struct {
	Items []Advertisement `json:"items"`
}

type Advertisement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
