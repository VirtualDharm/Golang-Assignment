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

// Define printUserInfo outside of the main function
func printUserInfo(user user_microservice.User) {
	fmt.Printf("Username: %s, Email: %s, UserType: %s\n", user.Username, user.Email, user.Type)
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

	// Example: Adding users
	// err := userMicroservice.AddUser(adminUser, "john_doe", "john@example.com", user_microservice.DefaultUser)
	// err1 := userMicroservice.AddUser(adminUser, "john_doe1", "john1@example.com", user_microservice.DefaultUser)
	// err2 := userMicroservice.AddUser(adminUser, "john_doe2", "john2@example.com", user_microservice.AdminUser)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// err3 := userMicroservice.AddUser(adminUser, "john_doe3", "john3@example.com", user_microservice.AdminUser)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// if err1 != nil {
	// 	fmt.Println("Error:", err1)
	// }
	// if err2 != nil {
	// 	fmt.Println("Error:", err2)
	// }
	// if err3 != nil {
	// 	fmt.Println("Error:", err3)
	// }

	// Example: Deleting a user
	// err = userMicroservice.DeleteUser(adminUser, "john_doe")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// err2 = userMicroservice.DeleteUser(adminUser, "john_doe3")
	// if err2 != nil {
	// 	fmt.Println("Error:", err2)
	// }

	// Example: Fetching user information
	// user, err := userMicroservice.GetUser(adminUser, "john_doe1")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("User found:", *user)
	// }
	// user1, err1 := userMicroservice.GetUser(adminUser, "john_doe2")
	// if err1 != nil {
	// 	fmt.Println("Error:", err1)
	// } else {
	// 	fmt.Println("User found:", *user1)
	// }

	// Add the task with both adminUser and assignedUser
	// assignedUser, err := userMicroservice.GetUser(adminUser, "john_doe1")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// err = taskMicroservice.AddTask(adminUser, *assignedUser, "Task 1", "Description 1", task_microservice.MediumPriority, time.Now().AddDate(0, 0, 7))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// assignedUser1, err := userMicroservice.GetUser(adminUser, "john_doe2")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// err = taskMicroservice.AddTask(adminUser, *assignedUser1, "Task 2", "Description 2", task_microservice.MediumPriority, time.Now().AddDate(0, 0, 8))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// Example: Editing a task
	// updatedTask := task_microservice.Task{
	// 	Title:       "Task 1 Changed",
	// 	Description: "Updated Description 1",
	// 	Priority:    task_microservice.MediumPriority,
	// 	DueDate:     time.Now().AddDate(0, 0, 5),
	// }
	// err = taskMicroservice.EditTask(adminUser, "Task 1", updatedTask)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// Example: Fetching a specific task
	// taskTitle := "Task 1 Changed"
	// fetchedTask, err := taskMicroservice.GetTask(adminUser, taskTitle)
	// if err != nil {
	// 	fmt.Println("Error fetching task:", err)
	// } else {
	// 	fmt.Println("Fetched Task:", *fetchedTask)
	// }

	// Example: Fetching all tasks
	// allTasks := taskMicroservice.GetAllTasks()
	// fmt.Println("All Tasks:", allTasks)

	// Example: Searching tasks by priority
	// searchPriority := task_microservice.MediumPriority
	// mediumPriorityTasks, err := taskMicroservice.SearchTasksByPriority(adminUser, searchPriority)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("\nTasks with Priority:", searchPriority)
	// 	for _, task := range mediumPriorityTasks {
	// 		fmt.Printf("Title: %s, Assigned To: %s\n", task.Title, task.User.Username)
	// 	}
	// }

	// Example: Sorting tasks by due date
	// sortedTasks, err := taskMicroservice.SortTasksByDueDate(adminUser)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Tasks sorted by due date:")
	// 	for _, task := range sortedTasks {
	// 		fmt.Printf("Title: %s, Due Date: %s\n", task.Title, task.DueDate.Format("2006-01-02"))
	// 	}
	// }

	// Example: Marking a task as complete
	// taskTitleToComplete := "Task 2"
	// err = taskMicroservice.MarkTaskAsComplete(adminUser, taskTitleToComplete)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

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
			// Adding users
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
		case 2:
			// Deleting a user
			fmt.Print("Enter username to delete: ")
			var usernameToDelete string
			fmt.Scan(&usernameToDelete)
			userMicroservice.DeleteUser(adminUser, usernameToDelete)
		case 3:
			// Fetching user information
			fmt.Print("Enter username to fetch information: ")
			var usernameToFetch string
			fmt.Scan(&usernameToFetch)
			user, err := userMicroservice.GetUser(adminUser, usernameToFetch)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("User found:", *user)
			}
		case 4:
			// Add the task with both adminUser and assignedUser
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
		case 5:
			// Editing a task
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
			fmt.Println("taskTitleToEdit : ",taskTitleToEdit," existingTask : ",existingTask)

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
			fmt.Println("updatedtask : ",updatedTask)

			// Update the task
			err = taskMicroservice.EditTask(adminUser, taskTitleToEdit, updatedTask)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task updated successfully.")
			}

		case 6:
			// Fetching a specific task
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
		case 7:
			// Fetching all tasks
			allTasks := taskMicroservice.GetAllTasks()
			fmt.Println("All Tasks:", allTasks)
		case 8:
			// Searching tasks by priority
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
		case 9:
			// Sorting tasks by due date
			sortedTasks, err := taskMicroservice.SortTasksByDueDate(adminUser)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Tasks sorted by due date:")
				for _, task := range sortedTasks {
					fmt.Printf("Title: %s, Due Date: %s\n", task.Title, task.DueDate.Format("2006-01-02"))
				}
			}
		case 10:
			// Marking a task as complete
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
		case 11:
			// Get all users
			allUsers, err := userMicroservice.GetAllUsers(adminUser)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("All Users:")
				for _, user := range allUsers {
					printUserInfo(user)
				}
			}
		case 0:
			// Exit the program
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 0 and 10.")
		}
	}
}
