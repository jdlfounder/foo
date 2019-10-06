package foo

import "fmt"

func Greet(name string) string {
	//return fmt.Sprintf("%s, 你好！", name)
	return fmt.Sprintf("%s, 你好！ Version v1.0.2", name)
}
