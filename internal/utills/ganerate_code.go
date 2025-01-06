package utills

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

func GenerateRoomCode(creatorID uuid.UUID) (code string) {
	code = fmt.Sprintf("%d", creatorID.ID())

	if len(code) > 7 {
		code = code[:7]
	}

	if len(code) != 7 {
		code = ""
		for i := 0; i < 7; i++ {
			d := fmt.Sprintf("%d", rand.Int31n(9))
			code += d
		}
		return code
	}

	return
}
