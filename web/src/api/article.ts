import request from '@/utils/request'

export interface Media {
  id: number
  type: string
  url: string
  objectKey: string
  size: number
  duration: number
}

export interface Article {
  id: number
  title: string
  content: string
  user_id: number
  media: Media[]
  created_at: string
  updated_at: string
}

export function articleList() {
  return request.get<{ articles: Article[] }>('/articles')
}

export function articleDetail(id: number) {
  return request.get<{ article: Article }>(`/articles/${id}`)
}

export function createArticle(formData: FormData) {
  return request.post<{ msg: string; article: Article }>('/articles/create', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}

export function deleteArticle(id: number) {
  return request.delete<{ msg: string }>(`/articles/${id}`)
}

export function getComments(articleId: number) {
  return request.get<any[]>(`/articles/${articleId}/comments`)
}

export function createComment(articleId: number, data: { content: string; parent_id?: number }) {
  return request.post(`/articles/${articleId}/comments`, data)
}

export function deleteComment(commentId: number) {
  return request.delete(`/articles/comments/${commentId}`)
}
