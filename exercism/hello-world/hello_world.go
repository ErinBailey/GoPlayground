package greeting

import (
	"fmt"
)

// HelloWorld should say hello to someone based on their name.
func HelloWorld(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
