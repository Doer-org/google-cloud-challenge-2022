import { components, operations, paths } from '../openapi/openapi';
export type Host = {
  user_id: string;
};

export type Participant = {
  participant_name: string;
  comment: string;
};

export type Event = {
  event_id: string;
  event_name: string;
  detail: string;
  location: string;
  host: Host;
  participants: Participant[];
};

export type EventState = 'open' | 'close' | 'cancel';
