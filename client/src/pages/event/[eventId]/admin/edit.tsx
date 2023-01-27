import { useEffect, useState } from 'react';
import { Button } from '../../../../components/atoms/text/Button';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { MapForm } from '../../../../components/atoms/form/MapForm';
import { Textarea } from '../../../../components/atoms/form/Textarea';
import { TypoWrapper } from '../../../../components/atoms/text/TypoWrapper';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { updateEvent } from '../../../../core/api/event/update';
import { useRouter } from 'next/router';
import { TMapPosition } from '../../../../components/atoms/map/MapBasicInfo';
import { Event } from '../../../../core/types/event';
import { tryGetEventInfo } from '../../../../core/api/event/getInfo';
import { useNoticeStore } from '../../../../store/noticeStore';
import * as TE from 'fp-ts/TaskEither';
import { pipe } from 'fp-ts/lib/function';

export default function Edit(event: Event) {
  const { changeNotice } = useNoticeStore();
  const [name, setName] = useState(event.event_name);
  const [capacity, setCapacity] = useState(event.event_size);
  const [detail, setDetail] = useState(event.detail);
  const [location, setLocation] = useState<null | TMapPosition>(null);
  const [limit, setLimit] = useState(event.close_limit);
  const router = useRouter();
  const update = updateEvent(
    (ok) => {
      router.push('/event');
      changeNotice({ type: 'Success', text: 'イベント情報を更新しました' });
    },
    (err) => {
      changeNotice({ type: 'Error', text: '更新に失敗しました' });
    }
  );
  let now = new Date();
  let tomorrow = new Date();
  tomorrow.setDate(now.getDate() + 1);
  tomorrow.setHours(23, 59, 59, 999);
  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">編集する</h1>
      </TypoWrapper>

      <FormWrapper
        onSubmit={() =>
          update({
            id: event.event_id,
            name: name,
            detail: detail,
            location: JSON.stringify(location),
            size: Number(capacity),
            type: '???',
            state: '???',
            created_at: new Date(Date.now()),
            limit_time: new Date(Date.now()), // FIXME: 締め切り時間設定
          })
        }
      >
        <Input
          type="text"
          label="イベント名"
          minLength={1}
          maxLength={100}
          content={name}
          changeContent={setName}
          required={true}
        />
        <Input
          type="number"
          min={1}
          max={5}
          label="募集人数"
          content={capacity}
          changeContent={setCapacity}
          required={true}
        />
        <Input
          type="datetime-local"
          label="締切"
          min={
            now.toISOString().slice(0, 11) +
            now.toLocaleTimeString().slice(0, 5)
          }
          max={
            tomorrow.toISOString().slice(0, 11) +
            tomorrow.toLocaleTimeString().slice(0, 5)
          }
          content={limit}
          changeContent={setLimit}
          required={true}
        />
        <Textarea
          label="詳細"
          minLength={1}
          maxLength={500}
          content={detail}
          changeContent={setDetail}
          required={true}
        />
        <MapForm location={location} setLocation={setLocation} />
        <Button className="flex m-auto my-5">編集完了する</Button>
      </FormWrapper>
    </BasicTemplate>
  );
}

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
