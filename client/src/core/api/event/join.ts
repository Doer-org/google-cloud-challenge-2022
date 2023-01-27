import { flow, pipe } from 'fp-ts/lib/function';
import * as TE from 'fp-ts/TaskEither';
import { Event } from '../../types/event';
import { fptsHelper } from '../../utils/fptsHelper';
import { EventApi } from '../../utils/gcChallengeApi';

export const tryJoinEvent = (param: {
  event_id: string;
  participant_name: string;
  comment: string;
}) => {
  console.log(param)
  return pipe (
    EventApi.join(param.event_id, {
      name: param.participant_name,
      comment: param.comment,
    }),
    TE.chain((r) => {  
      const body = 
        TE.tryCatch(
          () => r.json(),
          (e : any) => Error(e)
        ) 
      return pipe(
        body,
        TE.chain((resp200) =>  
            (resp200.code !== 400) 
            ? TE.right(r)
            : TE.left(Error(`response: ${r}`))  
        )
      )
    }),
  )
}

export const joinEvent = (
  okHandler: (success: unknown) => void,
  errorHandler: (e: Error) => void
) =>
  flow(
    tryJoinEvent,
    TE.match(errorHandler, okHandler),
    (task) => task().then(() => {}),
    () => {}
  );
