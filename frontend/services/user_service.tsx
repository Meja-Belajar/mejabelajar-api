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
  console.log(username, password)
  if(username == null && password == null) {
    const cookies = document.cookie;
    // if(false) {
    if(true) {
      // Parse the cookies and check if the user is logged in
      // If logged in, return the user data
      // Otherwise, return null
      console.log('Checking login');
      return {
        "status": 200,
        "user": {
          "userid": "USER0001",
          "account_type": "USER",
          "email": "agus@bagus.com",
          "account_detail": {
            "name": "Agus Bagus",
            "profile_picture": "https://cdn.discordapp.com/attachments/1211481783082680480/1211481789206495252/temporary-profile-placeholder-1.png?ex=65ee5b56&is=65dbe656&hm=57aa7174806ef6b5e1a04f987b93948e2e9031a85ace1a6d31056c6d70c890f8&", 
            "bod": "2000-05-15T00:00:00.000Z",
            "social": {
              "linkedin": "https://www.linkedin.dummy.com/in/agusbagus",
              "whatsapp": "0812 1212 1212"
            },
            "description": "Hi, I'm BINUSIAN B22"
          },
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