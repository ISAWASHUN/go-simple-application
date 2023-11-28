package usecase

import (
	"section2/model"
	"section2/repository"
)

type TaskUsecase interface {
	CreateTask(title string) (int, error)
	GetTasks(id int) (*model.Task, error)
	UpdateTask(id int, title string) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	r repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{r: r}
}

func (u *taskUsecase) CreateTask(title string) (int, error) {
	task := model.Task{Title: title}
	err := task.Validate()
	if err != nil {
		return 0, err
	}


	id, err := u.r.Create(&task)
	return id, err
}

func (u *taskUsecase) GetTasks(id int) (*model.Task, error) {
	t, err := u.r.Read(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (u *taskUsecase) UpdateTask(id int, title string) error {
	task := model.Task{ID: id, Title: title}
	err := u.r.Updated(&task)
	return err
}

func (u *taskUsecase) DeleteTask(id int) error {
	err := u.r.Delete(id)
	return err
}