import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  userName: string;
  image?: string;
};
// TODO:ユーザー画像を入れる
export const UserIcon = ({ userName }: TProps) => {
  return (
    <div className="m-auto">
      <div className="w-8 h-8 rounded-full bg-orange-500 m-auto my-1"></div>
      <TypoWrapper size="small" line="shin">
        <p>{userName}</p>
      </TypoWrapper>
    </div>
  );
};
