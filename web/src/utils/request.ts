import axios from 'axios'
import { ElMessage } from 'element-plus'

const request = axios.create({
  baseURL: '/api',
  timeout: 15000,
})

request.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

request.interceptors.response.use(
  (response) => response,
  (error) => {
    const status = error.response?.status
    const msg = error.response?.data?.error

    if (status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      window.location.href = '/login'
      return Promise.reject(error)
    }

    if (!error.response) {
      console.error('网络错误:', error.message, error.code)
      ElMessage.error('无法连接到服务器，请确认后端已启动')
    } else if (msg) {
      ElMessage.error(msg)
    } else {
      ElMessage.error(`请求失败 (${status})`)
    }

    return Promise.reject(error)
  }
)

export default request
