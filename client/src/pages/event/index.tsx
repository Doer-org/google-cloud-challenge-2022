import { useEffect, useState } from 'react';
import { useQuery } from 'react-query';
import { TypoWrapper } from '../../components/atoms/text/TypoWrapper';
import { EventBasicInfo } from '../../components/molecules/Event/EventBasicInfo';
import { EventWrapper } from '../../components/molecules/Event/EventWrapper';
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate';
import { MyHead } from '../../components/templates/shared/Head/MyHead';
import { getEventList } from '../../core/api/user/getEventList';
import { useUserInfoStore } from '../../store/userStore';
import { Event } from '../../core/types/event';
import { components } from '../../core/openapi/openapi';
import { EventBasicInfoCard } from '../../components/molecules/Event/EventBasicInfoCard';
export default function Index() {
  // TODO: 自分の作ったevent一覧をとってくるhooks使う

  const [events, setEvents] = useState<
    components['schemas']['User_EventsList'][]
  >([]);
  // const { userId } = UserStore();
  // TODO: SSR化する
  const { userInfo, setUserInfo } = useUserInfoStore();
  useEffect(() => {
    const getEvents = getEventList(
      (response) => {
        setEvents(response);
      },
      (error) => {}
    );
    getEvents(userInfo.userId);
  }, [userInfo]);

  return (
    <>
      <MyHead title="サービス名" description="サービスの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className="my-10">自分のイベント一覧</h1>
          {/* TODO: stateがcloseじゃなかったらreturn */}
        </TypoWrapper>
        {events.map((event: components['schemas']['User_EventsList']) => {
          return (
            <div key={event.name}>
              <EventWrapper>
                <EventBasicInfoCard
                  id={event.id}
                  eventName={event.name}
                  detail={event.detail as string}
                />
              </EventWrapper>
            </div>
          );
        })}
      </BasicTemplate>
    </>
  );
}
