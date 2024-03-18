export const validateLogin = (email: string, password: string) => {
  if(email === '' || password === '') {
    return false;
  }

  if(email.length < 5 || password.length < 8) {
    return false;
  }
  
  return true;
}