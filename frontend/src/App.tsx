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
const App: React.FC = () => {

  return (
    <>
      <BrowserRouter>
        <UserProvider>

          <AnimatePresence>

            <Routes key="routes">
              <Route path="/" element={<Landing />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
            </Routes>

            <Footer key="footer"/>
          </AnimatePresence>
        
        </UserProvider>
      </BrowserRouter>
    </>
  )
}

export default App
