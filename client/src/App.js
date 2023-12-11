import React from 'react';
import UserList from './components/User/UserList';
import AddUserForm from './components/User/AddUserForm';
import DeleteUserButton from './components/User/DeleteUserButton'; // Add this line
import TaskList from './components/Task/TaskList'; // Add this line
import AddTaskForm from './components/Task/AddTaskForm'; // Add this line
import EditTaskForm from './components/Task/EditTaskForm'; // Add this line
import MarkTaskCompleteButton from './components/Task/MarkTaskCompleteButton'; // Add this line

const App = () => {
  return (
    <div>
      <h1>Task Management System</h1>
      
      {/* User Components */}
      <h2>User Management</h2>
      <UserList />
      <AddUserForm />
      {/* You can add and use other user-related components here */}
      <DeleteUserButton /> {/* Placeholder for DeleteUserButton */}
      
      {/* Task Components */}
      <h2>Task Management</h2>
      <TaskList />
      <AddTaskForm />
      <EditTaskForm />
      <MarkTaskCompleteButton />
      {/* You can add and use other task-related components here */}
    </div>
  );
};

export default App;
