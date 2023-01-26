import { useNoticeStore } from '../../../../store/noticeStore';
import { Button } from '../../../atoms/text/Button';
import { TypoWrapper } from '../../../atoms/text/TypoWrapper';
type TProps = {
  type: 'Error' | 'Success' | 'None';
  text: string;
};
export const Notice = ({ type, text }: TProps) => {
  const { notice, changeNotice, resetNotice } = useNoticeStore();
  const isError = type === 'Error';
  if (type === 'None') {
    return null;
  } else {
    setTimeout(() => {
      resetNotice();
    }, 5000);
  }

  return (
    <div
      className={`absolute top-6 right-1/2 py-1 rounded-md md:w-1/3 w-2/3 transition translate-x-[50%] shadow-2xl border-black border-2 text-black z-50 ${
        isError ? 'bg-red-400' : 'bg-green-400'
      }`}
    >
      <div className="relative">
        <Button onClick={() => resetNotice()} className="absolute left-0 p-0">
          ×
        </Button>
        <TypoWrapper>
          <p className="text-center">{text}</p>
        </TypoWrapper>
      </div>
    </div>
  );
};
