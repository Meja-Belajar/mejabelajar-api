import { userServiceApi } from "@src/configs/envConfig";
import { LoginUserRequest, RegisterUserRequest } from "@src/models/requests/user_request";
import Cookies from 'universal-cookie';
import { jwtDecode } from 'jwt-decode';
import { decode } from "punycode";
import { LoginUserResponse, RegisterUserResponse } from "@src/models/responses/user_response";

const cookies = new Cookies();

export const registerService = async (requestData: RegisterUserRequest): Promise<RegisterUserResponse> => {
  const { user_name, email, password, phone_number, bod, confirm_password, created_by } = requestData;
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
    
    if(registerResponse.code !== 200) {
      throw new Error(registerResponse.message);
    }
    
    return registerResponse as RegisterUserResponse;
  } catch (error) {

    console.error('Error registering user:', error);
    throw new Error('Failed to register user. You may have entered used data. Please try again.');
  }
}

export const loginService = async ( requestData : LoginUserRequest): Promise<LoginUserResponse> => {
  const { email, password } = requestData;
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
    
    if(loginResponse.code !== 200) {
      throw new Error(loginResponse.message);
    }

    // Handle specific error cases here if necessary
    
    // Decode and manage JWT token
    // const decoded = jwtDecode(loginResponse.data.token);
    // cookies.set('token', loginResponse.data.token, { 
    //   expires: new Date(decoded.exp! * 1000) 
    // });

    // Store user data in local storage
    localStorage.setItem('user', JSON.stringify(loginResponse.data));

    return loginResponse as LoginUserResponse;
  } catch (error) {

    console.error('Error logging in:', error);
    throw new Error('Failed to login. Please try again.');
  }
}


export const isLoggedService = () => {
  // return cookies.get('token');

  if(localStorage.getItem('user')) {
    return JSON.parse(localStorage.getItem('user')!);
  } else {
    return null;
  }
}

export const logoutService = () => {
  // cookies.remove('token');
  // return null;

  if(localStorage.getItem('user')){
    localStorage.removeItem('user');
  }
  return null;
}
