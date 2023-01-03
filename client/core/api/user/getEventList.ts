  
import {flow} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
 

const tryGetEventList = (user_id : number) => {
    return TE.right("events")
}


export const getEventList = (
    okHandler : (events: any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryGetEventList,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)
 