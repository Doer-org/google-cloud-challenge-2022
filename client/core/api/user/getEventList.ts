  
import {flow} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { Event } from '../../types/event'
 

export module GetEventListInputsExample {
    export const causeError = -1
}

const events : Event[]  =  [ 
    {
        event_id: "10",
        event_name: "a" ,
        detail: "a" ,
        location: "a" ,
        host: {
            user_id : "1"
        } ,
        participants: [] 
    },
    {
        event_id: "20",
        event_name: "b" ,
        detail: "b" ,
        location: "b" ,
        host: {
            user_id : "1"
        } ,
        participants: [] 
    }
]
 


const tryGetEventList = (user_id : number) : TE.TaskEither<Error, Event[]> => {
    if(user_id===GetEventListInputsExample.causeError) {
        return TE.left(Error("tryGetEventList > "))
    }
    return TE.right(events)
}


export const getEventList = (
    okHandler : (events: Event[]) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryGetEventList,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)
 