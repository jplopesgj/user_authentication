package services

type AuthService interface {
	GenerateJWT(email string, username string) (tokenString string, err error)
	ValidateToken(signedToken string) (err error)
}
