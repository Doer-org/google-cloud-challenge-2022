import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { Event } from '../../../../core/types/event';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { closeEvent } from '../../../../core/api/event/close';
import { getEventInfo } from '../../../../core/api/event/getInfo';

export default function Close() {
  // ここは締め切るページ
  // TODO:event締切hooksをonClickへ
  const [event, setEvent] = useState<Event>({
    event_id: '',
    event_name: '',
    detail: '',
    location: '',
    host: {
      user_id: '',
      user_name: '',
      icon: '',
    },
    event_size: 1,
    event_state: '',
    participants: [],
  });
  const getEvent = getEventInfo(
    (response) => {
      setEvent(response);
    },
    (error) => {}
  );
  const eventId = useRouter().query.eventId;

  useEffect(() => {
    getEvent(eventId as string);
  }, [getEvent, eventId]);
  const [result, setResult] = useState('not closed');
  const close = closeEvent(
    (ok) => {
      setResult('ok!');
    },
    (err) => {
      setResult('error!');
    }
  );

  return (
    <>
      <MyHead title="イベント締切ページ" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
        />
        <Button
          className="flex m-auto my-5"
          onClick={() => close(eventId as string)}
        >
          締め切る
        </Button>
        <>{result}</>
      </BasicTemplate>
    </>
  );
}
