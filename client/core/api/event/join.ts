import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { Event } from '../../types/event'
import { fptsHelper } from '../../uitls/fptsHelper'
import { EventApi } from '../../uitls/mockApi'

export module JoinRetExample { 
    export const ok : Event = {
        event_id: "10",
        event_name: "string" ,
        detail: "string" ,
        location: "string" ,
        host: {
            user_id : "1"
        } ,
        participants: [] 
    }
}

export module JoinInputExample {
    export const causeError : Event = {
        event_id: "-1",
        event_name: "string" ,
        detail: "string" ,
        location: "string" ,
        host: {
            user_id : "1"
        } ,
        participants: [] 
    }
    export const ok : Event = {
        event_id: "10",
        event_name: "string" ,
        detail: "string" ,
        location: "string" ,
        host: {
            user_id : "1"
        } ,
        participants: [] 
    }
}

const tryJoinEvent = ( 
    param : {
        event_id : string, 
        participant_name: string,
        comment : string,
    },
)  => {  
    return pipe (
        {
            // id : param.event_id,
            // name : param.participant_name,
            // comment : param.cooment
        },
        EventApi.join,
        fptsHelper.TE.ofApiResponse, 
    )
    // const next_event_state : Event = JoinRetExample.ok
    // return TE.right(next_event_state) 
}
/**
* 未実装，patch?
*/
export  const joinEvent = (
    okHandler : (success :  unknown) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryJoinEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
) 