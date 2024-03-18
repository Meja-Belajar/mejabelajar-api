import { UserContext } from "@src/contexts/UserContext";
import { useContext } from "react";
import GuestLanding from "./GuestLandingPage";
import UserLanding from "./UserLandingPage";

const Landing = () => {
  const { login } = useContext(UserContext);
  
  const roles: string = 'Guest';

  // landing page controller based on user roles
  switch(roles){
    case 'User':
      return <UserLanding />
    case 'Mentor':
      return <></>
    default:
      return <GuestLanding />
  }
}


export default Landing