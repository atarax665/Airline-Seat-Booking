import React, { useState, useEffect } from "react";
import axios from "axios";

const GetUser = () => {
  const [userId, setUserId] = useState("");
  const [user, setUser] = useState(null);
  const [fetchUser, setFetchUser] = useState(false);

  useEffect(() => {
    if (!fetchUser) return;

    axios
      .get(`http://localhost:8080/v1/user?userId=${userId}`)
      .then((response) => {
        console.log("User data: ", response.data);
        setUser(response.data);
        setFetchUser(false);
      })
      .catch((error) => {
        console.error("Error fetching data: ", error);
        setUser(null);
        setFetchUser(false);
      });
  }, [fetchUser]);

  return (
    <div className="absolute top-0 right-0 p-4">
      <h2 className="text-2xl font-bold mb-4 text-center">
        Get User from Database
      </h2>
      <input
        type="text"
        value={userId}
        onChange={(e) => setUserId(e.target.value)}
        placeholder="Enter user ID"
        className="border-2 border-gray-300 rounded-md p-2 mr-2"
      />
      <button
        onClick={() => setFetchUser(true)}
        className="bg-blue-500 text-white rounded-md p-2"
      >
        Get User
      </button>
      {user && (
        <div className="mt-4">
          <h2 className="text-xl font-bold">{user.name}</h2>
        </div>
      )}
    </div>
  );
};

export default GetUser;
