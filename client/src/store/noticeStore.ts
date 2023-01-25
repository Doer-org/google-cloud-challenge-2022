import create from 'zustand';
type NoticeInfo = {
  type: 'Error' | 'Success' | 'None';
  text: string;
};
type Notice = {
  notice: NoticeInfo;
  changeNotice: (state: NoticeInfo) => void;
  resetNotice: () => void;
};

export const useNoticeStore = create<Notice>((set) => ({
  notice: { type: 'None', text: '' },
  changeNotice: (state) => {
    set({ notice: { type: state.type, text: state.text } });
  },
  resetNotice: () => {
    set({ notice: { type: 'None', text: '' } });
  },
}));
