package model

type ArticlesRequest struct {
	Genre	string	`json:"genre"`
	Year	*int	`json:"Year"`
	Month	*int	`json:"month"`
}

type ArticlesFromTagResponse struct {
	Tag	string	`json:"tag"`
}

type ArticleRequest struct {
	ID	string	`json:"articleID"`
}
