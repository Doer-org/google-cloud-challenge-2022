  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import { fptsHelper } from '../../uitls/fptsHelper'
import { EventApi } from '../../uitls/mockApi' 
import { components, operations, paths } from "../../openapi/openapi"
 
  
const tryUpdateEvent = (param : components["schemas"]["EventUpdate"])   => { 
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
 