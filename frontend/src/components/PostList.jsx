// src/components/PostList.jsx
import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import instance from "../api/axios"
import '../styles/board.css'

function formatDate(dateStr) {
  const date = new Date(dateStr)
  return date.toISOString().slice(0, 10) // YYYY-MM-DD
}

export default function PostList() {
  const [posts, setPosts] = useState([])
  const navigate = useNavigate()

  useEffect(() => {
    instance.get('/posts')
      .then(res => setPosts(res.data.reverse())) // 최신순으로 정렬
      .catch(err => console.error('Failed to fetch posts:', err))
  }, [])

  console.log(posts)

  return (
    <div className="board-container">
      <h2 className="board-title">GoBoard</h2>
      <table className="board-table">
        <thead>
          <tr>
            <th>No</th>
            <th>Title</th>
            <th>User</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>
          {posts.map((post, idx) => (
            <tr key={post.id}>
              <td>{posts.length - idx}</td>
              <td>
              <button type="button" onClick={() => navigate(`/posts/${post.id}`)}
                style={{
                  all: 'unset', // 스타일 초기화 (td랑 어울리게)
                  cursor: 'pointer',
                  color: '#1e88e5'
                }}
              >
                {post.title}
              </button>
            </td>
              <td>{post.username}</td>
              <td>{formatDate(post.created_at)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}