import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { UserInfo } from '../organisms/User/UserInfo';
import { Event } from '../../core/types/event';
import { FC } from 'react';
type TProps = {
  event: Event;
};
export const EventInfo = ({ event }: TProps) => {
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
        <div className="flex items-end justify-center gap-5 mx-5 overflow-x-scroll md:w-5/6 md:mx-auto pt-5">
          <UserInfo name={event.host.user_name} image={event.host.icon} />
          {event.participants &&
            event.participants.map((participant) => {
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

        <EventBasicInfo
          eventName={event.event_name}
          detail={event.detail}
          limitTime={event.close_limit ? event.close_limit : ''}
        />
        {isJson(event.location) ? (
          <div className="lg:m-10 m-3">
            <Map
              lat={JSON.parse(event.location).lat}
              lng={JSON.parse(event.location).lng}
            />
          </div>
        ) : (
          <></>
        )}
      </EventWrapper>
    </>
  );
};
