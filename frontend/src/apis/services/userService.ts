import { userServiceApi } from "@src/configs/envConfig";
import { LoginUserRequest, RegisterUserRequest } from "@src/models/requests/user_request";
import Cookies from 'universal-cookie';
import { jwtDecode } from 'jwt-decode';
import { decode } from "punycode";

const cookies = new Cookies();

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
    
    // const decoded = jwtDecode(loginResponse.data.token);
    
    // jwt token
    // cookies.set('token', loginResponse.data.token, { 
      // expires: new Date(decoded.exp! * 1000) 
    // });

    localStorage.setItem('user', JSON.stringify(loginResponse.data));

    return loginResponse;
    
  } catch (error) {
    
    throw error;  
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