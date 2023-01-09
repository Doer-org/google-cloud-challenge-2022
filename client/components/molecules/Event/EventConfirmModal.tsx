import { Button } from '../../atoms/text/Button';
type TProps = {
  eventId: number;
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
      <Button onClick={onParticipate}>本当に参加しますか？</Button>
      <Button onClick={onCancel}>考え直す</Button>
    </div>
  );
};
