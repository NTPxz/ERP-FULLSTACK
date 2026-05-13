import { useState, useCallback } from 'react'
import { apiGet, apiPost, apiPut, apiDelete } from '../api/client'
import type {
  Department, Position, Employee,
  CreateDepartmentInput, CreatePositionInput,
  CreateEmployeeInput, UpdateEmployeeInput,
} from '../types'

export function useDepartments() {
  const [departments, setDepartments] = useState<Department[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setDepartments(await apiGet<Department[]>('/departments')) }
    catch { setError('Failed to load departments') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateDepartmentInput) => {
    await apiPost<Department>('/departments', data); await load()
  }
  const update = async (id: number, data: CreateDepartmentInput) => {
    await apiPut<Department>(`/departments/${id}`, data); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/departments/${id}`); await load()
  }

  return { departments, loading, error, load, create, update, remove }
}

export function usePositions() {
  const [positions, setPositions] = useState<Position[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setPositions(await apiGet<Position[]>('/positions')) }
    catch { setError('Failed to load positions') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreatePositionInput) => {
    await apiPost<Position>('/positions', data); await load()
  }

  return { positions, loading, error, load, create }
}

export function useEmployees() {
  const [employees, setEmployees] = useState<Employee[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setEmployees(await apiGet<Employee[]>('/employees')) }
    catch { setError('Failed to load employees') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateEmployeeInput) => {
    await apiPost<Employee>('/employees', data); await load()
  }
  const update = async (id: number, data: UpdateEmployeeInput) => {
    await apiPut<Employee>(`/employees/${id}`, data); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/employees/${id}`); await load()
  }

  return { employees, loading, error, load, create, update, remove }
}
