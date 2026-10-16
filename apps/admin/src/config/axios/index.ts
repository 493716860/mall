import { ElMessage } from 'element-plus'
import { service } from './service'

import { config } from './config'

const { default_headers } = config

const request = async (option: any) => {
  if (import.meta.env.VITE_DEMO_MODE !== 'false') {
    const { demoRequest } = await import('@/demo')
    return { code: 0, ...(await demoRequest(option)) }
  }
  const { headersType, headers, ...otherOption } = option
  return service({
    ...otherOption,
    headers: {
      'Content-Type': headersType || default_headers,
      ...headers
    }
  })
}
export default {
  get: async <T = any>(option: any) => {
    const res = await request({ method: 'GET', ...option })
    return res.data as unknown as T
  },
  post: async <T = any>(option: any) => {
    const res = await request({ method: 'POST', ...option })
    return res.data as unknown as T
  },
  postOriginal: async (option: any) => {
    const res = await request({ method: 'POST', ...option })
    return res
  },
  delete: async <T = any>(option: any) => {
    const res = await request({ method: 'DELETE', ...option })
    return res.data as unknown as T
  },
  put: async <T = any>(option: any) => {
    const res = await request({ method: 'PUT', ...option })
    return res.data as unknown as T
  },
  download: async <T = any>(option: any) => {
    if (import.meta.env.VITE_DEMO_MODE === 'true') {
      ElMessage.warning('导出需要正式后端接口')
      throw new Error('演示模式不执行后端导出')
    }
    const res = await request({ method: 'GET', responseType: 'blob', ...option })
    return res as unknown as Promise<T>
  },
  upload: async <T = any>(option: any) => {
    option.headersType = 'multipart/form-data'
    const res = await request({ method: 'POST', ...option })
    return res as unknown as Promise<T>
  }
}
