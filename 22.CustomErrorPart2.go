package main
import (
	"fmt"
	"strings"
)
type PostError struct {
	Code    int
	Message string
}
func (e *PostError) Error() string {
	return fmt.Sprintf("Error %d %s", e.Code, e.Message)
}
// Validate a post and return error if it is invalid
func validatePost(content string) error {
	//Check for emplt content
	if strings.TrimSpace(content) == "" {
		return &PostError{1001, "post can not be empty"}
	}

	//check for to many hashtags
	hashtags := strings.Count(content, "#")

	if hashtags > 5 {
		return &PostError{1002, "Maximum hashtag 5 allowd"}
	}
	return nil //valid post
}
func main() {
	//making instance
	err := validatePost("")

	if err != nil {
		fmt.Println(err)
	}

	err = validatePost(" Go lanf is awesome. #code #go #coder #insta #newskill #c++")
	if err != nil {
		fmt.Println(err)
	}

}
