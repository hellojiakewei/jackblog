import { Iuserinfo } from '@/model/user'
import fetch from '@/utils/fetch'

// 登录
export function loginApi(data:Iuserinfo){
    return fetch({
        url:"user/login",
        method:"POST",
        data,
    })
}