import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { EventInfo } from '../../../components/molecules/EventInfo';
import { getEventInfo, tryGetEventInfo } from '../../../core/api/event/getInfo';
import { Event } from '../../../core/types/event';
import { pipe } from 'fp-ts/lib/function';
import * as TE from 'fp-ts/TaskEither';
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

export default function Show() {
  // TODO: Comment取れてねーじゃん！
  const [event, setEvent] = useState<Event>({
    event_id: '',
    event_name: '',
    detail: '',
    location: '',
    host: {
      user_name: '',
      user_id: '',
      icon: '',
    },
    event_size: 1,
    event_state: '',
    participants: [],
  });

  const eventId = useRouter().query.eventId;
  const getEvent = getEventInfo(
    (response) => {
      setEvent(response);
    },
    (error) => {}
  );
  useEffect(() => {
    getEvent(eventId as string);
  }, []);
  // TODO: commentだけ空文字で返ってくる
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          participants={event.participants}
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
        />
      </BasicTemplate>
    </>
  );
}
