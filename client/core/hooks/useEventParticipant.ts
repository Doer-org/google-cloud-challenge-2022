 
import { fptsHelper } from '../uitls/fptsHelper'
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as T from 'fp-ts/Task'
import * as E from 'fp-ts/Either'
import { EventApi }from '../uitls/mockApi'
import { Event, Host } from '../types/event'
 

const tryJoinEvent = (current_event_state : Event) => ( 
    param : {
        event_id : number, 
        participant_name: string,
        comment : string,
    },
)  => { 
    const next_event_state = {
        ...current_event_state,
        participants : [ 
            {
                participant_name : param.participant_name,
                commemt : param.comment
            } ,
            ...current_event_state.participants
        ]
    }
    return TE.right(next_event_state) 
}

export default () => {  

    const joinEvent = (
        okHandler : (event :  Event) => void,
        errorHandler : (e: Error) => void,
    ) => (aaaabbbbbcccc : Event) => flow (
        tryJoinEvent(aaaabbbbbcccc),
        TE.match(
            errorHandler,
            okHandler
        )
    )
    return { 
        joinEvent
    }
}