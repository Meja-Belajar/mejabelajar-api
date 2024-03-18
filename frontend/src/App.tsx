import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { AnimatePresence } from 'framer-motion'

import Landing from '@src/pages/Landing/Landing'
import Login from '@pages/Login'
import Register from '@pages/Register'
import ErrorPage from '@pages/ErrorPage'
import Profile from '@pages/Profile';

import '@assets/global.css';
import { UserProvider } from '@contexts/UserContext'
import Auth from '@utils/Auth';

const App: React.FC = () => {

  return (
    <>
      <BrowserRouter>
        <UserProvider>

          <AnimatePresence>

            <Routes key="routes">
              <Route path="/" element={<Landing />} />
              <Route path='*' element={<ErrorPage />}/>
              
              <Route element={<Auth />} >
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
              </Route>

              
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

          </AnimatePresence>
        
        </UserProvider>
      </BrowserRouter>
    </>
  )
}

export default App
