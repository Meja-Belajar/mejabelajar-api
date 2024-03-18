import {motion} from 'framer-motion';
import { animate, exit, initial } from '@assets/PageTransition';
import Footer from '@components/Footer';
import Logo from '@src/components/Logo';
import { Button, Input, Navbar, NavbarMenu, NavbarMenuItem, NavbarMenuToggle, Skeleton } from '@nextui-org/react';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { faArrowUpRightFromSquare, faSearch } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import image from '@assets/images/teacher.png';
import star from '@assets/images/icon/star.png';

const Programs = [
  {
    title: "Feature 1",
    desc: "description of the text at here will be added later",
    img: "star",
  },
  {
    title: "Feature 1",
    desc: "description of the text at here will be added later",
    img: "star",
  },
  {
    title: "Feature 1",
    desc: "description of the text at here will be added later",
    img: "star",
  },
]

const NavigationLists = [
  {
    name: "mentor",
    link: "/mentor"
  },
  {
    name: "forum",
    link: "/forum"
  },
  {
    name: "about",
    link: "/about"
  },
  {
    name: "login",
    link: "/login"
  }
];

const UserReviews = [
  {
    name: "Matt James",
    role: "Mentor",
    review: "with lorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsum",
    image: "",
  },
  {
    name: "Matt James",
    role: "Mentor",
    review: "with lorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsum",
    image: "",
  },
  {
    name: "Matt James",
    role: "Mentor",
    review: "with lorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsumlorem ipsum",
    image: "",
  },
];

const PopularCourses = [
  {
    title: "Database Design",
    img: "",
    alt: ""
  },
  {
    title: "Database Design",
    img: "",
    alt: ""
  },
  {
    title: "Database Design",
    img: "",
    alt: ""
  },
];

const BestMentors = [
  {
    name: "Matt James",
    image: "",
    degree: "Computer Science",
  },
  {
    name: "Matt James",
    image: "",
    degree: "Computer Science",
  },
  {
    name: "Matt James",
    image: "",
    degree: "Computer Science",
  },
  {
    name: "Matt James",
    image: "",
    degree: "Computer Science",
  },
  {
    name: "Matt James",
    image: "",
    degree: "Computer Science",
  },
]

