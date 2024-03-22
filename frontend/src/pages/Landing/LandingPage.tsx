import { UserContext } from "@src/contexts/UserContext";
import { useContext } from "react";
import GuestLanding from "./GuestLandingPage";
import UserLanding from "./UserLandingPage";

const Landing = () => {
  const { user } = useContext(UserContext);
  
  if(user){
    return <UserLanding />
  }

  return <GuestLanding />
}


export default Landing