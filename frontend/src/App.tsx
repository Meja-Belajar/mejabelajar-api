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

            <Routes >
              <Route path="/" element={<Landing />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
            </Routes>

            <Footer/>
          </AnimatePresence>
        
        </UserProvider>
      </BrowserRouter>
    </>
  )
}

export default App
