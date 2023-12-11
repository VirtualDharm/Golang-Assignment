import React, { useState } from 'react';
import axios from 'axios';

const AddTaskForm = () => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [priority, setPriority] = useState('');

  const addTask = () => {
    axios.post('http://localhost:8080/add-task', { title, description, priority }) // Update the endpoint
      .then(response => console.log('Task added successfully:', response.data))
      .catch(error => console.error('Error adding task:', error));
  };

  return (
    <div>
      <h2>Add Task</h2>
      <form onSubmit={addTask}>
        <label>Title:</label>
        <input type="text" value={title} onChange={e => setTitle(e.target.value)} />

        <label>Description:</label>
        <input type="text" value={description} onChange={e => setDescription(e.target.value)} />

        <label>Priority:</label>
        <input type="text" value={priority} onChange={e => setPriority(e.target.value)} />

        <button type="submit">Add Task</button>
      </form>
    </div>
  );
};

export default AddTaskForm;
