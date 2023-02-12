import { useState } from 'react';
import { useRouter } from 'next/router';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { tryGetEventInfo } from '../../../../core/api/event/getInfo';
import { Event } from '../../../../core/types/event';
import { pipe } from 'fp-ts/lib/function';
import * as TE from 'fp-ts/TaskEither';
import { joinEvent } from '../../../../core/api/event/join';
import { EventConfirmModal } from '../../../../components/templates/shared/Modal/EventConfirmModal';
import { useNoticeStore } from '../../../../store/noticeStore';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';

export async function getServerSideProps(context: any) {
  const eventId = context.query.eventId;
  return pipe(
    eventId,
    tryGetEventInfo,
    TE.match(
      (err) => {
        throw err;
      },
      (response) => {
        return {
          props: {
            ...response,
          },
        };
      }
    )
  )();
}

export default function Participate(event: Event) {
  // TODO: SSRで実装してリンクを貼った時にOGPを表示させるようにする
  const [isConfirm, setIsConfirm] = useState(false);
  const { changeNotice } = useNoticeStore();
  const router = useRouter();
  const joinApi = joinEvent(
    () => {
      router.push(
        `${process.env.NEXT_PUBLIC_FRONT_URL}/event/${event.event_id}`
      );
      changeNotice({ type: 'Success', text: '参加しました！' });
    },
    () => {
      router.reload();
      setIsConfirm(false);
      changeNotice({ type: 'Error', text: '参加に失敗しました' });
    }
  );
  const [name, setName] = useState('');
  const [comment, setComment] = useState('');
  const isCapacityOver = event.participants.length >= event.event_size;
  const isTimeOver = new Date(event.close_limit) < new Date();
  const isClosed = event.event_state === 'close';
  return (
    <>
      <MyHead title={event.event_name} description={event.detail} />
      <BasicTemplate className="text-center">
        <EventConfirmModal
          isShow={isConfirm}
          onClose={setIsConfirm}
          onParticipate={() =>
            joinApi({
              event_id: event.event_id,
              participant_name: name,
              comment: comment,
            })
          }
          eventId={event.event_id}
          currentUser={{ participant_name: name, comment, icon: '' }}
        >
          <EventInfo event={event} />
          {!isCapacityOver && !isTimeOver && !isClosed ? (
            <FormWrapper onSubmit={() => setIsConfirm(true)}>
              <Input
                type="text"
                label="名前"
                maxLength={20}
                minLength={1}
                content={name}
                changeContent={setName}
                required={true}
              />
              <Input
                type="text"
                label="ひとこと"
                maxLength={50}
                content={comment}
                changeContent={setComment}
              />
              <Button className="flex m-auto my-5">参加する</Button>
            </FormWrapper>
          ) : (
            <>
              <p className="mb-5">このイベントは締切られました</p>
            </>
          )}
        </EventConfirmModal>
      </BasicTemplate>
    </>
  );
}
