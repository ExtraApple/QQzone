import request from '@/utils/request'

export interface UserInfo {
  id: number
  username: string
  role: string
}

export interface LoginParams {
  username: string
  password: string
}

export interface RegisterParams {
  username: string
  password: string
}

export function login(params: LoginParams) {
  return request.post<{ token: string }>('/user/login', params)
}

export function register(params: RegisterParams) {
  return request.post<{ user: UserInfo }>('/user/register', params)
}

export function addFriend(id: number) {
  return request.post(`/user/friends/add/${id}`)
}

export function deleteFriend(id: number) {
  return request.delete(`/user/friends/delete/${id}`)
}

export function listFriends() {
  return request.get<{ friends: UserInfo[] }>('/user/friends')
}

export function adminCheck() {
  return request.get<{ msg: string; user: any }>('/user/admin')
}

export function logout() {
  return request.delete('/user/logout')
}
