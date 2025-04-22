// src/components/SignupForm.jsx
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from '../api/axios'

export default function SignupForm() {
  const [email, setEmail] = useState('')
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState(null)
  const navigate = useNavigate()

  const handleSignup = async () => {
    try {
      await axios.post('/users/signup', { email, username, password })
      setError(null)
      alert('회원가입 성공! 이제 로그인하세요.')
      navigate('/') // 로그인 화면으로
    } catch (err) {
      console.error(err)
      setError(err.response?.data?.error || '회원가입 실패')
    }
  }

  return (
    <div className="post-form">
      <h2>회원가입</h2>
      <input
        type="text"
        placeholder="이메일"
        value={email}
        onChange={e => setEmail(e.target.value)}
      />
      <input
        type="text"
        placeholder="이름"
        value={username}
        onChange={e => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="비밀번호 (8자 이상)"
        value={password}
        onChange={e => setPassword(e.target.value)}
      />
      <div className="form-buttons">
        <button type="button" onClick={() => navigate('/')}>취소</button>
        <button type="button" onClick={handleSignup}>가입하기</button>
      </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  )
}