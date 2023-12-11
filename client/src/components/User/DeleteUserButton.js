import React from 'react';
import axios from 'axios';

const DeleteUserButton = ({ username }) => {
  const deleteUser = () => {
    axios.delete(`http://localhost:8080/delete-user/${username}`) // Update the endpoint
      .then(response => console.log('User deleted successfully:', response.data))
      .catch(error => console.error('Error deleting user:', error));
  };

  return (
    <button onClick={deleteUser}>Delete User</button>
  );
};

export default DeleteUserButton;
