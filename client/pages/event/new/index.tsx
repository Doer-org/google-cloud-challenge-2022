import { useState } from 'react';
import Link from 'next/link';
import useHostApi from '../../../core/hooks/useEventHost';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';

export default function New() {
  const { createNewEvent } = useHostApi();
  const [createdEventId, setCreatedEventId] = useState(0);
  const createEvent = createNewEvent((v) => {
    console.log(v);
    setCreated(true);
    setCreatedEventId(v.created_event.event_id);
  }, console.log);

  const [created, setCreated] = useState(false);
  const [event_name, setEventName] = useState('event name');
  const [max_n, setMaxN] = useState(10);
  const [detail, setDetail] = useState('hogehogefugafuga');
  const [location, setLocation] = useState('earth');

  return (
    <BasicTemplate className="text-center">
      <h1>募集する</h1>
      <h2>イベント名</h2>
      <input
        value={event_name}
        onChange={(e) => setEventName(e.target.textContent ?? '')}
      />
      <h2>募集人数</h2>
      <input
        value={max_n}
        onChange={(e) => setMaxN(Number(e.target.textContent) ?? 0)}
      />
      <h2>詳細</h2>
      <input
        value={detail}
        onChange={(e) => setDetail(e.target.textContent ?? '')}
      />
      <h2>場所</h2>
      <input
        value={location}
        onChange={(e) => setLocation(e.target.textContent ?? '')}
      />
      <button
        onClick={() =>
          createEvent(
            { user_id: 'abc' },
            {
              event_name: event_name,
              max_member: max_n,
              detail: detail,
              location: location,
            }
          )
        }
      >
        募集
      </button>
      {created ?? (
        <Link href={`/event?id=${createdEventId}`}>募集したイベント</Link>
      )}
    </BasicTemplate>
  );
}
