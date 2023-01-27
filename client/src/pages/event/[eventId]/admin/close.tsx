import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { Event } from '../../../../core/types/event';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { closeEvent } from '../../../../core/api/event/close';
import {
  getEventInfo,
  tryGetEventInfo,
} from '../../../../core/api/event/getInfo';
import { pipe } from 'fp-ts/lib/function';
import * as TE from 'fp-ts/TaskEither';
import { useNoticeStore } from '../../../../store/noticeStore';
export default function Close(event: Event) {
  // ここは締め切るページ
  // TODO:event締切hooksをonClickへ
  const { changeNotice } = useNoticeStore();
  const router = useRouter();
  const close = closeEvent(
    (ok) => {
      router.push('/event');
      changeNotice({ type: 'Success', text: '締切ました' });
    },
    (err) => {
      changeNotice({ type: 'Error', text: '締切に失敗しました' });
    }
  );
  return (
    <>
      <MyHead title="イベント締切ページ" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
          hostImage={event.host.icon}
          hostName={event.host.user_name}
        />
        <Button
          className="flex m-auto my-5"
          onClick={() => close(event.event_id)}
        >
          締め切る
        </Button>
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
