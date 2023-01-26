import { ReactNode } from 'react';
import { Button } from '../../../atoms/text/Button';
import { TypoWrapper } from '../../../atoms/text/TypoWrapper';
import { UserFullInfo } from '../../../organisms/User/UserFullInfo';
import { UserInfo } from '../../../organisms/User/UserInfo';
import { BasicModal } from './BasicModal';

type TProps = {
  children: ReactNode;
  isShow: boolean;
  onClose: (isShow: boolean) => void;
  onParticipate: () => void;
  currentUser: {
    name: string;
    comment?: string;
    icon: string;
  };
  participant: {
    name: string;
    image: string;
  }[];
  isParticipate?: boolean;
};
export const EventConfirmModal = ({
  children,
  isShow,
  onClose,
  currentUser,
  participant,
  onParticipate,
}: TProps) => {
  return (
    <>
      {isShow ? (
        <div>
          {children}
          <BasicModal>
            <UserFullInfo
              name={currentUser.name}
              comment={currentUser.comment}
              isParticipate
              image={currentUser.icon}
            />
            <TypoWrapper>
              <h1 className="mt-10">他にも以下の参加者がいます</h1>
            </TypoWrapper>
            <div className="w-2/3 overflow-x-scroll m-auto my-3">
              <div className="flex gap-5">
                {participant.map((p) => {
                  return (
                    <UserInfo key={p.name} name={p.name} image={p.image} />
                  );
                })}
              </div>
            </div>

            <div className="m-auto w-full">
              <Button onClick={onParticipate} className="mx-3">
                参加する
              </Button>
              <Button onClick={() => onClose(false)} className="mx-3">
                考え直す
              </Button>
            </div>
          </BasicModal>
        </div>
      ) : (
        <>{children}</>
      )}
    </>
  );
};
