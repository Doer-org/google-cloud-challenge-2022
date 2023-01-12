import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { UserIcon } from './User/UserIcon';
import { Participant } from '../../core/types/event';
import { ParticipateInfo } from './User/ParticipateInfo';
import { TypoWrapper } from '../atoms/text/TypoWrapper';
type TProps = {
  userId?: string;
  participants?: Participant[];
  eventName: string;
  detail: string;
  location: string;
};
export const EventInfo = ({
  userId,
  eventName,
  participants,
  detail,
  location,
}: TProps) => {
  return (
    <>
      <Hanging />
      <EventWrapper>
        <div className="grid md:grid-cols-6 grid-cols-4 items-end">
          <div>
            <TypoWrapper line="bold">ホスト</TypoWrapper>
            <UserIcon userName="miso" />
          </div>

          {participants ? (
            <ParticipateInfo participants={participants} />
          ) : (
            <></>
          )}
        </div>

        <EventBasicInfo eventName={eventName} detail={detail} />
        {location ? (
          <div className="lg:m-10 m-3">
            <Map
              lat={JSON.parse(location).lat}
              lng={JSON.parse(location).lng}
            />
          </div>
        ) : (
          <></>
        )}
      </EventWrapper>
    </>
  );
};
