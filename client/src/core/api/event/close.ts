  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import { fptsHelper } from '../../utils/fptsHelper'
import { EventApi } from '../../utils/gcChallengeApi' 
 
export const tryCloseEvent = (event_id : string)   => 
     EventApi.updateEventState({id: event_id},  "close")
// { 
//     return pipe (
//         {
//             // id : param.event_id,
//             // state: "close"
//         },
//         EventApi.updateEventState,
//         fptsHelper.TE.ofApiResponse, 
//     )
// }
 

/**
* 未実装，patch?
*/
export const closeEvent = (
    okHandler : (ok : any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCloseEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)
 