import { useState } from 'react'
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom'
import LoginForm from './components/LoginForm'
import PostList from './components/PostList'
import PostForm from './components/PostForm'

export default function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('access_token'))
  const [refreshKey, setRefreshKey] = useState(0)
  const navigate = useNavigate()

  const reloadPosts = () => setRefreshKey(prev => prev + 1)

  const handleLogout = () => {
    localStorage.removeItem('access_token')
    setIsLoggedIn(false)
    navigate('/')
  }

  return (
    <div>
      {!isLoggedIn && <LoginForm onLogin={() => setIsLoggedIn(true)} />}

      <Routes>
        <Route path="/" element={
          <>
            {isLoggedIn && (
              <div className="header-buttons">
                <button type="button" onClick={() => navigate('/write')}>글쓰기</button>
                <button type="button" onClick={handleLogout}>로그아웃</button>
              </div>
            )}
            <PostList key={refreshKey} />
          </>
        } />

        <Route path="/write" element={
          isLoggedIn ? (
            <PostForm onPostCreated={() => navigate('/')} />
          ) : (
            <Navigate to="/" />
          )
        } />
      </Routes>
    </div>
  )
}