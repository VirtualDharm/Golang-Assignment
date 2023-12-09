import React from 'react';
import AddUser from './components/AddUser';
import TaskList from './components/TaskList';

const App = () => {
  const adminUser = {
    Username: 'admin_user',
    Email: 'admin@example.com',
    Type: 'admin',
  };

  const currentUser = {
    Username: 'john_doe1', // Change this to the logged-in user's username
    Email: 'john1@example.com',
    Type: 'default',
  };

  return (
    <div>
      <h1>Task Management System</h1>
      <AddUser adminUser={adminUser} />
      <TaskList currentUser={currentUser} />
    </div>
  );
};

export default App;
