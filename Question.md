Golang Coding Assignment: Building a Task Management System
Objective:
Develop a task management system that consists of two microservices:
User microservices
Task management microservice

Deadline:
Please complete this assignment within the given timeline.

Requirements:
User Microservices:
Create a user struct with the following fields:
Username
Email
Type (default or admin)

OR

Create a user in user microservice with the info 
{ 
username, 
email, 
type // default and admin 
}
Functions:
Admin should be able to add/delete users.
Default users should be able to fetch their own information only.


Task Management Microservice:
Create a task struct with the following fields:
Title
Description
Priority
Due date
Functions:
Admins should be able to create/edit tasks.
Admins should be able to search for tasks and sort them by completion status, due date, and priority.
Users should only be able to fetch their own tasks and mark them as complete.
Events:
When a task is marked as complete, raise an event for the system.

Additional Instructions:
Choose any database of your liking.
Write unit and integration tests for at least some part of one microservice.
Do not build an authentication system within the microservices.