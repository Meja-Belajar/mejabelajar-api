import { Button, Input, Navbar, NavbarBrand, NavbarContent, NavbarItem, NavbarMenu, NavbarMenuItem, NavbarMenuToggle } from '@nextui-org/react'
import logo from '../../public/vite.svg'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faArrowRight, faClose, faLeftLong, faRightLong, faSearch } from '@fortawesome/free-solid-svg-icons'
import { Link, useNavigate } from 'react-router-dom'
import { useState } from 'react'
import {AnimatePresence, motion} from 'framer-motion'
const navigationList = [
  "Home",
  "Announcement",
  "Ask Mentor",
  "Find Mentor",
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
  return (
    <>
      <AnimatePresence>
        {
          !isHidden && (
            <motion.div 
              className='h-6 bg-white-accent-1 mb-3 flex items-center justify-center p-5 lato-bold transition ease-linear'
              exit={{ y: -100 }}
            >
              <h1 className='text-red-500 peer transition ease hover:opacity-50 cursor-pointer'>50% OFF BY USING THIS VOUCHER</h1>
              <FontAwesomeIcon 
                icon={faArrowRight} 
                className='pl-3 text-red-500 cursor-pointer transition ease-linear peer-hover:opacity:50' 
                fade 
                onClick={() => navigate('/promotion')}
              /> 
              <FontAwesomeIcon icon={faClose} className='cursor-pointer absolute right-5 transition ease-linear hover:opacity-50' onClick={handleIconClick} />
            </motion.div>
          )}
      </AnimatePresence>

      <Navbar
        onMenuOpenChange={setIsMenuOpen}
        shouldHideOnScroll
        className='w-full mt-1'
        maxWidth='xl'
      >
        <div className='flex gap-5 items-center justify-center flex-row'>
          <NavbarMenuToggle
            aria-label={isMenuOpen ? "Close" : "Open"}
            className='sm:hidden'
          />
          <div className='flex items-center justify-center flex-row' >
            <img src={logo} alt="logo" className='w-6 mb-1' />
            <h1 className='special-font text-blue-accent-400'>MejaBelajar</h1>
          </div>

        </div>
        
        <form 
          className='flex gap-10 flex-row justify-end w-3/4 lg:w-1/2'
          onSubmit={(e) => handleFormSubmit(e) }
        >
          <Input
            type='text'
            placeholder='search courses'
            variant='bordered'
            className='hidden md:flex lato-regular p-3 w-full'
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            startContent={
              <FontAwesomeIcon icon={faSearch} className='text-blue-accent-300' />
            } 
          />
          <div className='gap-3 flex items-center'>
            <Button variant='bordered' className='border-blue-accent-300 text-blue-accent-300 lato-bold text-xs w-16 h-8 sm:w-24 sm:h-10' onClick={() => navigate('/login')}>Login</Button>
            <Button variant='shadow' className='hidden sm:flex bg-blue-accent-300 text-white lato-bold  text-xs w-16 h-8 sm:w-24 sm:h-10' onClick={() => navigate('/register')}>Register</Button>
          </div>
        </form>

        <NavbarMenu className='pt-20'>
          {navigationList.map((item, index) => (
            <NavbarMenuItem key={`${item}-${index}`}>
              <Link
                color={
                  index === 2 ? "primary" : index === navigationList.length - 1 ? "danger" : "foreground"
                }
                className="w-full text-xl lato-reguler"
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
          className='lato-regular flex p-3 w-full md:hidden'
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