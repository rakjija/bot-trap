import { useState } from 'react'
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom'

import LoginForm from './components/LoginForm'
import PostList from './components/PostList'
import PostForm from './components/PostForm'
import PostDetail from './components/PostDetail'
import SignupForm from './components/SignupForm'

export default function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('access_token'))
  const navigate = useNavigate()

  const handleLogout = () => {
    localStorage.removeItem('access_token')
    localStorage.removeItem('user_id')
    setIsLoggedIn(false)
    navigate('/')
  }

  return (
    <div>
      {!isLoggedIn && <LoginForm onLogin={() => setIsLoggedIn(true)} />}

      <Routes>
        <Route
          path="/"
          element={
            <>
              {isLoggedIn && (
                <div className="header-buttons">
                  <button type="button" onClick={() => navigate('/write')}>글쓰기</button>
                  <button type="button" onClick={handleLogout}>로그아웃</button>
                </div>
              )}
              <PostList />
            </>
          }
        />

        <Route
          path="/write"
          element={
            isLoggedIn ? (
              <PostForm onPostCreated={() => navigate('/')} />
            ) : (
              <Navigate to="/" />
            )
          }
        />

        <Route
          path="/posts/:id"
          element={<PostDetail />}
        />

        <Route
          path="/posts/:id/edit"
          element={isLoggedIn ? <PostForm mode="edit" /> : <Navigate to="/" />}
        />

        <Route path="/signup" element={<SignupForm />} />
      </Routes>
    </div>
  )
}