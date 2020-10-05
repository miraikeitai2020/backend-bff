package auth

import(
	"fmt"
	"time"
	"io/ioutil"
	_ "crypto/rsa"
    jwt "github.com/dgrijalva/jwt-go"
)

var(
	signKey []byte
)

func init() {
	signKey, _ = ioutil.ReadFile("./demo.rsa")
}

func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["sub"] = id
    claims["iat"] = time.Now()

    return token.SignedString(signKey)
}

func VerifyToken(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, Hmac)
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !(ok && t.Valid) {
		return nil, fmt.Errorf("[ERROR]: Faild get claims")
	}
	return claims, nil
}

func Hmac(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("[ERROR]: Faild Signing Method HMAC")
	}
	return signKey, nil
}
