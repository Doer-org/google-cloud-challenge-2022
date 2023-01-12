import { Participant } from '../../../core/types/event';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
import { UserIcon } from './UserIcon';

type TProps = {
  participants: Participant[];
};

export const ParticipateInfo = ({ participants }: TProps) => {
  return (
    <>
      {participants.map((participant) => {
        return (
          <div className="rounded-md" key={participant.participant_name}>
            <UserIcon userName={participant.participant_name} />
            <TypoWrapper size="small" line="shin">
              <p className="text-left overflow-x-scroll w-full whitespace-nowrap">
                {participant.comment}
              </p>
            </TypoWrapper>
          </div>
        );
      })}
    </>
  );
};
