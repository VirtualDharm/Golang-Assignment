import React, { useState } from 'react';
import axios from 'axios';

const AddUserForm = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [type, setType] = useState('');

  const addUser = () => {
    axios.post('http://localhost:8080/addUser', { username, email, type }) // Update the endpoint
      .then(response => console.log('User added successfully:', response.data))
      .catch(error => console.error('Error adding user:', error));
  };

  return (
    <div>
      <h2>Add User</h2>
      <form onSubmit={addUser}>
        <label>Username:</label>
        <input type="text" value={username} onChange={e => setUsername(e.target.value)} />

        <label>Email:</label>
        <input type="text" value={email} onChange={e => setEmail(e.target.value)} />

        <label>Type:</label>
        <input type="text" value={type} onChange={e => setType(e.target.value)} />

        <button type="submit">Add User</button>
      </form>
    </div>
  );
};

export default AddUserForm;
