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
  changeNotice: () =>
    set((state) => {
      return { notice: { type: state.notice.type, text: state.notice.text } };
    }),
  resetNotice: () =>
    set((state) => {
      return { notice: { type: 'None', text: '' } };
    }),
}));
