import Head from 'next/head';
import { Inter } from '@next/font/google';
import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { EventInfo } from '../../../components/molecules/EventInfo';
import { getEventInfo } from '../../../core/api/event/getInfo';
import { Event } from '../../../core/types/event';
const inter = Inter({ subsets: ['latin'] });

export default function Show() {
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
  console.log(event);
  const eventId = useRouter().query.eventId;
  const getEvent = getEventInfo(
    (response) => {
      setEvent(response);
    },
    (error) => {
      console.log(error);
    }
  );
  useEffect(() => {
    getEvent(eventId as string);
  }, []);
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
        />
      </BasicTemplate>
    </>
  );
}
