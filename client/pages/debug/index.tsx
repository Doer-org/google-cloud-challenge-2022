import { useState } from 'react';
import { Button } from '../../components/atoms/text/Button';
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate';
import {
  createNewEvent,
  CreateParamExample,
} from '../../core/api/event/create';
export default function New() {
  const [createdEventId, setCreatedEventId] = useState("1");
  const createEvent = createNewEvent(
    (v) => {
      console.log(v);
      setCreatedEventId(v.created_event.event_id);
    },
    (v) => {
      setCreatedEventId("1");
      console.log(v);
    }
  );

  return (
    <BasicTemplate className="text-center">
      eventid : {createdEventId}
      <Button
        className="flex m-auto my-5"
        onClick={
          () => { createEvent({ user_id: '' }, CreateParamExample.causeError)}
        }
      >
        fail
      </Button>
      <Button
        className="flex m-auto my-5"
        onClick={() => {
          createEvent(
            { user_id: 'abc' },
            { ...CreateParamExample.causeError, max_member: 1 }
          );
        }}
      >
        success
      </Button>
    </BasicTemplate>
  );
}
