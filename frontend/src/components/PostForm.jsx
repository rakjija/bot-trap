import { useState } from 'react'
import axios from '../api/axios'
import { useNavigate } from 'react-router-dom'

export default function PostForm({ onPostCreated }) {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [error, setError] = useState(null)
  const navigate = useNavigate()

  const handleSubmit = async () => {
    try {
      await axios.post('/posts', { title, content })
      setTitle('')
      setContent('')
      setError(null)
      onPostCreated?.()
    } catch (err) {
      console.error(err)
      setError('글 작성 실패! 로그인 상태를 확인해주세요.')
    }
  }

  return (
    <div className="post-form">
      <input
        type="text"
        placeholder="제목"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />
      <textarea
        placeholder="내용"
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />
      <div className="form-buttons">
        <button type="button" onClick={() => navigate('/')}>취소</button>
        <button type="button" onClick={handleSubmit}>작성하기</button>
      </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  )
}