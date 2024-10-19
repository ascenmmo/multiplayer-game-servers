package storage

import (
	"encoding/json"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"os"
)

type ServerConfiguration interface {
	Create(conf types.ServerConfiguration) (err error)
	Read() (conf types.ServerConfiguration, err error)
}

type serverConfiguration struct {
	generator tokengenerator.TokenGenerator
}

func NewServerConfiguration(generator tokengenerator.TokenGenerator) ServerConfiguration {
	return &serverConfiguration{
		generator: generator,
	}
}

func (s *serverConfiguration) Create(conf types.ServerConfiguration) (err error) {
	create, err := os.Create("server_configuration.txt")
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	data, err := s.generator.GenerateHash(string(marshal))
	if err != nil {
		return err
	}

	_, err = create.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (s *serverConfiguration) Read() (conf types.ServerConfiguration, err error) {
	file, err := os.ReadFile("server_configuration.txt")
	if err != nil {
		return conf, err
	}

	txt, err := s.generator.ParseHash(string(file))
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal([]byte(txt), &conf)

	return conf, err
}
