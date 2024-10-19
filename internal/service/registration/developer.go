package registration

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"time"
)

type developerService struct {
	developerStorage storage.DeveloperStorage
	token            tokengenerator.TokenGenerator

	logger *zerolog.Logger
}

func (c *developerService) SignUp(ctx context.Context, developer types.Developer) (token string, refresh string, err error) {
	developer.ID = uuid.NewMD5(uuid.NameSpaceX500, []byte(developer.Email))

	developer.Password = c.token.PasswordHash(developer.Password)

	err = c.developerStorage.CreateDeveloper(developer)
	if err != nil {
		return
	}

	token, err = c.token.GenerateToken(tokentype.Info{
		UserID: developer.ID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	refresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: developer.ID,
		TTL:    time.Hour * 24,
	}, tokengenerator.JWT)

	if err != nil {
		return
	}

	return token, refresh, nil
}

func (c *developerService) SignIn(ctx context.Context, developer types.Developer) (token, refresh string, err error) {
	developer.Password = c.token.PasswordHash(developer.Password)

	developer, err = c.developerStorage.FindByPassword(developer.Email, developer.Nickname, developer.Password)
	if err != nil {
		return
	}

	token, err = c.token.GenerateToken(tokentype.Info{
		UserID: developer.ID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	refresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: developer.ID,
		TTL:    time.Hour * 24 * 30,
	}, tokengenerator.JWT)

	if err != nil {
		return
	}

	return token, refresh, nil
}

func (c *developerService) RefreshToken(ctx context.Context, token string, refresh string) (newToken, newRefresh string, err error) {
	info, err := c.token.ParseToken(refresh)
	if err != nil {
		return
	}

	developer, err := c.developerStorage.FindByID(info.UserID)
	if err != nil {
		return
	}

	newToken, err = c.token.GenerateToken(tokentype.Info{
		UserID: developer.ID,
		TTL:    time.Minute * 15,
	}, tokengenerator.JWT)

	newRefresh, err = c.token.GenerateToken(tokentype.Info{
		UserID: developer.ID,
		TTL:    time.Hour * 24 * 30,
	}, tokengenerator.JWT)

	return newToken, newRefresh, nil
}

func (c *developerService) GetDeveloper(ctx context.Context, token string) (developer types.Developer, err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return developer, err
	}

	developer, err = c.developerStorage.FindByID(info.UserID)
	if err != nil {
		return developer, err
	}

	developer.Password = ""

	return developer, nil
}

func (c *developerService) UpdateDeveloper(ctx context.Context, token string, developer types.Developer) (err error) {
	info, err := c.token.ParseToken(token)
	if err != nil {
		return err
	}

	developer.Password = c.token.PasswordHash(developer.Password)

	developer.ID = info.UserID

	err = c.developerStorage.Update(developer)
	if err != nil {
		return err
	}

	return nil
}

func NewDeveloperService(developerStorage storage.DeveloperStorage, token tokengenerator.TokenGenerator, logger *zerolog.Logger) multiplayer.Developers {
	return &developerService{developerStorage: developerStorage, token: token, logger: logger}
}
