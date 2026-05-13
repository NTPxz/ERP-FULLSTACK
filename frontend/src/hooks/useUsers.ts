import { useState, useCallback } from 'react'
import { apiGet, apiPost, apiPatch, apiDelete } from '../api/client'
import type { User, AdminCreateUserInput, AdminUpdateUserInput } from '../types'

export function useUsers() {
  const [users, setUsers] = useState<User[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setUsers(await apiGet<User[]>('/users')) }
    catch { setError('Failed to load users') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: AdminCreateUserInput) => {
    await apiPost<User>('/users', data); await load()
  }
  const updateRole = async (id: number, data: AdminUpdateUserInput) => {
    await apiPatch<User>(`/users/${id}/role`, data); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/users/${id}`); await load()
  }

  return { users, loading, error, load, create, updateRole, remove }
}
