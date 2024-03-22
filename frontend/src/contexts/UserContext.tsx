import React, { createContext, useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom';
import Loading from '@src/components/Loading';
import { LoginUserResponse, UserAsMentorResponse } from '@src/models/responses/user_response';
import { isLoggedService } from '@src/apis/services/userService';

interface Children {
  children: React.ReactNode
}

interface UserContextProps {
  user?: LoginUserResponse;
  setUser?: (c: LoginUserResponse) => void;
}

const UserContext = createContext<UserContextProps>({})

const UserProvider = ( { children } : Children ) => {
  const [isLoading, setLoading] = useState<boolean>(false);

  const [user, setUser] = useState<LoginUserResponse>();
  const [mentor, setMentor] = useState<UserAsMentorResponse>();

  useEffect(() => {
    const isLogged = isLoggedService();

    if(isLogged) {
      setUser(isLogged);

      // getUserAsMentorService();
    }

  }, [user]);

  const location = useLocation();

  useEffect(() => {
    return () => {
      window.scrollTo(0, 0);
    }
  }, [location]);

  if(isLoading) return ( <Loading /> )
  else return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}  
    </UserContext.Provider>
  )
}

export {
  UserContext,
  UserProvider
}
