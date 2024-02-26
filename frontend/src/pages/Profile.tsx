import { motion } from 'framer-motion'
import { animate, exit, initial } from '../assets/PageTransition'
import { useNavigate } from 'react-router-dom';
import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../context/UserContext';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCamera, faSignOut } from '@fortawesome/free-solid-svg-icons';
import { Button, Input, Textarea } from '@nextui-org/react';

const Profile = () => {
  const { login } = useContext(UserContext);

  const navigate = useNavigate();
  
  useEffect(() => {
    // If user is already logged in, redirect to home
    if(!login || login?.status !== 200) {
      navigate('/');
    }  
  }, [])
  
  const [name, setName] = useState<string>(login?.user?.account_detail?.name || '');
  const [email, setEmail] = useState<string>(login?.user?.email || ''); 
  const [profile, setProfile] = useState<string>(login?.user?.account_detail?.profile_picture || '');
  const [description, setDescription] = useState<string>(login?.user?.account_detail?.description || '');

  const handleImageChange = async (e: any) => {

  }

  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='w-full p-4 mt-5'
      >

        <div className='ml-5 flex flex-col sm:flex-row items-start'>
          <div className='flex flex-col items-center p-2'>
            <h1 className='lato-bold mb-3 sm:hidden'>Profile Image</h1>
            
            <div className='group w-20 sm:w-40 rounded-full relative cursor-pointer sm:mt-10'>
              <div className='hidden transition ease-in-out duration-300 group-hover:flex cursor-pointer items-center justify-center absolute top-0 w-20 sm:w-40 h-20 sm:h-40 bg-gray-500 bg-opacity-50'>
                <FontAwesomeIcon icon={faCamera} className='text-white-accent-2 text-3xl absolute cursor-pointer '/>
              </div>
              <input 
                type="file"
                className='absolute top-0 w-20 sm:w-40 h-20 sm:h-40 opacity-0 cursor-pointer ' 
                onInput={(e) => handleImageChange(e) }
              />
              <img src={profile} alt="profile" className='cursor-pointer '/>
            </div>

            <p className='lato-regular text-xs sm:text-md mt-2'>Update Picture</p>
          </div>

          <div className='w-full mt-10 p-2 sm:ml-10'>
            <h1 className='lato-bold mb-3'>Profile Information</h1>
            <div className='w-[90%] flex gap-5 flex-col'>
              <Input
                type="email"
                label="Email"
                variant="bordered"
                placeholder="Enter your email"
                defaultValue={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-f"
              />
              <Input
                type="text"
                label="Name"
                variant="bordered"
                placeholder="Enter your name"
                defaultValue={name}
                onChange={(e) => setName(e.target.value)}
                className="w-f"
              />
              <Textarea
                variant='bordered'
                label="Description"
                placeholder="Enter your description"
                className="max-w-f"
                defaultValue={description}
                onChange={(e) => setDescription(e.target.value)}
                value={description}
              />
            </div>
          </div>
        </div>

        <div className='w-full mt-20 flex justify-end pr-10 mb-10'>
          <Button 
            color="danger"
            variant="bordered" 
            startContent={<FontAwesomeIcon icon={faSignOut} className='text-md'/>}
            onClick={() => {
              // handleLogout
              navigate('/login')
            }}
          >
            Log Out
          </Button>
        </div>
      </motion.div>
    </>
  )
}
export default Profile