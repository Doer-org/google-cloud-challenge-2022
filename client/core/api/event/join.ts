import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { Event } from '../../types/event'
import { fptsHelper } from '../../utils/fptsHelper'
import { EventApi } from '../../utils/gcChallengeApi'


export const tryJoinEvent = ( 
    param : {
        event_id : string, 
        participant_name: string,
        comment : string,
    },
)  => 
    EventApi.join(
        param.event_id,{
            name : param.participant_name,
            comment : param.comment
        }
    )

// {  
//     return pipe (
//         // {
//         //     // id : param.event_id,
//         //     // name : param.participant_name,
//         //     // comment : param.cooment
//         // },
//         EventApi.join(
//             param.event_id,{
//                 name : param.participant_name,
//                 comment : param.comment
//             }
//         ),
//         fptsHelper.TE.ofApiResponse, 
//     )
// }
/**
* 未実装，patch?
*/
export  const joinEvent = (
    okHandler : (success :  unknown) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryJoinEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {}),
    () => {}
) 