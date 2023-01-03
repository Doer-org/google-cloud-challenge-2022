  
import {flow} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
 

const tryCloseEvent = (event_id : number) => {
    return TE.right(event_id)
}


export const closeEvent = (
    okHandler : (event_id: number) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCloseEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)
 