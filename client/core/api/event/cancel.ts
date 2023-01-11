  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { fptsHelper } from '../../utils/fptsHelper'
import { EventApi } from '../../utils/gcChallengeApi'
 
// TODO: axios
export const tryCancel = (event_id : string) => 
    EventApi.updateEventState({id: event_id},  "cancel")

// { 
//     return pipe ( 
//         EventApi.updateEventState({id: event_id},  "cancel"),
//         fptsHelper.TE.ofApiResponse, 
//     )
// }

/**
* 未実装，patch?
*/
export const cancelEvent = (
    okHandler : (ok: any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCancel,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}).catch(()=>{}),
    () => {}
)
 