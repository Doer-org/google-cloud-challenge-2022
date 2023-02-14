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
        <div className="mx-auto md:w-2/3 flex justify-center">
          <div className="flex items-end gap-5 overflow-x-scroll pt-5 mx-5">
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
