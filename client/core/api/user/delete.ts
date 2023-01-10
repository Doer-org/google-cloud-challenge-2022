import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'  
import { fptsHelper } from '../../uitls/fptsHelper'
import { UserApi }from '../../uitls/mockApi'  
const tryCreateUser = (user : string) => pipe (
    UserApi.deleteById({id:user}),
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