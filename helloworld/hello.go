package helloworld

const prefix = "Hello, "

//Hello returns a string
func Hello(name string) string {

	if name == "" {
		name = "world"
	}
	return prefix + name
}
