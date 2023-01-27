import { useEffect, useState } from 'react';
import { TypoWrapper } from '../../components/atoms/text/TypoWrapper';
import { EventWrapper } from '../../components/molecules/Event/EventWrapper';
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate';
import { MyHead } from '../../components/templates/shared/Head/MyHead';
import { getEventList } from '../../core/api/user/getEventList';
import { useUserInfoStore } from '../../store/userStore';
import { components } from '../../core/openapi/openapi';
import { EventBasicInfoCard } from '../../components/molecules/Event/EventBasicInfoCard';
import { useRouter } from 'next/router';
import { LinkTo } from '../../components/atoms/text/LinkTo';
export default function Index() {
  const [events, setEvents] = useState<
    components['schemas']['User_EventsList'][]
  >([]);
  // TODO: SSR化する
  const router = useRouter();
  const { userInfo } = useUserInfoStore();
  useEffect(() => {
    const getEvents = getEventList(
      (response) => {
        setEvents(response);
      },
      (error) => {
        router.push('/');
      }
    );
    getEvents(userInfo.userId);
  }, []);

  return (
    <>
      <MyHead title="サービス名" description="サービスの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className="my-10">自分のイベント一覧</h1>
        </TypoWrapper>
        {events.map((event: components['schemas']['User_EventsList']) => {
          if (event.state === 'close') {
            return null;
          }
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
        <LinkTo href="/">トップに戻る</LinkTo>
      </BasicTemplate>
    </>
  );
}
