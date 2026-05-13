import { useState, useCallback } from 'react'
import { apiGet, apiPost, apiPut, apiPatch, apiDelete } from '../api/client'
import type {
  Customer, SalesOrder, SalesOrderStatus,
  CreateCustomerInput, UpdateCustomerInput, CreateSalesOrderInput,
} from '../types'

export function useCustomers() {
  const [customers, setCustomers] = useState<Customer[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setCustomers(await apiGet<Customer[]>('/customers')) }
    catch { setError('Failed to load customers') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateCustomerInput) => {
    await apiPost<Customer>('/customers', data); await load()
  }
  const update = async (id: number, data: UpdateCustomerInput) => {
    await apiPut<Customer>(`/customers/${id}`, data); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/customers/${id}`); await load()
  }

  return { customers, loading, error, load, create, update, remove }
}

export function useSalesOrders() {
  const [orders, setOrders] = useState<SalesOrder[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setOrders(await apiGet<SalesOrder[]>('/sales-orders')) }
    catch { setError('Failed to load sales orders') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateSalesOrderInput) => {
    await apiPost<SalesOrder>('/sales-orders', data); await load()
  }
  const updateStatus = async (id: number, status: SalesOrderStatus) => {
    await apiPatch<SalesOrder>(`/sales-orders/${id}/status`, { status }); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/sales-orders/${id}`); await load()
  }

  return { orders, loading, error, load, create, updateStatus, remove }
}
