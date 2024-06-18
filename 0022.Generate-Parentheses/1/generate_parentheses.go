package generate_parentheses

func generateParenthesis(n int) []string {
	var output []string

	var backtracking func(parenthesis string, open, close int)
	backtracking = func(parenthesis string, open, close int) {
		if close == n {
			output = append(output, parenthesis)
			return
		}
		if open < n {
			backtracking(parenthesis+"(", open+1, close)
		}
		if open > close {
			backtracking(parenthesis+")", open, close+1)
		}
	}

	backtracking("", 0, 0)

	return output
}
