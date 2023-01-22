import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { Participant } from '../../core/types/event';
import { TypoWrapper } from '../atoms/text/TypoWrapper';
import { UserInfo } from '../organisms/User/UserInfo';

type TProps = {
  participants?: Participant[];
  eventName: string;
  detail: string;
  location: string;
};
export const EventInfo = ({
  eventName,
  participants,
  detail,
  location,
}: TProps) => {
  const isJson = (location: string) => {
    try {
      JSON.parse(location).lat;
      JSON.parse(location).lng;
      return true;
    } catch (error) {
      return false;
    }
  };
  return (
    <>
      <Hanging />
      <EventWrapper>
        <div className="flex items-end gap-5 mx-5 overflow-x-scroll md:w-5/6 md:mx-auto py-5">
          <UserInfo name={'ホストです'} />
          {participants &&
            participants.map((participate) => {
              return (
                <UserInfo
                  key={participate.participant_name}
                  name={participate.participant_name}
                  comment={participate.comment}
                  isParticipate
                />
              );
            })}
        </div>

        <EventBasicInfo eventName={eventName} detail={detail} />
        {isJson(location) ? (
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
