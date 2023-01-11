import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'  
import { fptsHelper } from '../../utils/fptsHelper'
import { UserApi }from '../../utils/gcChallengeApi' 
import { components, operations, paths } from "../../openapi/openapi"
export const tryUpdateUser = (user : components["schemas"]["UserUpdate"]) => pipe (
    UserApi.createUser(user),
    fptsHelper.TE.ofApiResponse 
)  

/**
* 多分動く
*/
export const updateUser = (
    okHandler : (a : any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryUpdateUser,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)