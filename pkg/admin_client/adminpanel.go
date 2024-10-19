package adminclient

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/assets/pages"

const (
	ru  = "ru"
	eng = "eng"
)

func DevDocs(lengudage string) string {
	switch lengudage {
	case ru:
		return pages.Docs
	case eng:
		return pages.Docs
	default:
		return "not found"
	}
}

func Auth(lengudage string) string {
	switch lengudage {
	case ru:
		return pages.Auth
	case eng:
		return pages.Auth
	default:
		return "not found"
	}
}

func GameCollection(lengudage string) string {
	switch lengudage {
	case ru:
		return pages.GameCollection
	case eng:
		return pages.GameCollection
	default:
		return "not found"
	}
}

func MainPage(lengudage string) string {
	switch lengudage {
	case ru:
		return pages.MainPage
	case eng:
		return pages.MainPage
	default:
		return "not found"
	}
}

func GameInfo(lengudage string) string {
	switch lengudage {
	case ru:
		return pages.GameInfo
	case eng:
		return pages.GameInfo
	default:
		return "not found"
	}
}
