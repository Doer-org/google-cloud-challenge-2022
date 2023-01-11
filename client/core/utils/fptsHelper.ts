import * as FPTE from 'fp-ts/TaskEither'
import {pipe} from 'fp-ts/lib/function'
import { ApiResponse } from 'openapi-typescript-fetch'
export module fptsHelper {  
    export module TE { 
        export const ofApiResponse= <T> (resp : Promise<ApiResponse<T>>) => { // : FPTE.TaskEither<Error,T> => { 
            return pipe( 
                FPTE.tryCatch(
                    () => resp,
                    (e : any) => Error(e)
                ),
                FPTE.chain((r) => 
                    (r.ok)
                    ? FPTE.right(r.data) 
                    : FPTE.left(Error(`response: ${r.status} : ${r.headers}`))
                )
                // { // bind
                //     if (r.ok) {
                //         return FPTE.right(r.data)
                //     } else {
                //         return FPTE.left(Error(JSON.stringify(r)))
                //     }
                // }) 
            ) 
        }  
    } 
}  