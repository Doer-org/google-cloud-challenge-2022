  
import {flow} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
 

const tryCancel = (event_id : number) => {
    return TE.right("cancel!")
}


export const cancelEvent = (
    okHandler : (res: any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCancel,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)
 