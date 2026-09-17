package utility

import (
	"fmt"
	"strings"
)

// CommentEngine provides code comment insertion utilities.
type CommentEngine struct{}

// NewCommentEngine creates a new CommentEngine.
func NewCommentEngine() *CommentEngine {
	return &CommentEngine{}
}

// InsertLineComments inserts single-line comments into code.
func (e *CommentEngine) InsertLineComments(code string, comment string) string {
	lines := strings.Split(code, "\n")
	var result strings.Builder
	for i, line := range lines {
		if i > 0 && i%5 == 0 {
			fmt.Fprintf(&result, "  // %s\n", comment)
		}
		result.WriteString(line)
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}

// InsertBlockComments wraps sections of code in block comments.
func (e *CommentEngine) InsertBlockComments(code string, blockComment string) string {
	lines := strings.Split(code, "\n")
	var result strings.Builder
	for i, line := range lines {
		if i%10 == 0 && i > 0 {
			fmt.Fprintf(&result, "/* %s */\n", blockComment)
		}
		result.WriteString(line)
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}

// HideInComments conceals data inside comment blocks.
func (e *CommentEngine) HideInComments(code string, data string) string {
	comment := fmt.Sprintf(" /* %s */ ", data)
	lines := strings.Split(code, "\n")
	var result strings.Builder
	for i, line := range lines {
		if i > 0 && i%7 == 0 {
			result.WriteString(comment)
			result.WriteString("\n")
		}
		result.WriteString(line)
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}

// DocStringInject inserts data as fake documentation.
func (e *CommentEngine) DocStringInject(code string, data string) string {
	doc := fmt.Sprintf("// %s", data)
	lines := strings.Split(code, "\n")
	var result strings.Builder
	result.WriteString(doc)
	result.WriteString("\n")
	for _, line := range lines {
		result.WriteString(line)
		result.WriteString("\n")
	}
	return result.String()
}