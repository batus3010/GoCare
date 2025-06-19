package jwt

import (
	"GoCare/components/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	PayLoad tokenprovider.TokenPayLoad `json:"pay_load"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayLoad, expiry int) (*tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaims{
		PayLoad: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayLoad, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return token
	return &claims.PayLoad, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
