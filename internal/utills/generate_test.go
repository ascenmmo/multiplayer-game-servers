package utills

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestGenerateRoomCode(t *testing.T) {
	found := map[string]int{}
	for i := 0; i < 1000; i++ {
		//fmt.Println(GenerateRoomCode(uuid.New()))
		code := GenerateRoomCode(uuid.New())
		if len(code) != 7 {
			fmt.Println(code)
			t.Error("mast equals 7")
		}
		d, ok := found[code]
		if ok {
			t.Error("dublicate code", d)
		}
		found[code] += 1
	}
}
