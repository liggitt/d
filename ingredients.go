package d

import (
	"strings"

	"github.com/liggitt/c"
)

func Ingredients() []string {
	return []string{"eggs", "flour", "sugar", "butter"}
}

func CakeAvailable() bool {
	return !strings.Contains(c.CakeOrDeath(), "death")
}
