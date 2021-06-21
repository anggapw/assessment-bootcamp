package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	err = godotenv.Load
	key = os.Getenv("PASS_KEY")
)

type Service interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(UserID int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": UserID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}
