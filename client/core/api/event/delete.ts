import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'  
import { fptsHelper } from '../../utils/fptsHelper'
import { EventApi }from '../../utils/gcChallengeApi' 
  
export const tryDeleteEvent = (id : string) => pipe (
    EventApi.deleteEvent({id: id}),
    fptsHelper.TE.ofApiResponse 
)  

/**
* 多分動く
*/
export const deleteEvent = (
    okHandler : (a : unknown) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryDeleteEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)
 