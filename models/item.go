package models

type Item struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ImageUrl string `json:"image_url"`
}