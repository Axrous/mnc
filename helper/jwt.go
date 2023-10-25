package helper

import (
	"errors"
	"time"

	"github.com/Axrous/mnc/model/domain"
	"github.com/golang-jwt/jwt/v5"
)

type CustomerData struct {
	Id string
	Role string
}

type MyCustomClaims struct {
	CustomerId string
	jwt.RegisteredClaims
}

func CreateToken(customer domain.Customer) string{

	key := "123123123123123"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		CustomerId:              customer.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(30 * time.Minute)},
		},
	})

	stringToken, err := token.SignedString([]byte(key))
	PanicIfError(err)
	return stringToken
}

func CompareToken(tokenString string) (string, error) {

	key := "123123123123123"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unauthorized")
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(key), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["CustomerId"].(string), nil
	} else {
		return "", err
	}
}