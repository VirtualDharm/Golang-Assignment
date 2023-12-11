package main

import (
	"fmt"
	// "net/http"
	"server/task_microservice"
	"server/user_microservice"
	"time"
	"bufio"
	"os"
	"strings"
)

// func handleRequests(w http.ResponseWriter, r *http.Request) {
//     // Print the message to the console
//     fmt.Println("Request received from:", r.RemoteAddr)

//     // Write the message to the response writer
//     fmt.Fprintf(w, "Hello, this is your microservices server!")
// }

// Function to add a user
func addUser(userMicroservice *user_microservice.UserMicroservice, adminUser user_microservice.User) {
	fmt.Println("Enter user details:")
	var username, email, userType string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("User Type (admin/default): ")
	fmt.Scan(&userType)

	var userTypeEnum user_microservice.UserType
	switch userType {
	case "admin":
		userTypeEnum = user_microservice.AdminUser
	case "default":
		userTypeEnum = user_microservice.DefaultUser
	default:
		fmt.Println("Invalid user type. Setting user type to default.")
		userTypeEnum = user_microservice.DefaultUser
	}

	err := userMicroservice.AddUser(adminUser, username, email, userTypeEnum)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Function to delete a user
func deleteUser(userMicroservice *user_microservice.UserMicroservice, adminUser user_microservice.User) {
	fmt.Print("Enter username to delete: ")
	var usernameToDelete string
	fmt.Scan(&usernameToDelete)
	userMicroservice.DeleteUser(adminUser, usernameToDelete)
}

// Function to fetch user information
func fetchUserInfo(userMicroservice *user_microservice.UserMicroservice, adminUser user_microservice.User) {
	fmt.Print("Enter username to fetch information: ")
	var usernameToFetch string
	fmt.Scan(&usernameToFetch)
	user, err := userMicroservice.GetUser(adminUser, usernameToFetch)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User found:", *user)
	}
}

// Function to add a task
func addTask(taskMicroservice *task_microservice.TaskMicroservice, userMicroservice *user_microservice.UserMicroservice, adminUser user_microservice.User) {
	fmt.Println("Enter task details:")
	var taskTitle, taskDescription, dueDateString, getUserDetails string
	var priority task_microservice.Priority
	var dueDate time.Time

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Task Title: ")
	taskTitle, _ = reader.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)
	fmt.Print("Task Description: ")
	taskDescription, _ = reader.ReadString('\n')
	taskDescription = strings.TrimSpace(taskDescription)
	fmt.Print("Priority (High/Medium/Low): ")
	fmt.Scan(&priority)
	fmt.Print("Due Date (YYYY-MM-DD): ")
	fmt.Scan(&dueDateString)
	fmt.Print("Enter Existing Username: ")
	getUserDetails, _ = reader.ReadString('\n')
	getUserDetails = strings.TrimSpace(getUserDetails)

	// Parse the dueDate string into time.Time
	dueDate, err := time.Parse("2006-01-02", dueDateString)
	if err != nil {
		fmt.Println("Error parsing Due Date:", err)
		return
	}

	assignedUser, err := userMicroservice.GetUser(adminUser, getUserDetails)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = taskMicroservice.AddTask(adminUser, *assignedUser, taskTitle, taskDescription, priority, dueDate)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Function to edit a task
func editTask(taskMicroservice *task_microservice.TaskMicroservice, adminUser user_microservice.User) {
	reader := bufio.NewReader(os.Stdin)

	var taskTitleToEdit string
	fmt.Print("Enter task title to edit: ")
	taskTitleToEdit, _ = reader.ReadString('\n')
	taskTitleToEdit = strings.TrimSpace(taskTitleToEdit)

	// Get the existing task
	existingTask, err := taskMicroservice.GetTask(adminUser, taskTitleToEdit)
	if err != nil {
		fmt.Println("Error fetching task:", err)
		return
	}

	fmt.Println("taskTitleToEdit : ", taskTitleToEdit, " existingTask : ", existingTask)

	// Take user input for the updated task details
	var editedTaskTitle, editedDescription, editedDueDateString string
	var editedPriority task_microservice.Priority
	var editedDueDate time.Time

	fmt.Print("Enter edited task title: ")
	editedTaskTitle, _ = reader.ReadString('\n')
	editedTaskTitle = strings.TrimSpace(editedTaskTitle)
	fmt.Print("Enter edited task description: ")
	editedDescription, _ = reader.ReadString('\n')
	editedDescription = strings.TrimSpace(editedDescription)
	fmt.Print("Enter edited task priority (High/Medium/Low): ")
	fmt.Scan(&editedPriority)
	fmt.Print("Enter edited task due date (YYYY-MM-DD): ")
	fmt.Scan(&editedDueDateString)

	// Parse the dueDate string into time.Time
	editedDueDate, parseErr := time.Parse("2006-01-02", editedDueDateString)
	if parseErr != nil {
		fmt.Println("Error parsing Due Date:", parseErr)
		return
	}
	fmt.Println("editedDueDate: ", editedDueDate)

	// Create the updated task
	updatedTask := task_microservice.Task{
		Title:       editedTaskTitle,
		Description: editedDescription,
		Priority:    editedPriority,
		DueDate:     editedDueDate,
		Completed:   existingTask.Completed,
		User:        existingTask.User,
	}
	fmt.Println("updatedtask : ", updatedTask)

	// Update the task
	err = taskMicroservice.EditTask(adminUser, taskTitleToEdit, updatedTask)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Task updated successfully.")
	}
}

