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
import { useEffect } from 'react';
import { EventConfirmModal } from '../../../../components/templates/shared/Modal/EventConfirmModal';
import { useNoticeStore } from '../../../../store/noticeStore';

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
  const [origin, setOrigin] = useState('');
  useEffect(() => {
    setOrigin(window.location.origin);
  }, []);
  console.log(event);
  const joinApi = joinEvent(
    (response) => {
      console.log(response)
      router.push(`${origin}/event/${event.event_id}/`);
      changeNotice({ type: 'Success', text: '参加しました！' });
    },
    (e) => {
      router.reload()
      setIsConfirm(false) // TODO:
      changeNotice({ type: 'Error', text: '参加に失敗しました' });
    }
  );
  const [name, setName] = useState('');
  const [comment, setComment] = useState('');
  const isCapacityOver = event.participants.length >= event.event_size; 
  console.log("参加者", event.participants.length, ", event_size", event.event_size, ", isOver", isCapacityOver)
  const isTimeOver = new Date(event.close_limit) < new Date();
  const isClosed = event.event_state === 'close';
  return (
    <>
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
          currentUser={{ name, comment, icon: '' }}
          participant={event.participants.map((participant) => {
            return {
              name: participant.participant_name,
              image: participant.icon,
            };
          })}
        >
          <div>
            <EventInfo
              participants={event.participants}
              eventName={event.event_name}
              detail={event.detail}
              location={event.location}
              hostImage={event.host.icon}
              hostName={event.host.user_name}
              limitTime={event.close_limit}
            />
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
          </div>
        </EventConfirmModal>
      </BasicTemplate>
    </>
  );
}
