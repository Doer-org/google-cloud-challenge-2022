import { fptsHelper } from '../../utils/fptsHelper'
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { EventApi }from '../../utils/gcChallengeApi' 
import { Event, Host } from '../../types/event'

export const tryCreateNewEvent = (
    host : Host,
    param : {
        event_name: string,
        max_member : number,
        detail : string,
        location : string,
        created_at : Date,
        limit_time : Date
    } 
) : TE.TaskEither<Error, {url:string, created_event : Event}> => { 
    return pipe ( 
        
        EventApi.createEvent({
            name: param.event_name,
            admin: host.user_id, 
            detail: param.detail,
            location: param.location,
            type:  "??",
            state:  "open",
            size: param.max_member,
            created_at: param.created_at.toISOString(),
            limit_time: param.limit_time.toISOString()
        }, 
        {
          credentials: 'include',
        },),
        fptsHelper.TE.ofApiResponse,
        TE.map((res) => {
            const e : Event =  {
                event_id : res.id,
                event_name : res.name,
                event_size : res.size,
                event_state : res.state,
                detail : res.detail || "",
                location : res.location || "",
                host : host,
                participants : []
            }
            return {
                url : `http://localhost/event/${res.id}/participate`,
                created_event : e
            } 
        }) 
    )
}

/**
* 引数type, stateに，何を与えるといいか分からない．
*/
export const createNewEvent = (
    okHandler : (ok: {url:string, created_event : Event}) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCreateNewEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)