// Landing page for guest
const GuestLanding = () => {
  const [isMenuOpen, setIsMenuOpen] = useState<boolean>(false);
  const [search, setSearch] = useState<string>("");
  const navigate = useNavigate();

  const handleSearch: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
    // do query here

    navigate(`/search/${encodeURIComponent(search)}`);
  }

  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='bg-white-accent-1 pb-32'
      >  
        <main className='relative overflow-hidden'> 
          <div className='absolute top-0 -z-6 hidden lg:block'>
            <img src={image} alt="images" className='w-[100vw]' />
          </div>

          <Navbar 
            className='w-full border-none bg-transparent blur-none flex flex-row justify-around items-center p-3 px-12'
            onMenuOpenChange={setIsMenuOpen}
            shouldHideOnScroll
            maxWidth='xl'
            isBlurred={false}
          >
            <div>
              <Logo />
            </div>
            <NavbarMenuToggle
              aria-label={isMenuOpen ? "Close" : "Open"}
              className='md:hidden'
            />
            <div className='hidden md:flex flex-row items-center px-3 py-2 rounded-full bg-white text-sm shadow-lg'>
              <ul className='flex flex-row items-center gap-4 mr-6 open-sans-600 ml-2'>
                {
                  NavigationLists.map((item, index) => (
                    <li key={index} className='hover:opacity-50 transition duration-300 ease-linear'>
                      <Link to={item.link}>{item.name}</Link>
                    </li>
                  ))
                }
              </ul>
              <Button className='bg-blue-accent-300 text-white rounded-full px-5 height-2 open-sans-600 text-sm' size='sm'>
                <Link to='/register'>sign up</Link>
              </Button>
            </div>

            <NavbarMenu className='pt-20 px-10 mt-5'>
              {
                NavigationLists.map((item, index) => (
                  <NavbarMenuItem key={`${item}-${index}`}>
                    <Link
                      color={
                        index === 2 ? "primary" : index === NavigationLists.length - 1 ? "danger" : "foreground"
                      }
                      className="w-full text-xl "
                      to={item.link}
                    >
                      {item.name}
                    </Link>
                  </NavbarMenuItem>
                ))
              }
            </NavbarMenu>
          </Navbar>

          <form 
            className='w-full flex flex-col items-center justify-center py-20'
            onSubmit={ handleSearch }
          >
            <div className='text-center z-10'>
              <h1 className='p-5 sm:p-0 text-6xl text-blue-accent-400 mb-0 sm:mb-5'>Discover Your Perfect <br/> Mentor Today!</h1>
              <h3 className='p-5 sm:p-0 text-xl text-blue-accent-400'>Meja Belajar now is your gateway to <br/>connecting with experienced mentors who <br/>are eager to share their knowledge and <br/>expertise with you.</h3>
            </div>
            
            <Input
              color='default'
              type='text'
              placeholder='search your course'
              variant='bordered'
              className='relative w-3/4 sm:w-1/2 md:w-1/3 mt-10 focus:border-white'
              classNames={{
                label: 'text-blue-accent-400',
                input: 'bg-white focus:border-white-accent-1 border-white-accent-1',
                inputWrapper: [
                  'bg-white border border-none',
                ],
              }}
              radius='full'
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              startContent={
                <FontAwesomeIcon icon={faSearch} className='text-blue-accent-300 mr-3' />
              } 
            />
          </form>

        </main>
        
        <section className='mt-20 w-full flex items-center justify-center'>
          <div className='w-full m-5 sm:m-10 p-8 flex flex-col sm:flex-row justify-between shadow-md drop-shadow-lg sm:h-96 rounded-3xl sm:gap-10 bg-gradient-to-r from-purple-500 via-purple-400 to-blue-accent-300'>
            <div className='p-3 w-full sm:h-auto mb-10 sm:mb-0 sm:w-1/3 rounded-2xl bg-white-accent-1 drop-shadow-md shadow-md'>
              <Skeleton className="w-full flex flex-col rounded-lg">
                <div className="h-[5vh] rounded-lg bg-default-200" />
              </Skeleton>
              <Skeleton className="w-3/4 flex flex-col mt-3 rounded-lg">
                <div className="h-[5vh] rounded-lg bg-default-200" />
              </Skeleton>
              <Skeleton className="w-3/4 flex flex-col mt-3 rounded-lg">
                <div className="h-[5vh] rounded-lg bg-default-200" />
              </Skeleton>
              <Skeleton className="w-1/2 flex flex-col mt-3 rounded-lg">
                <div className="h-[5vh] rounded-lg bg-default-200" />
              </Skeleton>
            </div>
            <div className='p-1 sm:p-3 w-full text-white-accent-1 sm:w-2/3 flex justify-center flex-col sm:pr-10'>
              <h1 className='text-3xl p-3'>Enhance Academic Performance</h1>
              <h1 className='text-xl sm:text-justify p-2 pl-3'>Meja Belajar assists university students in enhancing their academic performance with the aid of our mentors. Whenever you're ready, our mentors are always prepared to provide assistance.</h1>
              <div className='mt-10 flex items-center'>
                <Link to='/announcement' className='peer p-3 underline-offset-4 underline decoration-transparent transition ease-linear duration-300 hover:decoration-white-accent-1'>
                  learn more 
                </Link>
                <FontAwesomeIcon icon={faArrowUpRightFromSquare} fade className='peer-hover:scale-110'/>
              </div>
            </div>
          </div>
        </section>

        <section className='pt-20 pb-28 sm:pt-36 sm:pb-48 bg-gradient-to-b from-white-accent-1 via-white to-white-accent-1 flex flex-col lg:flex-row items-center justify-between px-5 sm:px-20 lg:px-36 gap-5'>
          <div className='px-5 py-10'>
            <h1 className='text-4xl'>Our Program</h1>
            <h2 className='text-sm mt-2'>Whether you're a student, professional, or entrepreneur</h2>
          </div>
          {
            Programs.map((item, index) => (
              <div key={index} className='px-5 py-10'>
                <img src={star} className='mb-4'/>
                <h1 className='text-xl open-sans-600'>{item.title}</h1>
                <p className='text-sm mt-2'>{item.desc}</p>
              </div>
            ))
          }
        </section>

        <section className='w-full items-center flex flex-col justify-center p-10'>
          <div className='text-center text-4xl text-blue-accent-400 mb-20'>
            <h1>View some reviews <br/>from <a className='special-font'>Meja Belajar'</a> users</h1>
          </div>
          <div className='flex flex-col md:flex-row w-full items-center justify-between gap-14 sm:gap-8 lg:px-14'>
            {
              UserReviews.map((item, index) => (
                <div key={index} className=' w-3/4 lg:max-w-[25%] bg-blue-accent-300 shadow-xl drop-shadow-md shadow-blue-accent-300 p-5 rounded-3xl'>
                  <div className='flex flex-row justify-around items-center gap-5 p-2'>
                    <div className='overflow-hidden w-14 h-14 bg-gray-200 rounded-full'>
                      <img src={item.image} alt={item.name} />
                    </div>
                    <div className='text-white'>
                      <h1 className='text-xl open-sans-600'>{item.name}</h1>
                      <h2 className='text-sm'>{item.role}</h2>
                    </div>
                  </div>

                  <div className='text-white line-clamp-3 pt-3 px-3 overflow-hidden'>
                    <p>
                      {item.review}
                    </p>
                  </div>

                  <div className='p-3 mt-2 flex flex-row items-center justify-start'>
                    <img src={star} />
                    <img src={star} />
                    <img src={star} />
                  </div>
                </div>
              ))
            }

          </div>
        </section>
        
        <section className='mt-20 md:mt-32 p-10 w-full bg-gradient-to-r from-purple-500 to-blue-accent-300'>
          <div className='text-white p-3 border border-black'>
            <h1 className='open-sans-600 text-3xl'>Our Best Mentor</h1>
            <p>Congratulations to our top mentor for the odd semester, whose exceptional performance and dedication have set a new standard for excellence in mentorship.</p>
          </div>

          <div className='p-4 mt-10 flex flex-col md:flex-row gap-3 justify-center'>
            {
              BestMentors.map((item, index) => (
                <div key={index} className='border border-black p-4'>
                  <div className='w-full border border-black p-3 h-28 rounded-2xl mb-3'>
                    <img src={item.image} />
                  </div>
                  <div className='text-white text-center'>
                    <h1 className='text-lg open-sans-600'>{item.name}</h1>
                    <h2 className='opacity-80 text-sm'>{item.degree}</h2>
                  </div>
                </div>
              ))
            }

          </div>
        </section>

      </motion.div>  
      
      <Footer />  
    </>
  )
}

export default GuestLanding;