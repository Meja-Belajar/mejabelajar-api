import React, { createContext, useEffect, useState } from 'react'
import { getLive, getUser, loginService } from '../../services/user_service';
import { Live, User, UserLogin } from '../../models/user_model';
import Login from '../pages/Login/Login';
import { useLocation } from 'react-router-dom';

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
        const userData = await getUser();
        setUser(userData);

        const liveData = await getLive();
        setlive(liveData);

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
      console.log(data);
      
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
  

  if(isLoad) return ( <div> loading </div> )
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
