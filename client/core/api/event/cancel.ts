  
import {flow} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
 
export module CancelInputsExample {
    export const causeError = "-1"
}


const tryCancel = (event_id : string) : TE.TaskEither<Error,string> => {
    switch (event_id) {
        case (CancelInputsExample.causeError) : {
            return TE.left(Error("fail > tryCancel"))
        }
        default : {
            return TE.right("cancel!")
        }
    }  
}

/**
* 未実装，patch?
*/
export const cancelEvent = (
    okHandler : (res: any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCancel,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}).catch(()=>{})
)
 