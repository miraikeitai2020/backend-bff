package model

type ArticlesResponse struct {
	ArticleList []*ArticleHeader `json:"articleList"`
}

type AddLikeResponse struct {
	Nice int `json:"nice"`
}

type SpotResponse struct {
	Spot SpotInfo `json:"spot"`
}

type SpotInfo struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type DetourResponse struct {
	Detour []DetourInfo `json:"detours"`
}

type DetourInfo struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Image		string	`json:"image"`
	Description	string	`json:"description"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type MutationResponse struct {
	Status	bool	`json:"status"`
}
