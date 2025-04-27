import { useState, useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import instance from "../api/axios"

export default function PostForm({ mode = 'create', onPostCreated }) {
  const navigate = useNavigate()
  const { id } = useParams()

  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [error, setError] = useState(null)

  // 수정 모드일 경우 기존 데이터 로드
  useEffect(() => {
    if (mode === 'edit' && id) {
      instance.get(`/posts/${id}`)
        .then(res => {
          setTitle(res.data.title)
          setContent(res.data.content)
        })
        .catch(() => setError('글을 불러오지 못했습니다.'))
    }
  }, [mode, id])

  const handleSubmit = async () => {
    try {
      if (mode === 'edit') {
        await instance.put(`/posts/${id}`, { title, content })
        navigate(`/posts/${id}`)
      } else {
        await instance.post('/posts', { title, content })
        setTitle('')
        setContent('')
        setError(null)
        onPostCreated?.()
      }
    } catch (err) {
      console.error(err)
      setError(mode === 'edit' ? '글 수정 실패!' : '글 작성 실패!')
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
        <button type="button" onClick={handleSubmit}>
          {mode === 'edit' ? '수정하기' : '작성하기'}
        </button>
      </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  )
}