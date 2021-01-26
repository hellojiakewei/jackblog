// import { urlProp } from '@/interface/urlProp'
import axios from 'axios'
import qs from 'qs'
import { useStore } from 'vuex'
import { ElLoading } from 'element-plus'
const store = useStore()
const service = axios.create({
    timeout: 4000,
    timeoutErrorMessage: "请求超时",
    transformRequest: [data => qs.stringify(data)],
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
})

service.interceptors.request.use((config) => {
    return config
}, (error: any) => {
    return Promise.reject(error.message)
})
service.interceptors.response.use((response) => {
    return response.data
}, error => {
    return Promise.reject(error.message)
})


export default function () {


}