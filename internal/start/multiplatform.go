package start

import (
	"fmt"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/multiplayer-game-servers/internal/service/access"
	devtools "github.com/ascenmmo/multiplayer-game-servers/internal/service/dev_tools"
	"github.com/ascenmmo/multiplayer-game-servers/internal/service/registration"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	adminclient "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/transport"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"html/template"
)

func Multiplayer(logger zerolog.Logger) {
	logger = logger.With().Str("server:", "multiplayer").Logger()
	client := storage.NewMongoConnection(env.MongoURL)
	token, err := tokengenerator.NewTokenGenerator(env.TokenKey)
	mastNil(err)

	developerStorage, err := storage.NewDeveloperStorage(client)
	mastNil(err)
	clientStorage, err := storage.NewClientStorage(client)
	mastNil(err)
	gameStorage, err := storage.NewGameStorage(client)
	mastNil(err)
	serverStorage, err := storage.NewServersStorage(client)
	mastNil(err)
	accessGameStorage, err := storage.NewAccessGameStorage(client)
	mastNil(err)
	//accessClientStorage, err := storage.NewAccessClientStorage(client)
	//mastNil(err)
	roomStorage, err := storage.NewRoomsStorage(client)
	mastNil(err)
	gameConfigStorage, err := storage.NewGameConfigsStorage(client)
	mastNil(err)

	accessGameService := access.NewAccessGame(accessGameStorage)

	developerService := registration.NewDeveloperService(developerStorage, token, &logger)
	clientService := registration.NewClientService(clientStorage, gameStorage, roomStorage, token, &logger)

	devToolsConnectionServicc := devtools.NewConnections(gameStorage, serverStorage, roomStorage, token, &logger)
	devToolsService := devtools.NewDevTools(accessGameService, gameStorage, serverStorage, token, &logger)
	devToolsServerService := devtools.NewServerService(accessGameService, gameStorage, serverStorage, token, &logger)
	devToolsGameConfigs := devtools.NewGameConfigs(accessGameService, gameConfigStorage, token, &logger)

	services := []transport.Option{
		transport.MaxBodySize(10 * 1024 * 1024),
		transport.Developers(transport.NewDevelopers(developerService)),
		transport.DevToolsClient(transport.NewDevToolsClient(clientService)),
		transport.DevToolsConnections(transport.NewDevToolsConnections(devToolsConnectionServicc)),
		transport.DevTools(transport.NewDevTools(devToolsService)),
		transport.DevToolsServer(transport.NewDevToolsServer(devToolsServerService)),
		transport.DevToolsGameConfigs(transport.NewDevToolsGameConfigs(devToolsGameConfigs)),
	}

	srv := transport.New(logger, services...).WithLog()

	if env.RunAdminPanel {
		app := srv.Fiber()
		adminPanel(app)
	}

	logger.Info().Str("bind", fmt.Sprintf("http://%s:%s", env.ServerAddress, env.MultiplayerPort)).Msg("listen on")
	if err := srv.Fiber().Listen(":" + env.MultiplayerPort); err != nil {
		logger.Panic().Err(err).Stack().Msg("server error")
	}
}

func mastNil(err error) {
	if err != nil {
		panic(err)
	}
}

func adminPanel(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.MainPage("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}
		return nil
	})

	app.Get("/developer/doc", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.DevDocs("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}

		return nil
	})

	app.Get("/admin/auth", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.Auth("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}
		return nil
	})

	app.Get("/admin/games", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.GameCollection("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}
		return nil
	})

	app.Get("/admin/game_info", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.GameInfo("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}
		return nil
	})

	app.Get("/admin/game_info/config", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.GameConfig("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}
		return nil
	})

	app.Get("/admin/info", func(c *fiber.Ctx) error {
		tmpl, err := template.New("docs").Parse(string(adminclient.DeveloperInfo("ru")))
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		err = tmpl.Execute(c.Response().BodyWriter(), dev_doc.GetCategory())
		if err != nil {
			return err
		}
		return nil
	})
}
