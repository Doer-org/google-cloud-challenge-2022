import { flow } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { Event } from '../../types/event'
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
) : TE.TaskEither<Error,Event> => { 
    if (param.event_id === JoinInputExample.causeError.event_id) {
        return TE.left(Error("tryJoinEvent > fail"))
    }
    const next_event_state : Event = JoinRetExample.ok
    return TE.right(next_event_state) 
}
export  const joinEvent = (
    okHandler : (event :  Event) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryJoinEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)

export default () => {  

    return { 
        joinEvent
    }
}