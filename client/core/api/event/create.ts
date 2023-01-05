import { fptsHelper } from '../../uitls/fptsHelper'
import {flow, pipe} from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither' 
import { EventApi }from '../../uitls/mockApi' 
import { Event, Host } from '../../types/event'

const tryCreateNewEvent = (
    host : Host,
    param : {
        event_name: string,
        max_member : number,
        detail : string,
        location : string
    } 
) : TE.TaskEither<Error, {url:string, created_event : Event}> => { 
    return pipe (
        {
            name : param.event_name,
            detail : param.detail,
            location : param.location 
        },
        EventApi.create,
        fptsHelper.TE.ofApiResponse,
        TE.map((res) => {
            const e : Event =  {
                event_id : res.id,
                event_name : res.name,
                detail : res.detail,
                location : res.location,
                host : host,
                participants : []
            }
            return {
                url : `http://localhost/event?id=${res.id}`,
                created_event : e
            } 
        }) 
    )
}

export const createNewEvent = (
    okHandler : (ok: {url:string, created_event : Event}) => void,
    errorHandler : (e: Error) => void,
) => flow (
    tryCreateNewEvent,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => {})
)