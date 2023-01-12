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
  userId: '5a288c56-09fd-422a-9753-a73216f64f01',
  setUserId: (id: string) => {
    return {
      userId: id,
    };
  },
}));
