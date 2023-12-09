import React, { useEffect, useState } from 'react';
import axios from 'axios';

const TaskList = ({ currentUser }) => {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/get-tasks/${currentUser.Username}`
        );
        setTasks(response.data);
      } catch (error) {
        console.error('Error fetching tasks:', error.message);
      }
    };

    fetchTasks();
  }, [currentUser.Username]);

  return (
    <div>
      <h2>Task List</h2>
      <ul>
        {tasks.map((task) => (
          <li key={task.Title}>{task.Title}</li>
        ))}
      </ul>
    </div>
  );
};

export default TaskList;
