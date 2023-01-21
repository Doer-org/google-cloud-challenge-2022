import { Participant } from '../../../core/types/event';
import { Button } from '../../atoms/text/Button';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  eventId: string;
  participants: Participant[];
  onCancel: () => void;
  onParticipate: () => void;
};
export const EventConfirmModal = ({
  participants,
  onParticipate,
  onCancel,
  eventId,
}: TProps) => {
  return (
    <div className="h-screen flex justify-center flex-col gap-5">
      <div className="grid md:grid-cols-6 grid-cols-4 items-end">
        <div>
          <TypoWrapper line="bold">ホスト</TypoWrapper>
        </div>
      </div>
      <TypoWrapper>
        <h1>本当に参加しますか？</h1>
      </TypoWrapper>
      <Button onClick={onParticipate}>参加する</Button>
      <Button onClick={onCancel}>考え直す</Button>
    </div>
  );
};
