import { useState } from 'react';
import { createNewEvent } from '../../../core/api/event/create';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../../../components/atoms/text/TypoWrapper';
import { Input } from '../../../components/atoms/form/Input';
import { Textarea } from '../../../components/atoms/form/Textarea';
import { FormWrapper } from '../../../components/atoms/form/FormWrapper';
import { Button } from '../../../components/atoms/text/Button';
import { MapForm } from '../../../components/atoms/form/MapForm';
import { useUserInfoStore } from '../../../store/userStore';
import { useRouter } from 'next/router';
import { TMapPosition } from '../../../components/atoms/map/MapBasicInfo';

export default function New() {
  const router = useRouter();
  // TODO: 今はBad Requestが飛んできているのでuser_idとか用意しないといけないっぽい？それかtypeとかstateを埋めてないから？
  const createEvent = createNewEvent(
    (ok) => {
      router.push(ok.url);
    },
    (e) => {
      console.log(e);
    }
  );
  const { userInfo } = useUserInfoStore();
  const [name, setName] = useState('');
  const [capacity, setCapacity] = useState(1);
  const [detail, setDetail] = useState('');
  const [location, setLocation] = useState<null | TMapPosition>(null);

  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">募集する</h1>
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
        <MapForm location={location} setLocation={setLocation} />
        <Button
          className="flex m-auto my-5"
          onClick={() => {
            createEvent(
              { user_id: 'a88d4cba-6211-40ee-8a23-3f259d0166d5' },
              {
                event_name: name,
                max_member: capacity,
                detail: detail,
                location: JSON.stringify(location),
                timestamp: Date.now(),
              }
            );
          }}
        >
          募集する
        </Button>
      </FormWrapper>
    </BasicTemplate>
  );
}
