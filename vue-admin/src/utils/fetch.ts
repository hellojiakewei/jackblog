// import { urlProp } from '@/interface/urlProp'
import axios, { Canceler, AxiosRequestConfig, AxiosResponse } from 'axios'
import qs from 'qs'
// import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import store from '@/store/index'
const router = useRouter()
// if(process.env.Env === 'development'){
//     axios.defaults.baseURL = process.env.
// }
const service = axios.create({
    timeout: 40000,
    // transformRequest: [data => qs.stringify(data)],
    headers: {
        'Content-Type': 'application/json'
    },
    // withCredentials: true
});
console.log(process.env.VUE_APP_BLOG_ADMIN)
service.defaults.baseURL = process.env.VUE_APP_BLOG_ADMIN

let sourceAjaxList: Canceler[] = []

service.interceptors.request.use((config: AxiosRequestConfig) => {
    config.cancelToken = new axios.CancelToken(function executor(cancel: Canceler): void {
        sourceAjaxList.push(cancel)
    })
    if (store.state.token) {
        config.headers['Authorization'] = `Bearer ${store.state.token}`;
    }
    return config
}, function (error) {
    // 抛出错误
    return Promise.reject(error)
})
service.interceptors.response.use((response: AxiosResponse) => {
    const { status, data } = response
    if (status === 200) {
        return data
    } else if (data.code === 401) {
        sourceAjaxList.length && sourceAjaxList.length > 0 && sourceAjaxList.forEach((ajaxCancel, index) => {
            ajaxCancel()
            delete sourceAjaxList[index]
        })
        ElMessage({
            showClose: true,
            message: response.data,
            type: 'error'
        });
        return router.push('/login')
    } else {
        return data
    }
}, error => {
    const { response } = error
    if (!response || response.status === 404 || response.status === 500) {
        if (!response) {
            console.error(`404 error %o ${error}`)
        } else {
            if (response.data && response.data.message) {
                ElMessage({
                    message: '请求异常，请稍后再试!',
                    duration: 2000,
                    type: 'error'
                })
            }
        }
    }
    return Promise.reject(error.message)
})
export default service

