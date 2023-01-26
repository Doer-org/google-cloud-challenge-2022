import { ReactNode } from 'react';
import { Button } from '../../../atoms/text/Button';
import { UserFullInfo } from '../../../organisms/User/UserFullInfo';
import { BasicModal } from './BasicModal';

type TProps = {
  children: ReactNode;
  isShow: boolean;
  onClose: (isShow: boolean) => void;
  userInfo: {
    name: string;
    comment?: string;
    image: string;
    isParticipate?: boolean;
  };
};
export const UserModal = ({ children, isShow, onClose, userInfo }: TProps) => {
  return (
    <>
      {isShow ? (
        <>
          {children}
          <BasicModal>
            <UserFullInfo
              name={userInfo.name}
              comment={userInfo.comment}
              image={userInfo.image}
              isParticipate={userInfo.isParticipate}
            />
            <Button
              onClick={() => onClose(false)}
              className="text-white text-lg border-none my-5 hover:opacity-75"
            >
              閉じる
            </Button>
          </BasicModal>
        </>
      ) : (
        <>{children}</>
      )}
    </>
  );
};
