package token

import (
	"errors"
	"log"
	"time"

	usermgr "github.com/Ankr-network/dccn-common/protos/usermgr/v1/micro"
	jwt "github.com/dgrijalva/jwt-go"

	ankr_default "github.com/Ankr-network/dccn-common/protos"
)

var secret = []byte("14444749c1ecc982cd0f91113db98211")

type IToken interface {
	New(user *usermgr.User) (string, error)
	Verify(tokenString string) error
}

type Token struct {
	RefreshTokenValidTime int
	AccessTokenValidTime  int
}

// UserPayload is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type UserPayload struct {
	user *usermgr.User
	jwt.StandardClaims
}

// New returns Token instance.
func New() *Token {
	return &Token{
		AccessTokenValidTime:  ankr_default.AccessTokenValidTime,
		RefreshTokenValidTime: ankr_default.RefreshTokenValidTime,
	}
}

// New returns JWT string.
func (p *Token) New(user *usermgr.User) (string, error) {

	expireTime := time.Now().Add(time.Minute * time.Duration(p.RefreshTokenValidTime)).Unix()

	// Create the Claims
	payload := UserPayload{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "ankr.network",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign token and return
	return token.SignedString(secret)
}

// Verify a token string into a token object
func (p *Token) Verify(tokenString string) error {

	log.Println("Debug into Verify: ", tokenString)
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &UserPayload{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Validate the token
	if _, ok := token.Claims.(*UserPayload); ok && token.Valid {
		return nil
	}
	return errors.New("token is invalid")
}
