import { UserContext } from '@src/contexts/UserContext';
import { useContext, useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';

interface LoginAuthProps {
  mustLogin: boolean;
}

const AuthRedirector = ({ mustLogin }: LoginAuthProps) => {
  const { login } = useContext(UserContext);
  const navigate = useNavigate();

  useEffect(() => {
    // console.log(login?.status, mustLogin);

    // if user already login and want to access authentication page then navigate to home
    if(login?.status === 200 && mustLogin === false) {
      navigate('/');
    } 
  
    // if user not login and want to access protected page then navigate to login
    if(login?.status !== 200 && mustLogin === true) {
      navigate('/login');
    }
  }, []);
  
  return <Outlet />

};

export {
  AuthRedirector
};