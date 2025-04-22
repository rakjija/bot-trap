import React, { useState } from 'react'
import LoginForm from './components/LoginForm'
import PostList from './components/PostList'
import PostForm from './components/PostForm'

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('access_token'))
  const [refreshKey, setRefreshKey] = useState(0)

  const reloadPosts = () => setRefreshKey(prev => prev + 1)

  const handleLogout = () => {
    localStorage.removeItem('access_token')
    setIsLoggedIn(false)
  }

  return (
    <div>
      {!isLoggedIn && <LoginForm onLogin={() => setIsLoggedIn(true)} />}
      {isLoggedIn && (
        <>
          <div className="logout-container">
            <button onClick={handleLogout}>로그아웃</button>
          </div>
          <PostForm onPostCreated={reloadPosts} />
        </>
      )}
      <PostList key={refreshKey} />
    </div>
  )
}

export default App