// Function to fetch a specific task
func fetchTask(taskMicroservice *task_microservice.TaskMicroservice, adminUser user_microservice.User) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter task title to fetch: ")
	var taskTitleToFetch string
	taskTitleToFetch, _ = reader.ReadString('\n')
	taskTitleToFetch = strings.TrimSpace(taskTitleToFetch)
	fetchedTask, err := taskMicroservice.GetTask(adminUser, taskTitleToFetch)
	if err != nil {
		fmt.Println("Error fetching task:", err)
	} else {
		fmt.Println("Fetched Task:", *fetchedTask)
	}
}

// Function to fetch all tasks
func fetchAllTasks(taskMicroservice *task_microservice.TaskMicroservice) {
	allTasks := taskMicroservice.GetAllTasks()
	fmt.Println("All Tasks:", allTasks)
}

// Function to search tasks by priority
func searchTasksByPriority(taskMicroservice *task_microservice.TaskMicroservice, adminUser user_microservice.User) {
	fmt.Print("Enter priority to search (High/Medium/Low): ")
	var searchPriority task_microservice.Priority
	fmt.Scan(&searchPriority)
	mediumPriorityTasks, err := taskMicroservice.SearchTasksByPriority(adminUser, searchPriority)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Tasks with Priority:", searchPriority)
		for _, task := range mediumPriorityTasks {
			fmt.Printf("Title: %s, Assigned To: %s\n", task.Title, task.User.Username)
		}
	}
}

// Function to sort tasks by due date
func sortTasksByDueDate(taskMicroservice *task_microservice.TaskMicroservice, adminUser user_microservice.User) {
	sortedTasks, err := taskMicroservice.SortTasksByDueDate(adminUser)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Tasks sorted by due date:")
		for _, task := range sortedTasks {
			fmt.Printf("Title: %s, Due Date: %s\n", task.Title, task.DueDate.Format("2006-01-02"))
		}
	}
}

// Function to mark a task as complete
func markTaskAsComplete(taskMicroservice *task_microservice.TaskMicroservice, adminUser user_microservice.User) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter task title to mark as complete: ")
	var taskTitleToComplete string
	taskTitleToComplete, _ = reader.ReadString('\n')
	taskTitleToComplete = strings.TrimSpace(taskTitleToComplete)
	var err error
	err = taskMicroservice.MarkTaskAsComplete(adminUser, taskTitleToComplete)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Function to get all users
func getAllUsers(userMicroservice *user_microservice.UserMicroservice, adminUser user_microservice.User) {
	allUsers, err := userMicroservice.GetAllUsers(adminUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All Users:")
	for _, user := range allUsers {
		fmt.Printf("Username: %s, Email: %s, UserType: %s\n", user.Username, user.Email, user.Type)
	}
}


func main() {
	// User Microservice
	adminUser := user_microservice.User{
		Username: "admin_user",
		Email:    "admin@example.com",
		Type:     user_microservice.AdminUser,
	}

	// User Microservice
	userMicroservice := user_microservice.NewUserMicroservice()
	// Task Microservice
	taskMicroservice := task_microservice.NewTaskMicroservice()

	// Set up HTTP routes
    // http.HandleFunc("/check", handleRequests)

	// Start the HTTP server on port 8080
    // go func() {
    //     fmt.Println("Server listening on :8080...")
    //     if err := http.ListenAndServe(":8080", nil); err != nil {
    //         fmt.Println(err)
    //     }
    // }()


	for {
		// Display IVR Menu
		fmt.Println("\nIVR Menu:")
		fmt.Println("1. Adding users")
		fmt.Println("2. Deleting a user")
		fmt.Println("3. Fetching user information")
		fmt.Println("4. Add the task with both adminUser and assignedUser")
		fmt.Println("5. Editing a task")
		fmt.Println("6. Fetching a specific task")
		fmt.Println("7. Fetching all tasks")
		fmt.Println("8. Searching tasks by priority")
		fmt.Println("9. Sorting tasks by due date")
		fmt.Println("10. Marking a task as complete")
		fmt.Println("11. Get all users")
		fmt.Println("0. Exit")

		// Take user input
		var choice int
		fmt.Print("Enter your choice (0-10): ")
		fmt.Scan(&choice)

		// Perform action based on user choice
		switch choice {
		case 1:
			addUser(userMicroservice, adminUser)
		case 2:
			deleteUser(userMicroservice, adminUser)
		case 3:
			fetchUserInfo(userMicroservice, adminUser)
		case 4:
			addTask(taskMicroservice, userMicroservice, adminUser)
		case 5:
			editTask(taskMicroservice, adminUser)
		case 6:
			fetchTask(taskMicroservice, adminUser)
		case 7:
			fetchAllTasks(taskMicroservice)
		case 8:
			searchTasksByPriority(taskMicroservice, adminUser)
		case 9:
			sortTasksByDueDate(taskMicroservice, adminUser)
		case 10:
			markTaskAsComplete(taskMicroservice, adminUser)
		case 11:
			getAllUsers(userMicroservice, adminUser)
		case 0:
			// Exit the program
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 0 and 10.")
		}
	}
}
