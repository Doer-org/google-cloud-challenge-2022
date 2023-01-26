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
export default function Close(event: Event) {
  // ここは締め切るページ
  // TODO:event締切hooksをonClickへ
  const eventId = useRouter().query.eventId;
  const close = closeEvent(
    (ok) => {},
    (err) => {}
  );

  return (
    <>
      <MyHead title="イベント締切ページ" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
        />
        <Button
          className="flex m-auto my-5"
          onClick={() => close(eventId as string)}
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
