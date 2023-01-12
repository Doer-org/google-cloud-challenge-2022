import { useEffect, useState } from 'react';
import Router, { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { useCopyToClipboard } from 'usehooks-ts';
import { Event } from '../../../../core/types/event';
import { getEventInfo } from '../../../../core/api/event/getInfo';
export default function Participate() {
  // ここはイベントを作成したときにリンクをコピーする画面
  // TODO:URLコピー機能
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
  const [value, copy] = useCopyToClipboard();
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

  return (
    <>
      <MyHead title="イベントURLコピー" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
        />
        <Button
          className="flex m-auto my-5"
          onClick={() => {
            copy(
              `${process.env.NEXT_PUBLIC_SERVER_URL}event/${eventId}/participate`
            );
          }}
        >
          URLをコピー
        </Button>
      </BasicTemplate>
    </>
  );
}
