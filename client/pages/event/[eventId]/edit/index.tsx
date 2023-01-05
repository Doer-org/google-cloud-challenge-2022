import { useState } from 'react';
import { Button } from '../../../../components/atoms/Button';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { MapForm } from '../../../../components/atoms/form/MapForm';
import { Textarea } from '../../../../components/atoms/form/Textarea';
import { TypoWrapper } from '../../../../components/atoms/TypoWrapper';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import useHostApi from '../../../../core/hooks/useEventHost';

export default function Edit() {
  const { createNewEvent } = useHostApi();
  const [createdEventId, setCreatedEventId] = useState(0);
  const createEvent = createNewEvent((v) => {
    console.log(v);
    setCreated(true);
    setCreatedEventId(v.created_event.event_id);
  }, console.log);

  const [created, setCreated] = useState(false);
  const [name, setName] = useState('');
  const [capacity, setCapacity] = useState(1);
  const [detail, setDetail] = useState('');
  const [location, setLocation] = useState('');
  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">編集する</h1>
      </TypoWrapper>

      <FormWrapper>
        <Input
          type="text"
          label="イベント名"
          content={name}
          changeContent={setName}
          required={true}
        />
        <Input
          type="number"
          label="募集人数"
          content={capacity}
          changeContent={setCapacity}
          required={true}
        />
        <Textarea
          label="詳細"
          content={detail}
          changeContent={setDetail}
          required={true}
        />
        <MapForm />
        <Button
          className="flex m-auto my-5"
          onClick={() => {
            createEvent(
              { user_id: 'abc' },
              {
                event_name: name,
                max_member: capacity,
                detail: detail,
                location: location,
              }
            );
          }}
        >
          編集完了する
        </Button>
      </FormWrapper>

      {/* {created ?? (
        <Link href={`/event?id=${createdEventId}`}>募集したイベント</Link>
      )} */}
    </BasicTemplate>
  );
}
