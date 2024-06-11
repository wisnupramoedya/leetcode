package simplify_path

/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Simplify Path.
Memory Usage: 3.1 MB, less than 95.34% of Go online submissions for Simplify Path.
*/

import "strings"

type Stack []string

func (s *Stack) Push(value string) bool {
	*s = append(*s, value)
	return true
}

func (s *Stack) Pop() string {
	newStack := *s
	length := len(newStack)
	value := (newStack)[length-1]
	*s = newStack[0 : length-1]
	return value
}

func simplifyPath(path string) string {
	stack := &Stack{}
	length := len(path)
	prevSlash, nextSlash := 0, 0

	if path[length-1] != '/' {
		path += "/"
	}
	// println(path)
	length = len(path)
	for i := 0; i < length; i++ {
		// println(i, string(path[i]))
		if path[i] == '/' {
			// println("path", i, prevSlash, nextSlash)
			prevSlash = nextSlash
			nextSlash = i

			item := string(path[prevSlash:nextSlash])
			// println(i, item, prevSlash, nextSlash)
			if item == "/" {
				continue
			} else if item == "/." {
				continue
			} else if item == "/.." {
				if len(*stack) != 0 {
					stack.Pop()
				}
				continue
			}
			stack.Push(item)
		}
	}

	newPath := strings.Join(*stack, "")
	if newPath == "" {
		newPath = "/"
	}

	return newPath
}
