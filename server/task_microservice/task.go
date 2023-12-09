// task_microservice/task.go
package task_microservice

import (
	"fmt"
	"sort"
	"time"
	"server/user_microservice"
)

type Priority string

const (
	LowPriority    Priority = "low"
	MediumPriority Priority = "medium"
	HighPriority   Priority = "high"
)

type Task struct {
	Title       string
	Description string
	Priority    Priority
	DueDate     time.Time
	Completed   bool
	User        user_microservice.User
}

type TaskMicroservice struct {
	tasks []Task
}

func NewTaskMicroservice() *TaskMicroservice {
	return &TaskMicroservice{}
}

func (tm *TaskMicroservice) AddTask(adminUser user_microservice.User, assignedUser user_microservice.User, title, description string, priority Priority, dueDate time.Time) error {
	if adminUser.Type != user_microservice.AdminUser {
		return fmt.Errorf("only admin can add tasks")
	}

	task := Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		DueDate:     dueDate,
		Completed:   false,
		User:        assignedUser,
	}
	tm.tasks = append(tm.tasks, task)
	fmt.Println("Task added:", task)
	return nil
}

func (tm *TaskMicroservice) EditTask(adminUser user_microservice.User, title string, updatedTask Task) error {
	if adminUser.Type != user_microservice.AdminUser {
		return fmt.Errorf("only admin can edit tasks")
	}
	fmt.Println("check re: ",adminUser, title, updatedTask)

	for i, task := range tm.tasks {
		if task.Title == title {
			updatedTask.User = task.User
			tm.tasks[i] = updatedTask
			fmt.Println("Task edited:", updatedTask)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (tm *TaskMicroservice) GetTask(currentUser user_microservice.User, title string) (*Task, error) {
	if currentUser.Type != user_microservice.AdminUser {
		return nil, fmt.Errorf("unauthorized to fetch this task's information")
	}
	for _, task := range tm.tasks {
		if (task.Title == title) {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("unauthorized to fetch this task's information")
}

func (tm *TaskMicroservice) GetAllTasks() []Task {
	return tm.tasks
}

func (tm *TaskMicroservice) SearchTasksByPriority(adminUser user_microservice.User, priority Priority) ([]Task, error) {
	if adminUser.Type != user_microservice.AdminUser {
		return nil, fmt.Errorf("only admin can search tasks by priority")
	}

	var result []Task
	for _, task := range tm.tasks {
		if task.Priority == priority {
			result = append(result, task)
		}
	}

	return result, nil
}

func (tm *TaskMicroservice) SortTasksByDueDate(adminUser user_microservice.User) ([]Task, error) {
    if adminUser.Type != user_microservice.AdminUser {
        return nil, fmt.Errorf("only admin can sort tasks by due date")
    }

    // Sort tasks
    sortedTasks := make([]Task, len(tm.tasks))
    copy(sortedTasks, tm.tasks)
    sort.Slice(sortedTasks, func(i, j int) bool {
        return sortedTasks[i].DueDate.Before(sortedTasks[j].DueDate)
    })

    return sortedTasks, nil
}



func (tm *TaskMicroservice) MarkTaskAsComplete(currentUser user_microservice.User, title string) error {
	for i, task := range tm.tasks {
		if (currentUser.Type == user_microservice.AdminUser || currentUser.Username == title) && task.Title == title {
			tm.tasks[i].Completed = true
			fmt.Printf("'%s' marked as complete.\n", title)
			return nil
		}
	}
	return fmt.Errorf("unauthorized to mark this task as complete")
}
