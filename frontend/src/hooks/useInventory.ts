import { useState, useCallback } from 'react'
import { apiGet, apiPost, apiPut, apiDelete } from '../api/client'
import type {
  Category, Product, StockMovement,
  CreateCategoryInput, CreateProductInput, UpdateProductInput, StockAdjustInput,
} from '../types'

export function useCategories() {
  const [categories, setCategories] = useState<Category[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setCategories(await apiGet<Category[]>('/categories')) }
    catch { setError('Failed to load categories') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateCategoryInput) => {
    await apiPost<Category>('/categories', data); await load()
  }

  return { categories, loading, error, load, create }
}

export function useProducts() {
  const [products, setProducts] = useState<Product[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    setLoading(true); setError('')
    try { setProducts(await apiGet<Product[]>('/products')) }
    catch { setError('Failed to load products') }
    finally { setLoading(false) }
  }, [])

  const create = async (data: CreateProductInput) => {
    await apiPost<Product>('/products', data); await load()
  }
  const update = async (id: number, data: UpdateProductInput) => {
    await apiPut<Product>(`/products/${id}`, data); await load()
  }
  const remove = async (id: number) => {
    await apiDelete(`/products/${id}`); await load()
  }
  const adjustStock = async (data: StockAdjustInput) => {
    await apiPost<void>('/stock/adjust', data); await load()
  }
  const getMovements = async (productId: number): Promise<StockMovement[]> => {
    return apiGet<StockMovement[]>(`/products/${productId}/movements`)
  }

  return { products, loading, error, load, create, update, remove, adjustStock, getMovements }
}
