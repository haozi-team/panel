import type { AxiosResponse } from 'axios'

import { request } from '@/utils'

export default {
  // 获取容器列表
  containerList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get('/container/container', { params: { page, limit } }),
  // 添加容器
  containerCreate: (config: any): Promise<AxiosResponse<any>> =>
    request.post('/container/container', config),
  // 删除容器
  containerRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.delete(`/container/container/${id}`),
  // 启动容器
  containerStart: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/start`),
  // 停止容器
  containerStop: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/stop`),
  // 重启容器
  containerRestart: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/restart`),
  // 暂停容器
  containerPause: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/pause`),
  // 恢复容器
  containerUnpause: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/unpause`),
  // 杀死容器
  containerKill: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/kill`),
  // 重命名容器
  containerRename: (id: string, name: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/container/${id}/rename`, { name }),
  // 获取容器日志
  containerLogs: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/container/${id}/logs`),
  // 清理容器
  containerPrune: (): Promise<AxiosResponse<any>> => request.post(`/container/container/prune`),
  // 获取网络列表
  networkList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get(`/container/network`, { params: { page, limit } }),
  // 创建网络
  networkCreate: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/network`, config),
  // 删除网络
  networkRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.delete(`/container/network/${id}`),
  // 清理网络
  networkPrune: (): Promise<AxiosResponse<any>> => request.post(`/container/network/prune`),
  // 获取镜像列表
  imageList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get(`/container/image`, { params: { page, limit } }),
  // 拉取镜像
  imagePull: (config: any): Promise<AxiosResponse<any>> => request.post(`/container/image`, config),
  // 删除镜像
  imageRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.delete(`/container/image/${id}`),
  // 清理镜像
  imagePrune: (): Promise<AxiosResponse<any>> => request.post(`/container/image/prune`),
  // 获取卷列表
  volumeList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get(`/container/volume`, { params: { page, limit } }),
  // 创建卷
  volumeCreate: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/volume`, config),
  // 删除卷
  volumeRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.delete(`/container/volume/${id}`),
  // 清理卷
  volumePrune: (): Promise<AxiosResponse<any>> => request.post(`/container/volume/prune`)
}
