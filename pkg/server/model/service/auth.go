package service

import(
	"fmt"
	"time"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/dto"
)

var(
	signKey []byte
)

func init() {
	signKey, _ = ioutil.ReadFile("./private.key")
}

func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["sub"] = id
    claims["iat"] = time.Now()

    return token.SignedString(signKey)
}

func VerifyToken(token string) (*dto.ClimeLoader, error) {
	t, err := jwt.Parse(token, Hmac)
	if err != nil {
		return nil, err
	}
	claims, status := t.Claims.(jwt.MapClaims)
	if !(status && t.Valid) {
		return nil, fmt.Errorf("[ERROR]: Faild get claims")
	}
	return dto.MakeClimeLoader(claims), nil
}

func Hmac(token *jwt.Token) (interface{}, error) {
	if _, status := token.Method.(*jwt.SigningMethodHMAC); !status {
		return nil, fmt.Errorf("[ERROR]: Faild Signing Method HMAC")
	}
	return signKey, nil
}
