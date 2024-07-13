package app

import (
	"context"
	"time"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/internal/db"
	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	userStore       db.UserStore
	tokenSecret     string
	tokenExpTimeout time.Duration
}

func NewUserService(tokenSecret string, tokenExpTimeout time.Duration, userStore db.UserStore) *UserService {
	return &UserService{
		userStore:       userStore,
		tokenSecret:     tokenSecret,
		tokenExpTimeout: tokenExpTimeout,
	}
}

func (us *UserService) encodeToken(claims api.TokenClaims) (api.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": claims.Issuer,
		"exp": claims.ExpTime,
	})
	tokenStr, err := token.SignedString([]byte(us.tokenSecret))
	if err != nil {
		return "", err
	}
	return api.Token(tokenStr), nil
}

func (us *UserService) decodeToken(tokenStr api.Token) (api.TokenClaims, error) {
	token, err := jwt.Parse(string(tokenStr), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, api.Error("invalid signing method")
		}
		return []byte(us.tokenSecret), nil
	})
	if err != nil {
		return api.TokenClaims{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		iss, err := uuid.Parse(claims["iss"].(string))
		if err != nil {
			return api.TokenClaims{}, err
		}
		return api.TokenClaims{
			Issuer:  iss,
			ExpTime: int64(claims["exp"].(float64)),
		}, nil
	}
	return api.TokenClaims{}, api.Error("claims decode error")
}

func (us *UserService) Login(ctx context.Context, username, password string) (uuid.UUID, api.Token, error) {
	user, err := us.userStore.GetUserByName(ctx, username)
	if err != nil {
		return uuid.Nil, "", err
	}
	if user.Password != password {
		return uuid.Nil, "", api.Error("wrong password")
	}
	token, err := us.encodeToken(api.TokenClaims{Issuer: user.Id, ExpTime: time.Now().Local().Add(us.tokenExpTimeout).Unix()})
	return user.Id, token, err
}

func (us *UserService) ValidateToken(ctx context.Context, token api.Token) (api.TokenClaims, error) {
	claims, err := us.decodeToken(token)
	if err != nil {
		return api.TokenClaims{}, err
	}

	// Check exparation time
	if time.Now().Local().Before(time.Unix(claims.ExpTime, 0)) {
		return claims, nil
	}
	return api.TokenClaims{}, api.Error("token expired")
}

func (us *UserService) Register(ctx context.Context, user *types.User) error {
	if exist, err := us.userStore.DoExist(ctx, user.Username); err != nil || exist {
		if err != nil {
			return err
		}
		return db.ErrUserAlreadyExists
	}
	if err := us.userStore.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.userStore.DeleteUser(ctx, id)
}

func (svc *UserService) GetUser(ctx context.Context, id uuid.UUID) (*types.User, error) {
	u, err := svc.userStore.GetUser(ctx, id)
	return u, err
}
func (svc *UserService) GetUserByName(ctx context.Context, username string) (*types.User, error) {
	u, err := svc.userStore.GetUserByName(ctx, username)
	return u, err
}

func (svc *UserService) Close() error {
	return nil
}
