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
import {
  getEventInfo,
  tryGetEventInfo,
} from '../../../../core/api/event/getInfo';
import { useNoticeStore } from '../../../../store/noticeStore';
import * as TE from 'fp-ts/TaskEither';
import { pipe } from 'fp-ts/lib/function';

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
export default function Edit(event: Event) {
  const { changeNotice } = useNoticeStore();
  const [name, setName] = useState(event.event_name);
  const [capacity, setCapacity] = useState(event.event_size);
  const [detail, setDetail] = useState(event.detail);
  const [location, setLocation] = useState<null | TMapPosition>(null);
  const eventId = useRouter().query.eventId;
  const update = updateEvent(
    (ok) => {
      changeNotice({ type: 'Success', text: 'イベント情報を更新しました' });
    },
    (err) => {
      changeNotice({ type: 'Error', text: '更新に失敗しました' });
    }
  );

  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">編集する</h1>
      </TypoWrapper>

      <FormWrapper
        onSubmit={() =>
          update({
            id: eventId as string,
            name: name,
            detail: detail,
            location: JSON.stringify(location),
            size: Number(capacity),
            type: '???',
            state: '???',
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
