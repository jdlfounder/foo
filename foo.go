package foo

import "fmt"

func Greet(name string) string {
	//return fmt.Sprintf("%s, 你好！", name)
	return fmt.Sprintf("%s, 你好！ Version v0.2.0", name)
}
