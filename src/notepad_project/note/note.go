package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	CreatedDate time.Time `json:"created_date"`
}

func (n Note) Display() {
	fmt.Printf("Title: %s\n, Content: %s\n, Created Date: %s\n", n.Title, n.Body, n.CreatedDate)
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, body string) (*Note, error) {

	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if body == "" {
		return nil, errors.New("body cannot be empty")
	}

	return &Note{
		Title: title,
		Body:  body,
		CreatedDate: time.Now(),
	}, nil
}

func toJson(notes []*Note) (string, error) {

	notesJson, err := json.Marshal(notes)

	if err != nil {
		return "", err
	}

	return string(notesJson), nil
}