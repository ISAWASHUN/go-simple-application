package model

import "errors"

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}