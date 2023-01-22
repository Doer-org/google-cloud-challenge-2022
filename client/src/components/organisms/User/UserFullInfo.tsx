import { UserComment } from '../../molecules/User/UserComment';
import { UserIcon } from '../../molecules/User/UserIcon';
import { UserName } from '../../molecules/User/UserName';

type TProps = {
  name: string;
  comment?: string;
  image?: string;
  isParticipate?: boolean;
};

export const UserFullInfo = ({
  name,
  comment,
  image,
  isParticipate,
}: TProps) => {
  return (
    <div className="w-screen">
      {isParticipate ? <UserComment comment={comment} full /> : <></>}
      <UserIcon isParticipate={isParticipate} />
      <UserName name={name} full />
    </div>
  );
};
