import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { UserIcon } from './User/UserIcon';
import { useEffect, useState } from 'react';
type TProps = {
  eventName: string;
  detail: string;
  location: string;
};
export const EventInfo = ({ eventName, detail, location }: TProps) => {
  return (
    <>
      <Hanging />
      <EventWrapper>
        <UserIcon userName="miso" />
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
