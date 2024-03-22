import { userServiceApi } from "@src/configs/envConfig";
import { LoginUserRequest, RegisterUserRequest } from "@src/models/requests/account_request";

export const registerService = async ({ 
    user_name, 
    email, 
    password, 
    phone_number, 
    bod, 
    confirm_password, 
    created_by
  } : RegisterUserRequest) => {

  const apiurl = userServiceApi.register;
  
  try {
  
    const response = await fetch(apiurl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ 
        user_name, 
        email, 
        password,
        phone_number,
        bod,
        confirm_password,
        is_active: false,
        created_by,
      })
    });
    
    const registerResponse = await response.json();

    if(registerResponse && registerResponse.code !== 200) {
      throw new Error(registerResponse.message)
    } 

    return registerResponse;
  } catch(error) {

    console.error('Error', error);
    throw error;  
  }
}

export const loginService = async ({ email, password } : LoginUserRequest) => {
  const apiurl = userServiceApi.login;

  try {
    const response = await fetch(apiurl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    const loginResponse = await response.json();
    
    if(loginResponse && loginResponse.code !== 200) {
      throw new Error(loginResponse.message);
    }
    return loginResponse;
    
  } catch (error) {
    console.error('Error', error);
    throw error;  
  }

}

// Fetch live data
export const getLive = async () => {
  try {
    const response = await fetch('../../data/live.json');
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
    const response = await fetch('../../data/mentor.json');
    const data = response.json();
    return data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

// Fetch user data
export const getUser = async () => {
  try {
    const response = await fetch('../../data/user.json');
    const data = response.json();
    return data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};
