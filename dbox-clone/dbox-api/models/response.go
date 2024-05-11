package models

type Response struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Id          string `json:"id"`
	Folders     []Folder
	AccessToken string `json:"access_token"`
}
