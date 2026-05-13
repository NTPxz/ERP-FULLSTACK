import { useState, useCallback } from 'react'
import { apiGet, apiPost, apiPut, apiPatch, apiDelete } from '../api/client'
import type {
  Supplier, PurchaseOrder, PurchaseOrderStatus,
  CreateSupplierInput, UpdateSupplierInput, CreatePurchaseOrderInput,
} from '../types'

export function useSuppliers() {
  const [suppliers, setSuppliers] = useState<Supplier[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setSuppliers(await apiGet<Supplier[]>('/suppliers')) }
    catch { setError('Failed to load suppliers') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateSupplierInput) => {
    await apiPost<Supplier>('/suppliers', data); await load()
  }
  const update = async (id: number, data: UpdateSupplierInput) => {
    await apiPut<Supplier>(`/suppliers/${id}`, data); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/suppliers/${id}`); await load()
  }

  return { suppliers, loading, error, load, create, update, remove }
}

export function usePurchaseOrders() {
  const [orders, setOrders] = useState<PurchaseOrder[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setOrders(await apiGet<PurchaseOrder[]>('/purchase-orders')) }
    catch { setError('Failed to load purchase orders') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreatePurchaseOrderInput) => {
    await apiPost<PurchaseOrder>('/purchase-orders', data); await load()
  }
  const updateStatus = async (id: number, status: PurchaseOrderStatus) => {
    await apiPatch<PurchaseOrder>(`/purchase-orders/${id}/status`, { status }); await load()
  }
  const receive = async (id: number) => {
    await apiPost<PurchaseOrder>(`/purchase-orders/${id}/receive`, {}); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/purchase-orders/${id}`); await load()
  }

  return { orders, loading, error, load, create, updateStatus, receive, remove }
}
