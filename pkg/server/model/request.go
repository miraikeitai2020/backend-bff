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

type SpotRequest struct {
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
	Walktime	int		`json:"walktime"`
	Emotion		int		`json:"emotion"`
}

type DetourRequest struct {
	SpotLatitude	float64	`json:"spot_latitude"`
	SpotLongitude	float64	`json:"spot_longitude"`
	UserLatitude	float64	`json:"user_latitude"`
	UserLongitude	float64	`json:"user_longitude"`
	Walktime 		int		`json:"walktime"`
	Emotion 		int		`json:"emotion"`
}

type AddSpotRequest struct {
	Name		string	`json:"name"`
	Image		string	`json:"image"`
	Description	string	`json:"description"`
	Latitude	float64	`json:"latitude"`
	Longitude	float64	`json:"longitude"`
}

type AddEvaluationRequest struct {
	ID		string	`json:"id"`
	Emotion	int		`json:"emotion"`
	Status	bool	`json:"status"`
}
