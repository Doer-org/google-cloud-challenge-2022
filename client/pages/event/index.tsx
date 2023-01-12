import { useEffect, useState } from 'react';
import { useQuery } from 'react-query';
import { TypoWrapper } from '../../components/atoms/text/TypoWrapper';
import { EventBasicInfo } from '../../components/molecules/Event/EventBasicInfo';
import { EventWrapper } from '../../components/molecules/Event/EventWrapper';
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate';
import { MyHead } from '../../components/templates/shared/Head/MyHead';
import { getEventList } from '../../core/api/user/getEventList';
import { useUserInfoStore } from '../../store/userStore';

export default function Index() {
  // TODO: 自分の作ったevent一覧をとってくるhooks使う

  const [events, setEvents] = useState<any>([]);

  useEffect(() => {
    const getEvents = getEventList(
      (response) => {
        setEvents(response);
      },
      (error) => {
        console.log(error);
      }
    );
    getEvents('a88d4cba-6211-40ee-8a23-3f259d0166d5');
  }, []);

  return (
    <>
      <MyHead title="サービス名" description="サービスの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className="my-10">自分のイベント一覧</h1>
        </TypoWrapper>
        {events.map((event: any) => {
          return (
            <div key={event.name}>
              <EventWrapper>
                <EventBasicInfo eventName={event.name} detail={event.detail} />
              </EventWrapper>
            </div>
          );
        })}
      </BasicTemplate>
    </>
  );
}
