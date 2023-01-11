import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'  
import { fptsHelper } from '../../utils/fptsHelper'
import { UserApi }from '../../utils/gcChallengeApi' 
import { components, operations, paths } from "../../openapi/openapi"
export const tryCreateUser = (user : components["schemas"]["UserCreate"]) => pipe (
    UserApi.createUser(user),
    fptsHelper.TE.ofApiResponse 
)  

/**
* 多分動く
*/
export const createUser = (
    okHandler : (a : any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCreateUser,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)