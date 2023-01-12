import create from 'zustand';
import { persist } from 'zustand/middleware';

interface UserInfo {
  userId: string;
}
// miso
// a88d4cba-6211-40ee-8a23-3f259d0166d5
const initUserInfo: UserInfo = {
  userId: 'a88d4cba-6211-40ee-8a23-3f259d0166d5',
};

interface UserInfoStore {
  userInfo: UserInfo;
  setUserInfo: (info: UserInfo) => void;
}

export const useUserInfoStore = create<UserInfoStore>()(
  persist((set) => ({
    userInfo: initUserInfo,
    setUserInfo: (info: UserInfo) => set({ userInfo: info }),
  }))
);
