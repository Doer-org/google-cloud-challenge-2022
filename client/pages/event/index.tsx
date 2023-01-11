import { useState } from 'react';
import { TypoWrapper } from '../../components/atoms/text/TypoWrapper';
import { EventBasicInfo } from '../../components/molecules/Event/EventBasicInfo';
import { EventWrapper } from '../../components/molecules/Event/EventWrapper';
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate';
import { MyHead } from '../../components/templates/shared/Head/MyHead';

export default function Index() {
  // TODO: 自分の作ったevent一覧をとってくるhooks使う
  const [events] = useState([
    { eventName: 'お好み焼きいこ', detail: 'お好み焼き行きたかった' },
    { eventName: 'ラーメンいこ', detail: 'ラーメン行きたかった' },
    { eventName: '居酒屋いこ', detail: '居酒屋行きたかった' },
  ]);
  return (
    <>
      <MyHead title="サービス名" description="サービスの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className="my-10">自分のイベント一覧</h1>
        </TypoWrapper>
        {events.map((event) => {
          return (
            <div key={event.eventName}>
              <EventWrapper>
                <EventBasicInfo
                  eventName={event.eventName}
                  detail={event.detail}
                />
              </EventWrapper>
            </div>
          );
        })}
      </BasicTemplate>
    </>
  );
}
