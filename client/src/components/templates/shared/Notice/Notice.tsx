import { TypoWrapper } from '../../../atoms/text/TypoWrapper';
type TProps = {
  type: 'Error' | 'Success' | 'None';
  text: string;
};
// TODO: typeをNormalにしたときは非表示にする
// displaynoneとかだとanimationが出ないのでなんか考えとく、最悪アニメーションなしで
// 3秒後とかにNormalにするとかを考えとく（これは上のコンポーネントでtypeがNormalの場合は出さないようにする）
export const Notice = ({ type, text }: TProps) => {
  const isError = type === 'Error';
  if (type === 'None') {
    return null;
  }
  return (
    <div
      className={`absolute top-5 right-3 py-1 px-3 rounded-md lg:w-1/6 w-1/3 transition ${
        isError ? 'bg-red-400' : 'bg-green-400'
      }`}
    >
      <TypoWrapper>
        <p className="text-center">{text}</p>
      </TypoWrapper>
    </div>
  );
};
