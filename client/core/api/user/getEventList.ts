  
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { Event } from '../../types/event'
import { fptsHelper } from '../../utils/fptsHelper'
import { UserApi } from '../../utils/gcChallengeApi'
import { components, operations, paths } from "../../openapi/openapi"
 

export const tryGetEventList = (user_id : string)  : TE.TaskEither<Error, components["schemas"]["User_EventsList"][]>  => pipe (
    UserApi.getUsersEvents({id:user_id}),
    fptsHelper.TE.ofApiResponse 
)  

/**
* 多分動く
*/
export const getEventList = (
    okHandler : (events : components["schemas"]["User_EventsList"][]) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryGetEventList,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)
 