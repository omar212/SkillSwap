"use client"

import type { User } from '../types/User';
import axios from 'axios';
import { useEffect, useState } from 'react';

function UserList() {
  const [users, setUsers] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    setIsLoading(true);
    axios
      .get('http://localhost:8080/api/users')
      .then((res) => {
        if (res.status !== 200) throw new Error(`${res.status}: Error`);
        return res.data;
      })
      .then((data) => {
        console.log('data: ', data);
        setUsers(data);
        return data;
      })
      .catch((err: ErrorEvent) => setError(err.message))
      .finally(() => setIsLoading(false));
  }, []);

  if (isLoading) return <main>loading data ...</main>;
  if (error) return <main>{error}</main>;

  return (
    <>
      User List
      {users?.map((user: User) => (
        <ul key={user.id}>
          <li>{user.name}</li>
        </ul>
      ))}
    </>
  );
}

export default UserList;
