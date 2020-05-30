package controller

type QueryCmd struct {
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
	Sort    string `json:"sort"`
	Order   string `json:"order"`
}