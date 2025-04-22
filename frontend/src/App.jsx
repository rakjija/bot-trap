import { useState } from 'react'
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom'

import LoginForm from './components/LoginForm'
import PostList from './components/PostList'
import PostForm from './components/PostForm'
import PostDetail from './components/PostDetail' // ✅ 누락됐던 부분

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
      {/* 로그인하지 않았을 경우만 로그인폼 보여주기 */}
      {!isLoggedIn && <LoginForm onLogin={() => setIsLoggedIn(true)} />}

      <Routes>
        <Route
          path="/"
          element={
            <>
              {/* 로그인된 경우만 버튼 표시 */}
              {isLoggedIn && (
                <div className="header-buttons">
                  <button type="button" onClick={() => navigate('/write')}>글쓰기</button>
                  <button type="button" onClick={handleLogout}>로그아웃</button>
                </div>
              )}
              <PostList key={refreshKey} />
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
      </Routes>
    </div>
  )
}