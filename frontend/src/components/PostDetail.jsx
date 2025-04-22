import { useEffect, useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import axios from '../api/axios'

export default function PostDetail() {
  const { id } = useParams()
  const userId = localStorage.getItem('user_id')
  const navigate = useNavigate()
  const [post, setPost] = useState(null)
  const [error, setError] = useState(null)

  useEffect(() => {
    axios.get(`/posts/${id}`)
      .then(res => setPost(res.data))
      .catch(() => setError('글을 불러오지 못했습니다.'))
  }, [id])

  if (error) return <p style={{ color: 'red' }}>{error}</p>
  if (!post) return <p>불러오는 중...</p>

  return (
    <div style={{ width: '90%', margin: '40px auto', fontFamily: 'sans-serif' }}>
      <p style={{ fontSize: '0.9rem', color: '#777', marginBottom: '8px' }}>
        작성일: {new Date(post.created_at).toLocaleString()}
      </p>

      <h2 style={{ fontSize: '1.8rem', fontWeight: 'bold', marginBottom: '16px' }}>
        {post.title}
      </h2>

      <div
        style={{
          whiteSpace: 'pre-line',
          lineHeight: 1.6,
          fontSize: '1rem',
          marginBottom: '32px'
        }}
      >
        {post.content}
      </div>

      <div className="post-detail-buttons">
        {post.user_id?.toString() === userId && (
          <button type="button" onClick={() => navigate(`/posts/${id}/edit`)}>
            ✏️ 수정하기
          </button>
        )}

        <button type="button" onClick={() => navigate('/')}>
          ← 돌아가기
        </button>
      </div>
    </div>
  )
}