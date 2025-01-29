package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// struct tags: built in function in Go, meta data for struct field.

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%v' has the following content: \n\n%v\n\n", note.Title, note.Content)
}

// Since this does not edit note, we don't need pointer
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	// json.Marshal function only works with Public fields!!

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
	// 0644: read and edit permitions to owner and only read to other
	// this method should return nil at the end, but os.WriteFile also returns error if fail.
}