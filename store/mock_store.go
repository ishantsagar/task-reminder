package store

import "tskrm.com/model"

type MockStore struct {
}

func (s *MockStore) CreateTaskReminder(task *model.Task) (int64, error) {
	return 1, nil
}

func (s *MockStore) UpdateTaskReminder(task *model.Task) error {
	return nil
}

func (s *MockStore) DeleteTaskReminder(id string) error {
	return nil
}

func (s *MockStore) GetTaskReminder(id string) (*model.Task, error) {
	return &model.Task{ID: 1, Title: "Lunch", Description: "Lunch with Johnson"}, nil
}

func (s *MockStore) GetTaskReminders() ([]*model.Task, error) {
	return []*model.Task{
		{ID: 1, Title: "Lunch", Description: "Lunch with Johnson", Priority: 1},
		{ID: 2, Title: "Meeting", Description: "Meeting with David", Priority: 2},
	}, nil
}
