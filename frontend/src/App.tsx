import { useEffect } from 'react'
import './assets/global.css'
import { Button } from '@nextui-org/react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Landing from './pages/Landing'
import { AnimatePresence } from 'framer-motion'
import { UserProvider } from './context/UserContext'
import Login from './pages/Login'
import Register from './pages/Register'
import Footer from './components/Footer'
import Navigation from './components/Navigation'
import ErrorPage from './pages/ErrorPage'
import Profile from './pages/Profile'

const App: React.FC = () => {

  return (
    <>
      <BrowserRouter>
        <UserProvider>

          <AnimatePresence>
            <Navigation key="navigation"/>

            <Routes key="routes">
              <Route path="/" element={<Landing />} />
              <Route path='*' element={<ErrorPage />}/>
              
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />

              <Route path="/profile" element={<Profile />} />
                
              <Route
                path='/search'
                // element={
                  // <QueryProvider>
                  //   <Search/>
                  // </QueryProvider>
                // }
              >
                <Route index element={<></>} />
                <Route path=':query' element={<></>} />
              </Route>

              <Route path='/tutoring'/>

              <Route
                path='/mentoring'
              >
                <Route index/>
                <Route path=':filter'/>
              </Route>

              <Route path='/announcement' />
              <Route path='/history' />

            </Routes>

            <Footer key="footer"/>
          </AnimatePresence>
        
        </UserProvider>
      </BrowserRouter>
    </>
  )
}

export default App
