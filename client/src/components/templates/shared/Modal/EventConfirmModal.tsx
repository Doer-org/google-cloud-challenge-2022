import { ReactNode, useEffect, useState } from 'react';
import { getEventInfo } from '../../../../core/api/event/getInfo';
import { Participant } from '../../../../core/types/event';
import { Button } from '../../../atoms/text/Button';
import { TypoWrapper } from '../../../atoms/text/TypoWrapper';
import { UserFullInfo } from '../../../organisms/User/UserFullInfo';
import { BasicModal } from './BasicModal';
import { UserInfo } from '../../../organisms/User/UserInfo';
type TProps = {
  children: ReactNode;
  isShow: boolean;
  currentUser: Participant;
  eventId: string;
  onClose: (isShow: boolean) => void;
  onParticipate: () => void;
};
export const EventConfirmModal = ({
  children,
  isShow,
  onClose,
  currentUser,
  onParticipate,
  eventId,
}: TProps) => {
  const [participants, setParticipants] = useState<Participant[]>([]);
  useEffect(() => {
    const getParticipants = getEventInfo(
      (response) => {
        if (response) {
          setParticipants(response.participants);
        }
      },
      () => {}
    );
    getParticipants(eventId);
  }, [eventId]);
  return (
    <>
      {isShow ? (
        <div>
          {children}
          <BasicModal>
            <UserFullInfo
              name={currentUser.participant_name}
              comment={currentUser.comment}
              isParticipate
              image={currentUser.icon}
            />
            <TypoWrapper>
              <h1 className="mt-10">他にも以下の参加者がいます</h1>
            </TypoWrapper>
            <div className="w-2/3 overflow-x-scroll m-auto my-3">
              <div className="flex gap-5 justify-center">
                {participants.map((participant) => {
                  return (
                    <UserInfo
                      key={participant.participant_name}
                      name={participant.participant_name}
                      image={participant.icon}
                    />
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
