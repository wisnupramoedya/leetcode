package simplify_path

import "testing"

func Test71_1_1(t *testing.T) {
	path := "/home/"
	output := "/home"

	result := simplifyPath(path)

	if result != output {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test71_1_2(t *testing.T) {
	path := "/home//foo/"
	output := "/home/foo"

	result := simplifyPath(path)

	if result != output {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test71_1_3(t *testing.T) {
	path := "/home/user/Documents/../Pictures"
	output := "/home/user/Pictures"

	result := simplifyPath(path)

	if result != output {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test71_1_4(t *testing.T) {
	path := "/../"
	output := "/"

	result := simplifyPath(path)

	if result != output {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test71_1_5(t *testing.T) {
	path := "/.../a/../b/c/../d/./"
	output := "/.../b/d"

	result := simplifyPath(path)

	if result != output {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test71_1_6(t *testing.T) {
	path := "/.././../b/c/../d/./"
	output := "/b/d"

	result := simplifyPath(path)

	if result != output {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}
