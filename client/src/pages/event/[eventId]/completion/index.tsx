import { useEffect, useState } from 'react';
import Router, { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { useCopyToClipboard } from 'usehooks-ts';
import { Event } from '../../../../core/types/event';
import { tryGetEventInfo } from '../../../../core/api/event/getInfo';
import { pipe } from 'fp-ts/lib/function';
import * as TE from 'fp-ts/TaskEither';
import { useNoticeStore } from '../../../../store/noticeStore';
import { LinkTo } from '../../../../components/atoms/text/LinkTo';
export default function Participate(event: Event) {
  const [_, copy] = useCopyToClipboard();
  const [origin, setOrigin] = useState('');
  const eventId = useRouter().query.eventId;
  const { changeNotice } = useNoticeStore();

  useEffect(() => {
    setOrigin(window.location.origin);
  }, []);
  return (
    <>
      <MyHead title="イベントURLコピー" description="" />
      <BasicTemplate className="text-center py-5">
        <EventInfo
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
          hostImage={event.host.icon}
          hostName={event.host.user_name}
        />
        <Button
          className="flex m-auto my-5"
          border
          onClick={() => {
            copy(`${origin}/event/${eventId}/participate`);
            changeNotice({ type: 'Success', text: 'コピー完了しました' });
          }}
        >
          募集URLをコピー
        </Button>
        <div className="my-5">
          <LinkTo href={`${origin}/event`}>event一覧へ</LinkTo>
        </div>
      </BasicTemplate>
    </>
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
