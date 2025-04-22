import { useState } from 'react'
import axios from '../api/axios'

export default function LoginForm({ onLogin }) {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState(null)

  const handleLogin = async () => {
    try {
      const res = await axios.post('/users/login', { email, password })
      localStorage.setItem('access_token', res.data.access_token)
      setError(null)
      onLogin?.()
    } catch (err) {
      console.error(err)
      setError('로그인 실패: 이메일이나 비밀번호를 확인해주세요')
    }
  }

  return (
    <div className="login-form">
      <input
        type="email"
        placeholder="이메일"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="password"
        placeholder="비밀번호"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={handleLogin}>로그인</button>
    </div>
  )
}