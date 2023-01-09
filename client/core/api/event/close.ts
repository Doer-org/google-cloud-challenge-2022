  
import {flow} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
 

export module CloseInputsExample {
    export const causeError = "-1"
}

const tryCloseEvent = (event_id : string) : TE.TaskEither<Error,string> => {
    switch (event_id) {
        case (CloseInputsExample.causeError) : {
            return TE.left(Error("fail > tryCloseEvent"))
        }
        default : {
            return TE.right(event_id)
        }
    }  
}
 

/**
* 未実装，patch?
*/
export const closeEvent = (
    okHandler : (event_id: string) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCloseEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)
 