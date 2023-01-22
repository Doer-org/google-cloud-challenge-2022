import { useState } from 'react';
import { UserComment } from '../../molecules/User/UserComment';
import { UserIcon } from '../../molecules/User/UserIcon';
import { UserName } from '../../molecules/User/UserName';
import { UserModal } from '../../templates/shared/Modal/UserModal';

type TProps = {
  name: string;
  comment?: string;
  image?: string;
  isParticipate?: boolean;
};

export const UserInfo = ({ name, comment, image, isParticipate }: TProps) => {
  const [isShowModal, setIsShowModal] = useState(false);
  return (
    <UserModal
      isShow={isShowModal}
      onClose={setIsShowModal}
      userInfo={{ name, comment, image, isParticipate }}
    >
      <div
        className={`${isParticipate ? 'md:w-20 w-24' : 'md:w-32 w-20'}`}
        onClick={() => setIsShowModal(true)}
      >
        {isParticipate ? <UserComment comment={comment} /> : <></>}
        <UserIcon isParticipate={isParticipate} />
        <UserName name={name} isParticipate={isParticipate} />
      </div>
    </UserModal>
  );
};
