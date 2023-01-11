  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import { fptsHelper } from '../../utils/fptsHelper'
import { EventApi } from '../../utils/gcChallengeApi' 
import { components, operations, paths } from "../../openapi/openapi"
 
  
export const tryUpdateEvent = (param : components["schemas"]["EventUpdate"])   => { 
    return pipe (
        param,
        EventApi.updateEvent,
        fptsHelper.TE.ofApiResponse, 
    )
}
 

/**
* 未実装，patch?
*/
export const updateEvent = (
    okHandler : (ok : any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryUpdateEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)
 