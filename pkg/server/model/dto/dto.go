package dto

type ClimeLoader struct {
	claims	map[string]interface{}
}

func MakeClimeLoader(claims map[string]interface{}) *ClimeLoader {
	return &ClimeLoader{claims: claims}
}

func (c *ClimeLoader) Load(key string) string {
	return c.claims[key].(string)
}
