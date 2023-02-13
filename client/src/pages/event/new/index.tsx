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
import { useEffect } from 'react';
import { useNoticeStore } from '../../../store/noticeStore';
import { LinkTo } from '../../../components/atoms/text/LinkTo';

export default function New() {
  const router = useRouter();
  const { userInfo } = useUserInfoStore();
  const { changeNotice } = useNoticeStore();
  const [name, setName] = useState('');
  const [capacity, setCapacity] = useState<number>(1);
  const [detail, setDetail] = useState('');
  const [location, setLocation] = useState<null | TMapPosition>(null);
  const [limit, setLimit] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [limitMaxTime, setLimitMaxTime] = useState('');
  const [limitMinTime, setLimitMinTime] = useState('');
  const LIMIT_DAY_FROM_TODAY = 1;

  const createEvent = createNewEvent(
    (ok) => {
      router.push(
        `${process.env.NEXT_PUBLIC_FRONT_URL}/event/${ok.created_event.event_id}/completion`
      );
      changeNotice({ type: 'Success', text: '作成に成功しました' });
    },
    (e) => {
      changeNotice({ type: 'Error', text: '作成に失敗しました' });
    }
  );

  useEffect(() => {
    const now = new Date();
    let tomorrow = new Date();
    tomorrow.setDate(now.getDate() + LIMIT_DAY_FROM_TODAY);
    tomorrow.setHours(23, 59, 59, 999);
    now.setHours(now.getHours() + 9);
    const minTime = now.toISOString().replace(/\..+/, '').slice(0, 16);
    const maxTime = tomorrow.toISOString().replace(/\..+/, '').slice(0, 16);
    setLimitMinTime(minTime);
    setLimitMaxTime(maxTime);
    console.log(minTime, maxTime);
  }, []);

  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">募集する</h1>
      </TypoWrapper>

      <FormWrapper
        onSubmit={() => {
          if (!isLoading) {
            createEvent(
              {
                user_id: userInfo.userId,
                user_name: 'atode', //FIXME : user storeに保存・取得
                icon: 'mada',
              },
              {
                event_name: name,
                max_member: Number(capacity),
                detail: detail,
                location: JSON.stringify(location),
                created_at: new Date(Date.now()),
                limit_time: new Date(limit), // FIXME: 締め切り時間設定
              }
            );
            setIsLoading(true);
          }
        }}
      >
        <Input
          type="text"
          label="イベント名"
          content={name}
          changeContent={setName}
          required={true}
          maxLength={50}
          minLength={1}
        />
        <Input
          type="number"
          label="募集人数(最大5名)"
          min={1}
          max={5}
          content={capacity}
          changeContent={setCapacity}
          required={true}
        />
        <Input
          type="datetime-local"
          label="締切"
          min={limitMinTime}
          max={limitMaxTime}
          content={limit}
          changeContent={setLimit}
          required={true}
        />
        <Textarea
          label="詳細"
          content={detail}
          changeContent={setDetail}
          minLength={0}
          maxLength={300}
          required={true}
        />
        <MapForm location={location} setLocation={setLocation} />
        <Button className="flex m-auto my-5" disable={isLoading}>
          募集する
        </Button>
      </FormWrapper>
      <LinkTo href="/" className="m-1 block my-5" borderNone>
        戻る
      </LinkTo>
    </BasicTemplate>
  );
}
