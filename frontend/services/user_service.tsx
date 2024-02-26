import { RegisterUser } from "../models/user_model";

// Fetch live data
export const getLive = async () => {
  try {
    const response = await fetch('../data/live.json');
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

// Fetch mentor data
export const getMentor = async () => {
  try {
    const response = await fetch('../data/user.json');
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

// Fetch user data
export const getUser = async () => {
  try {
    const response = await fetch('../data/user.json');
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

export const loginService = async (username: string | null, password: string | null) => {
  if(username === null && password === null) {
    const cookies = document.cookie;
    if(cookies) {
      // Parse the cookies and check if the user is logged in
      // If logged in, return the user data
      // Otherwise, return null
      return {
        "status": 200,
        "user": {
          "userid": "USER0001",
          "account_type": "USER",
          "email": "agus@bagus.com",
          "account_detail": {},
          "mentor_detail": {},
        }
      }
    } else {
      return {
        "status": 404,
        "user": { }
      };
    }
  } 

  try {
    const response = await fetch('../data/login_success.json', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error', error);
    throw error;  
  }
}

export const registerService = async (registerUser : RegisterUser) => {
  try {
    const response = await fetch('../data/login_success.json', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(registerUser),
    });

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error', error);
    throw error;  
  }
}