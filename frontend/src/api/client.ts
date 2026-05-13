import axios, { type AxiosInstance, type InternalAxiosRequestConfig } from 'axios'

const http: AxiosInstance = axios.create({ baseURL: '/api' })

http.interceptors.request.use((cfg: InternalAxiosRequestConfig) => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

http.interceptors.response.use(
  res => res,
  err => {
    const url: string = err.config?.url ?? ''
    if (err.response?.status === 401 && !url.includes('/auth/')) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/login'
    }
    return Promise.reject(err)
  }
)

interface ApiResponse<T> { data: T }

export async function apiGet<T>(url: string): Promise<T> {
  const res = await http.get<ApiResponse<T>>(url)
  return res.data.data
}

export async function apiPost<T>(url: string, body: unknown): Promise<T> {
  const res = await http.post<ApiResponse<T>>(url, body)
  return res.data.data
}

export async function apiPut<T>(url: string, body: unknown): Promise<T> {
  const res = await http.put<ApiResponse<T>>(url, body)
  return res.data.data
}

export async function apiPatch<T>(url: string, body: unknown): Promise<T> {
  const res = await http.patch<ApiResponse<T>>(url, body)
  return res.data.data
}

export async function apiDelete(url: string): Promise<void> {
  await http.delete(url)
}

export default http
