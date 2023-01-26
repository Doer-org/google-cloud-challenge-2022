import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { Participant } from '../../core/types/event';
import { UserInfo } from '../organisms/User/UserInfo';

type TProps = {
  participants?: Participant[];
  eventName: string;
  detail: string;
  location: string;
  hostImage: string;
  hostName: string;
};
export const EventInfo = ({
  eventName,
  participants,
  detail,
  location,
  hostImage,
  hostName,
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
          <UserInfo name={hostName} image={hostImage} />
          {participants &&
            participants.map((participant) => {
              return (
                <UserInfo
                  key={participant.participant_name}
                  name={participant.participant_name}
                  comment={participant.comment}
                  image={participant.icon}
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
