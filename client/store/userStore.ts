import create from "zustand"
import { persist } from 'zustand/middleware'

interface UserInfo { 
    userId : string 
}
const initUserInfo : UserInfo = { 
    userId : ""
}

interface UserInfoStore {
    userInfo : UserInfo
    setUserInfo : (info : UserInfo) => void
}

export const useUserInfoStore = create<UserInfoStore>()(
    persist(   
        (set) => ({ 
            userInfo : initUserInfo,
            setUserInfo : (info : UserInfo) => set(({userInfo: info}))
        })
    ))