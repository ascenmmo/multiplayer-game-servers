package adminclient

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/assets/pages"

const (
	Ru  = "ru"
	Eng = "eng"
)

func DevDocs(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuDocs
	case Eng:
		return pages.EngDocs
	default:
		return pages.EngDeveloperInfo
	}
}

func Auth(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuAuth
	case Eng:
		return pages.EngAuth
	default:
		return pages.EngDeveloperInfo
	}
}

func GameCollection(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuGameCollection
	case Eng:
		return pages.EngGameCollection
	default:
		return pages.EngDeveloperInfo
	}
}

func MainPage(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuMainPage
	case Eng:
		return pages.EngMainPage
	default:
		return pages.EngDeveloperInfo
	}
}

func GameInfo(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuGameInfo
	case Eng:
		return pages.EngGameInfo
	default:
		return pages.EngDeveloperInfo
	}
}

func GameConfig(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuGameConfig
	case Eng:
		return pages.EngGameConfig
	default:
		return pages.EngDeveloperInfo
	}
}

func DeveloperInfo(lEngudage string) string {
	switch lEngudage {
	case Ru:
		return pages.RuDeveloperInfo
	case Eng:
		return pages.EngDeveloperInfo
	default:
		return pages.EngDeveloperInfo
	}
}
