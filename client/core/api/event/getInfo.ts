import { fptsHelper } from '../../uitls/fptsHelper'
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { EventApi }from '../../uitls/mockApi'  
import { Event, Host, Participant } from '../../types/event' 

const tryGetEventInfo = (event_id: number) => { 
    const getEventInfo = flow (
        EventApi.findById,
        fptsHelper.TE.ofApiResponse, 
    )
    
    const getEventHost = flow (
        EventApi.getEventMembers,
        fptsHelper.TE.ofApiResponse, 
    ) 
    return pipe ( 
        getEventInfo({id : event_id}), 
        TE.chain((event_info) => pipe (
            getEventHost({id : event_info.id}),
            TE.map((members) => {
                const e : Event =  {
                    event_id : event_info.id,
                    event_name : event_info.name,
                    detail : event_info.detail,
                    location : event_info.location,
                    host : {user_id: ""},
                    participants : members.map((member) => ( 
                        {
                            participant_name : member.name,
                            commemt : ""
                        }
                    ))
                } 
                return e 
            })
        )) 
    )
}

export const getEventInfo = (
    okHandler : (event :  Event) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryGetEventInfo,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)