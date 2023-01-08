import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  userName: string;
};
export const UserIcon = ({ userName }: TProps) => {
  return (
    <div className="my-5 m-auto">
      <div className="w-14 h-14 rounded-full bg-orange-500 m-auto my-1"></div>
      <TypoWrapper size="small" line="shin">
        <p>{userName}</p>
      </TypoWrapper>
    </div>
  );
};
