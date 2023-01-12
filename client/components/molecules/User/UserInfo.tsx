import { TypoWrapper } from '../../atoms/text/TypoWrapper';
import { UserIcon } from './UserIcon';

type TProps = {
  name: string;
  comment: string;
  image: string;
};

export const UserInfo = ({ name, comment, image }: TProps) => {
  return (
    <>
      <UserIcon userName={name} />

      <TypoWrapper>
        <p>{comment}</p>
      </TypoWrapper>
    </>
  );
};