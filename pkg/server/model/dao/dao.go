package dao

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type apiPath struct {
	Auth       string `envconfig:"AUTH_API"`
	User       string `envconfig:"USER_API"`
	Record     string `envconfig:"RECORD_API" default:""`
	Memory     string `envconfig:"MEMORY_API" default:""`
	Collection string `envconfig:"COLLECTION_API"`
	Log        string `envconfig:"LOG_API"`
	Spot       string `envconfig:"SPOT_API" default:""`
	Evaluation string `envconfig:"EVALUATION_API" default:""`
}

type AccessResourceRepository interface {
	Request(arg ...string) ([]byte, error)
}

var (
	pathOf apiPath

	// API RESOURCE PATH
	AUTH_API_QUERY_SIGNIN    = "%s/query/signin"
	AUTH_API_MUTATION_SIGNUP = "%s/mutation/signup"

	RECORD_API_QUERY_ARTICLES     = "%s/query/articles"
	RECORD_API_QUERY_TAG_ARTICLES = "%s/query/tag/articles"
	RECORD_API_QUERY_ARTICLE      = "%s/query/article"
	RECORD_API_MUTATION_LIKE      = "%s/mutation/add/like"

	MEMORY_API_MUTATION_COMMENT = "%s/comment/add"
	MEMORY_API_MUTATION_REQUEST = "%s/request/save"

	COLLECTION_API_QUERY_LIST = "%s/list/send"
	COLLECTION_API_QUERY_LIKE = "%s/list/change"

	SPOT_API_QUERY_SPOT          = "%s/query/spot"
	SPOT_API_QUERY_DETOUR        = "%s/query/detour"
	SPOT_API_MUTATION_EVALUATION = "%s/mutation/evaluate/spot"
	SPOT_API_MUTATION_ADD_SPOT   = "%s/mutation/add/spot"
)

func init() {
	if err := envconfig.Process("", &pathOf); err != nil {
		log.Fatal(err)
	}
}

/* Access Sign API */
type sign struct {
}

func MakeSignClient() AccessResourceRepository {
	return &sign{}
}
func (info *sign) Request(arg ...string) ([]byte, error) {
	var path string
	if arg[0] == "GET" {
		path = fmt.Sprintf(AUTH_API_QUERY_SIGNIN, pathOf.Auth)
	}
	if arg[0] == "POST" {
		path = fmt.Sprintf(AUTH_API_MUTATION_SIGNUP, pathOf.Auth)
	}
	request, err := http.NewRequest(
		arg[0],
		path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("UserID", arg[1])
	request.Header.Set("Password", arg[2])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

/* Access User API */
