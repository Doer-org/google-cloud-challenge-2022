 
import { fptsHelper } from '../uitls/fptsHelper'
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as T from 'fp-ts/Task'
import * as E from 'fp-ts/Either'
import { UserApi }from '../uitls/mockApi'
export default () => { 
    const getUser = flow (
        UserApi.findById,
        fptsHelper.TE.ofApiResponse
    )

    const test = async () => { 
        const r = await pipe(
            getUser({ id: 123 }),
            TE.match(
                (e) => E.left(e),
                (ok) => E.right(ok)
            ), 
        )()
        console.log(r)
        return
    } 
    const test2 = async () => { 
        console.log( UserApi.findById({ id: 123 }) )
        return
    } 

    return {
        test
    }
}