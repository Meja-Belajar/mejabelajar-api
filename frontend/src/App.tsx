import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { AnimatePresence } from 'framer-motion'

import Landing from '@src/pages/Landing/LandingPage'
import Login from '@src/pages/LoginPage'
import Register from '@src/pages/RegisterPage'
import ErrorPage from '@pages/ErrorPage'
import Profile from '@src/pages/ProfilePage';

import '@assets/global.css';
import { UserProvider } from '@contexts/UserContext'
import { AuthRedirector } from '@components/AuthRedirector';
import { QueryProvider } from '@contexts/SearchQueryContext'
import Search from '@pages/Search/SearchPage'
import SearchDefault from '@pages/Search/SearchDefault'
import SearchResult from '@pages/Search/SearchResult'

const App: React.FC = () => {

  return (
    <>
      <BrowserRouter>
        <UserProvider>

          <AnimatePresence>

            <Routes key="routes">
              {/* routes that don't require login */}
              <Route path="/" element={<Landing />} />
              <Route path='*' element={<ErrorPage />}/>
              <Route path='/announcement' />

              <Route
                path='/search'
                element={
                  <QueryProvider>
                    <Search />
                  </QueryProvider>
                }
              >
                <Route index element={<SearchDefault />} />
                <Route path=':query' element={<SearchResult />} />
              </Route>

              {/* routes that can't be accessed after logging in. */}
              <Route element={<AuthRedirector mustLogin={false}/>} >
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
              </Route>

              {/* routes that require login */}
              <Route element={<AuthRedirector mustLogin={true}/>} >
                <Route path="/profile" element={<Profile />} />
                <Route path='/history' />

                <Route path='/tutoring'/>
              </Route>

            </Routes>

          </AnimatePresence>
        </UserProvider>
      </BrowserRouter>
    </>
  )
}

export default App
