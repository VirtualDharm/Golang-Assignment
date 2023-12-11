import React, { useState } from 'react';
import axios from 'axios';

const EditTaskForm = ({ title }) => {
  const [updatedTitle, setUpdatedTitle] = useState('');
  const [updatedDescription, setUpdatedDescription] = useState('');
  const [updatedPriority, setUpdatedPriority] = useState('');

  const editTask = () => {
    axios.put(`http://localhost:8080/edit-task/${title}`, { title: updatedTitle, description: updatedDescription, priority: updatedPriority }) // Update the endpoint
      .then(response => console.log('Task edited successfully:', response.data))
      .catch(error => console.error('Error editing task:', error));
  };

  return (
    <div>
      <h2>Edit Task</h2>
      <form onSubmit={editTask}>
        <label>New Title:</label>
        <input type="text" value={updatedTitle} onChange={e => setUpdatedTitle(e.target.value)} />

        <label>New Description:</label>
        <input type="text" value={updatedDescription} onChange={e => setUpdatedDescription(e.target.value)} />

        <label>New Priority:</label>
        <input type="text" value={updatedPriority} onChange={e => setUpdatedPriority(e.target.value)} />

        <button type="submit">Edit Task</button>
      </form>
    </div>
  );
};

export default EditTaskForm;
