package model

type ArticlesResponse struct {
	ArticleList []*ArticleHeader `json:"articleList"`
}

type AddLikeResponse struct {
	Nice int `json:"nice"`
}

