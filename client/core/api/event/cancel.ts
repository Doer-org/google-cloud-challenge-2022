  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { fptsHelper } from '../../uitls/fptsHelper'
import { EventApi } from '../../uitls/mockApi'
 

const tryCancel = (event_id : string)   => { 
    return pipe (
        {
            // id : param.event_id,
            // state: "cancel"
        },
        EventApi.updateEventState,
        fptsHelper.TE.ofApiResponse, 
    )
}

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
 