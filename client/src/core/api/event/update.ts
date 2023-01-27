  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import { fptsHelper } from '../../utils/fptsHelper'
import { EventApi } from '../../utils/gcChallengeApi' 
import { components, operations, paths } from "../../openapi/openapi"
 
  
export const tryUpdateEvent = (
    param : {
        id : string // eventID
        name: string,
        size : number,
        detail : string,
        location : string,
        created_at : Date,
        limit_time : Date,
        type ?: string,
        state ?: string
    } 

)   => { 
    return pipe (
        EventApi.updateEvent(
            {
                name: param.name,
                detail: param.detail,
                location: param.location,
                size: param.size,
                created_at: param.created_at.toISOString(),
                limit_time: param.limit_time.toISOString(),
                id: param.id
            }, 
            {
                credentials: 'include',
            }
        ),
        fptsHelper.TE.ofApiResponse, 
    )
}
 

/**
* 未実装，patch?
*/
export const updateEvent = (
    okHandler : (ok : any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryUpdateEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)
 