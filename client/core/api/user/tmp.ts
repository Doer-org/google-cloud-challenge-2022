 
import { fptsHelper } from '../../uitls/fptsHelper'
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as E from 'fp-ts/Either'
import { UserApi }from '../../uitls/mockApi'

const getUser = flow (
    UserApi.findById,
    fptsHelper.TE.ofApiResponse
)

export const test = async () => { 
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

export const test2 = async () => { 
    console.log( UserApi.findById({ id: 123 }) )
    return
}  