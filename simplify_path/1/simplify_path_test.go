package simplify_path

import "testing"

func Test71_1_1(t *testing.T) {
	path := "/home/"

	expected := "/home"

	result := simplifyPath(path)
	print("result")

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Test71_1_2(t *testing.T) {
	path := "/home//foo/"

	expected := "/home/foo"

	result := simplifyPath(path)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Test71_1_3(t *testing.T) {
	path := "/home/user/Documents/../Pictures"

	expected := "/home/user/Pictures"

	result := simplifyPath(path)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Test71_1_4(t *testing.T) {
	path := "/../"

	expected := "/"

	result := simplifyPath(path)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Test71_1_5(t *testing.T) {
	path := "/.../a/../b/c/../d/./"

	expected := "/.../b/d"

	result := simplifyPath(path)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Test71_1_6(t *testing.T) {
	path := "/.././../b/c/../d/./"

	expected := "/b/d"

	result := simplifyPath(path)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
