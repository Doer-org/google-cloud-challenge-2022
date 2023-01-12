import { Button } from '../../atoms/text/Button';
import { TypoWrapper } from '../../atoms/text/TypoWrapper';
type TProps = {
  eventId: string;
  onCancel: () => void;
  onParticipate: () => void;
};
export const EventConfirmModal = ({
  onParticipate,
  onCancel,
  eventId,
}: TProps) => {
  return (
    <div className="h-screen flex justify-center flex-col gap-5">
      <p>{eventId}</p>
      <TypoWrapper>
        <h1>本当に参加しますか？</h1>
      </TypoWrapper>
      <Button onClick={onParticipate}>参加する</Button>
      <Button onClick={onCancel}>考え直す</Button>
    </div>
  );
};
