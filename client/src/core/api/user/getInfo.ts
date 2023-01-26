import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'  
import { fptsHelper } from '../../utils/fptsHelper'
import { UserApi }from '../../utils/gcChallengeApi' 
import { components, operations, paths } from "../../openapi/openapi"
export const tryGetUserInfo = (user_id : string ) => pipe (
    UserApi.getUser({id:user_id},{credentials: 'include'}),
    fptsHelper.TE.ofApiResponse 
)  

/**
* 多分動く
*/
export const getUserInfo = (
    okHandler : (a : any) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryGetUserInfo,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
)