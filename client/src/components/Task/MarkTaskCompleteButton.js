import React from 'react';
import axios from 'axios';

const MarkTaskCompleteButton = ({ title }) => {
  const markComplete = () => {
    axios.put(`http://localhost:8080/mark-task-complete/${title}`) // Update the endpoint
      .then(response => console.log('Task marked as complete:', response.data))
      .catch(error => console.error('Error marking task as complete:', error));
  };

  return (
    <button onClick={markComplete}>Mark as Complete</button>
  );
};

export default MarkTaskCompleteButton;
