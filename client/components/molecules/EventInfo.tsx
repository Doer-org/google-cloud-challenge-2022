import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { UserIcon } from './User/UserIcon';
import { LoadScriptTemplate } from '../templates/shared/LoadScriptTemplate';
export const EventInfo = () => {
  return (
    <>
      <Hanging />
      <EventWrapper>
        <UserIcon userName="miso" />
        <EventBasicInfo
          eventName="ラーメン行こう"
          detail="同志社周りのラーメン行こう！！あくたがわとかが良さげ"
        />
        <div className="lg:m-10 m-3">
          <LoadScriptTemplate>
            <Map lat={35.6809591} lng={139.7673068} />
          </LoadScriptTemplate>
        </div>
      </EventWrapper>
    </>
  );
};
