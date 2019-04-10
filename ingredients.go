package d

import (
	"strings"

	c "github.com/liggitt/c/v2"
)

func Ingredients() []string {
	return []string{"eggs", "flour", "sugar", "butter"}
}

func CakeAvailable() bool {
	return !strings.Contains(c.CakeOrDeath("cake"), "death")
}
