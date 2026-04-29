import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { UserInfo } from '@/api/user'
import * as userApi from '@/api/user'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref<UserInfo | null>(null)

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => userInfo.value?.role === 'admin')

  const storedUser = localStorage.getItem('userInfo')
  if (storedUser) {
    try {
      userInfo.value = JSON.parse(storedUser)
    } catch {}
  }

  async function login(params: userApi.LoginParams) {
    const res = await userApi.login(params)
    token.value = res.data.token
    localStorage.setItem('token', res.data.token)
    await decodeAndSave()
  }

  async function register(params: userApi.RegisterParams) {
    const res = await userApi.register(params)
    return res.data
  }

  async function decodeAndSave() {
    try {
      const payload = JSON.parse(atob(token.value.split('.')[1]))
      const info: UserInfo = {
        id: payload.user_id,
        username: payload.username,
        role: payload.role,
      }
      userInfo.value = info
      localStorage.setItem('userInfo', JSON.stringify(info))
    } catch {
      logout()
    }
  }

  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
  }

  return { token, userInfo, isLoggedIn, isAdmin, login, register, decodeAndSave, logout }
})
