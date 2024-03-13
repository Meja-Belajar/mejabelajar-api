import { Button, Input, Navbar, NavbarBrand, NavbarContent, NavbarItem, NavbarMenu, NavbarMenuItem, NavbarMenuToggle } from '@nextui-org/react'
import logo from '@public/logo.svg'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faArrowRight, faClose, faSearch } from '@fortawesome/free-solid-svg-icons'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { useContext, useState } from 'react'
import {AnimatePresence, motion} from 'framer-motion'
import Logo from '@utils/Logo'
import { UserContext } from '@contexts/UserContext'

const navigationList = [
  "Announcement",
  "History",
]

const Navigation: React.FC = () => {

  const navigate = useNavigate();
  const [search, setSearch] = useState<string>('');
  const [isMenuOpen, setIsMenuOpen] = useState<boolean>(false);

  const [isHidden, setIsHidden] = useState(false);

  const handleIconClick = () => {
    setIsHidden(true);
  };

  const handleFormSubmit = (e: any) => {
    e.preventDefault();

    if(search) {
      navigate(`/search/${search}`);
    }
  }

  const { login } = useContext(UserContext);

  const location = useLocation();

  if(['/login', '/register'].includes(location.pathname)) return null;
  else return (
    <>
      <AnimatePresence>
        {
          !isHidden && (
            <motion.div 
              className='h-6 bg-white-accent-1 mb-3 flex items-center justify-center p-5 open-sans-600 transition ease-linear'
              exit={{ y: -100 }}
            >
              <h1 className='text-red-500 peer text-xs sm:text-md transition ease hover:opacity-50 cursor-pointer'>50% OFF BY USING THIS VOUCHER</h1>
              <FontAwesomeIcon 
                icon={faArrowRight} 
                className='pl-3 text-red-500 cursor-pointer transition ease-linear peer-hover:opacity:50' 
                fade 
                onClick={() => navigate('/promotion')}
              /> 
              <FontAwesomeIcon icon={faClose} className='cursor-pointer absolute right-5 transition ease-linear hover:opacity-50' onClick={handleIconClick} />
            </motion.div>
          )
        }
      </AnimatePresence>

      <Navbar
        onMenuOpenChange={setIsMenuOpen}
        shouldHideOnScroll
        className='w-full mt-1 p-1'
        maxWidth='xl'
      >
        <div className='flex gap-5 items-center justify-center flex-row'>
          <NavbarMenuToggle
            aria-label={isMenuOpen ? "Close" : "Open"}
            className='sm:hidden'
          />
          <Logo />

        </div>
        
        <form 
          className='flex gap-10 flex-row justify-end w-3/4 lg:w-1/2'
          onSubmit={(e) => handleFormSubmit(e) }
        >
          <Input
            type='text'
            placeholder='search courses'
            variant='bordered'
            className='hidden md:flex p-3 w-full'
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            startContent={
              <FontAwesomeIcon icon={faSearch} className='text-blue-accent-300' />
            } 
          />
          <div className='gap-3 flex items-center'>
            {
              login?.status !== 200 ? (
                <>
                  <Button variant='bordered' className='border-blue-accent-300 text-blue-accent-300 open-sans-600 text-xs w-16 h-8 sm:w-24 sm:h-10' onClick={() => navigate('/login')}>Login</Button>
                  <Button variant='shadow' className='hidden sm:flex bg-blue-accent-300 text-white open-sans-600  text-xs w-16 h-8 sm:w-24 sm:h-10' onClick={() => navigate('/register')}>Register</Button>
                </>
              ) : (
                <>
                  <Link to='/profile' className='relative w-10 overflow-hidden rounded-full mr-4'>
                    <img src={login?.user?.account_detail?.profile_picture} alt="" className='w-full h-full' />
                  </Link>
                </>
              )
            }
          </div>
        </form>

        <NavbarMenu className='pt-20'>
          {navigationList.map((item, index) => (
            <NavbarMenuItem key={`${item}-${index}`}>
              <Link
                color={
                  index === 2 ? "primary" : index === navigationList.length - 1 ? "danger" : "foreground"
                }
                className="w-full text-xl"
                to={`/${item.toLowerCase()}`}
              >
                {item}
              </Link>
            </NavbarMenuItem>
          ))}
        </NavbarMenu>
      </Navbar>
      
      <nav className='w-full flex justify-center'>
        <Input
          type='text'
          placeholder='search courses'
          variant='bordered'
          className='flex p-3 w-full md:hidden'
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          startContent={
            <FontAwesomeIcon icon={faSearch} className='text-blue-accent-300' />
          } 
        />
      </nav>
    </>
  )
}

export default Navigation