package model

type Article struct {
	Info   *ArticleInfo `json:"info"`
	Errors []*Errors    `json:"errors"`
}

type ArticleHeader struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	ImagePath string    `json:"ImagePath"`
	Tags      []*string `json:"tags"`
}

type ArticleInfo struct {
	ID         string           `json:"id"`
	Title      string           `json:"title"`
	ImagePath  string           `json:"imagePath"`
	Nice       int              `json:"nice"`
	Context    string           `json:"context"`
	UserStatus *ArticleUserInfo `json:"userStatus"`
	Comment    []*Comment       `json:"comment"`
}

type ArticleUserInfo struct {
	Nice bool `json:"nice"`
	List bool `json:"list"`
}

type Articles struct {
	Articles []*ArticleHeader `json:"articles"`
	Errors   []*Errors        `json:"errors"`
}

type Comment struct {
	Name    string `json:"name"`
	Image   string `json:"image"`
	Comment string `json:"comment"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Detour struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Description string  `json:"Description"`
	Locate      *Locate `json:"locate"`
}

type Errors struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type Genres struct {
	Genres []string  `json:"genres"`
	Errors []*Errors `json:"errors"`
}

type Like struct {
	Status bool      `json:"status"`
	Errors []*Errors `json:"errors"`
}

type List struct {
	Articles []*ArticleHeader `json:"articles"`
	Errors   []*Errors        `json:"errors"`
}

type Locate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Log struct {
	Log    *LogInfo  `json:"log"`
	Errors []*Errors `json:"errors"`
}

type LogData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type LogInfo struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Date          string    `json:"date"`
	Worktime      int       `json:"worktime"`
	Concentration []float64 `json:"concentration"`
}

type Logs struct {
	Logs   []*LogData `json:"logs"`
	Errors []*Errors  `json:"errors"`
}

type Result struct {
	Status bool      `json:"status"`
	Errors []*Errors `json:"errors"`
}

type Spot struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Locate *Locate `json:"locate"`
}

type Spots struct {
	Spot   *Spot     `json:"spot"`
	Detour []*Detour `json:"detour"`
	Errors []*Errors `json:"errors"`
}

type Token struct {
	Value  string    `json:"value"`
	Errors []*Errors `json:"errors"`
}

type User struct {
	Name     string    `json:"name"`
	Birthday *Date     `json:"birthday"`
	Gender   int       `json:"gender"`
	Genre    []*string `json:"genre"`
}

type UserInfo struct {
	Info   *User     `json:"info"`
	Errors []*Errors `json:"errors"`
}
