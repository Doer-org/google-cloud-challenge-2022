import create from 'zustand';
import { persist } from 'zustand/middleware';

// interface UserInfo {
//   userId: string;
// }
// miso
// a88d4cba-6211-40ee-8a23-3f259d0166d5
// const initUserInfo: UserInfo = {
//   userId: 'a88d4cba-6211-40ee-8a23-3f259d0166d5',
// };
// 動作しなかったので一旦置き換え
// interface UserInfoStore {
//   userInfo: UserInfo;
//   setUserInfo: (info: UserInfo) => void;
// }

// export const useUserInfoStore = create<UserInfoStore>()(
//   persist((set) => ({
//     userInfo: initUserInfo,
//     setUserInfo: (info: UserInfo) => set({ userInfo: info }),
//   }))
// );

type UserInfoStore = {
  userId: string;
  setUserId: (id: string) => void;
};
export const UserStore = create<UserInfoStore>((set) => ({
  userId: 'c698649d-9f54-4bf7-9e0f-5a715fc04909',
  setUserId: (id: string) => {
    return {
      userId: id,
    };
  },
}));
