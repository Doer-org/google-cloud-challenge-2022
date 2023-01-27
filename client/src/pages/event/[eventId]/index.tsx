import { MyHead } from '../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { EventInfo } from '../../../components/molecules/EventInfo';
import { tryGetEventInfo } from '../../../core/api/event/getInfo';
import { Event } from '../../../core/types/event';
import { pipe } from 'fp-ts/lib/function';
import * as TE from 'fp-ts/TaskEither';
export default function Show(event: Event) {
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        <EventInfo
          participants={event.participants}
          eventName={event.event_name}
          detail={event.detail}
          location={event.location}
          hostImage={event.host.icon}
          hostName={event.host.user_name}
          limitTime={event.close_limit}
        />
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
        return { props: { ...response } };
      }
    )
  )();
}
