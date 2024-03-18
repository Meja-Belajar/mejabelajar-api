import { UserContext } from '@src/contexts/UserContext';
import React, { useContext, useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';

const Auth = () => {
  const { login } = useContext(UserContext);
  const navigate = useNavigate();
  
  useEffect(() => {
    if(login && login.status === 200) {
      navigate('/');
    }  
  }, [])

  return <Outlet />;
};

export default Auth;