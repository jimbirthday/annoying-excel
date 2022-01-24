package models

type Page struct {
	Page  int    `json:"page"`
	Pages int    `json:"pages"`
	Q     string `json:"q,omitempty"`
}
