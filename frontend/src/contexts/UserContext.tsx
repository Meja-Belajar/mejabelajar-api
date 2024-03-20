import React, { createContext, useEffect, useState } from 'react'
import { getLive, getMentor, loginService } from '@src/apis/services/userService';
import { Live, User, UserLogin } from '@src/models/database/userModel';
import { useLocation } from 'react-router-dom';
import Loading from '@src/components/Loading';

interface Children {
  children: React.ReactNode
}

interface UserContent {
  user?: User;
  live?: Live;
  login?: UserLogin;
  setLogin?: (c: UserLogin) => void;
}

const UserContext = createContext<UserContent>({})

const UserProvider = ( { children } : Children ) => {
  // loading state
  const [isLoad, setLoad] = useState<boolean>(true);

  const [user, setUser] = useState<User>();
  const [live, setlive] = useState<Live>();

  const [login, setLogin] = useState<UserLogin>();

  useEffect(() => {
    setLoad(true);

    const fetchData = async () => {
      try {
        // const userData = await getMentor();
        // setUser(userData);

        // const liveData = await getLive();
        // setlive(liveData);

        // console.log(userData); 
        // console.log(liveData); 

        setLoad(false); 
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);

  useEffect(() => {
    
    const checkLogin = async () => {
      setLoad(true);
      const data = await loginService(null, null);
      
      setLogin(data);
      
      setLoad(false);
    }
    
    checkLogin();
  }, []);

  const location = useLocation();

  useEffect(() => {
    return () => {
      window.scrollTo(0, 0);
    }
  }, [location]);
  

  if(isLoad) return ( <Loading /> )
  else return (
    <UserContext.Provider value={{ user, live, login, setLogin}}>
      {children}  
    </UserContext.Provider>
  )
}

export {
  UserContext,
  UserProvider
}
