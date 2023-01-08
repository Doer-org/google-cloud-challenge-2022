import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'  
import { fptsHelper } from '../../uitls/fptsHelper'
import { UserApi }from '../../uitls/mockApi' 
import { components, operations, paths } from "../../openapi/openapi"
const tryGetUserInfo = (user : string ) => pipe (
    UserApi.getUser({id:user}),
    fptsHelper.TE.ofApiResponse 
)  

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