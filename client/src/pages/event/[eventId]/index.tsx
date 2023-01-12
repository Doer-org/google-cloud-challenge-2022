import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { EventInfo } from '../../../components/molecules/EventInfo';
import { getEventInfo } from '../../../core/api/event/getInfo';
import { Event } from '../../../core/types/event';

export default function Show() {
  // TODO: Comment取れてねーじゃん！
  const [event, setEvent] = useState<Event>({
    event_id: '',
    event_name: '',
    detail: '',
    location: '',
    host: {
      user_id: '',
    },
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
