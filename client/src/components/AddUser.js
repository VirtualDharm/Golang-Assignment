import React, { useState } from 'react';
import axios from 'axios';

const AddUser = ({ adminUser }) => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [userType, setUserType] = useState('default');

  const addUser = async () => {
    try {
      await axios.post('http://localhost:8080/add-user', {
        adminUser,
        username,
        email,
        userType,
      });
      console.log('User added successfully!');
    } catch (error) {
      console.error('Error adding user:', error.message);
    }
  };

  return (
    <div>
      <h2>Add User</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="text"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <select value={userType} onChange={(e) => setUserType(e.target.value)}>
        <option value="default">Default</option>
        <option value="admin">Admin</option>
      </select>
      <button onClick={addUser}>Add User</button>
    </div>
  );
};

export default AddUser;
