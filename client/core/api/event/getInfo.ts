import { fptsHelper } from '../../uitls/fptsHelper'
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import { EventApi } from '../../uitls/mockApi'
import { Event } from '../../types/event'

export module GetEventInputsExample {
    export const causeError = "-1"
} 

const tryGetEventInfo = (event_id: string) => {
    if (event_id === GetEventInputsExample.causeError) {
        return TE.left(Error("fail > tryGetEventInfo"))
    }
    const getEventInfo = flow(EventApi.getEvent, fptsHelper.TE.ofApiResponse)
    const getEventHost = flow(EventApi.getEventHost, fptsHelper.TE.ofApiResponse,)
    const getEventMembers = flow(EventApi.getEventMembers, fptsHelper.TE.ofApiResponse,) 
    const getHostAndMembers = flow(
        getEventHost,
        TE.chain((host) => pipe(
            getEventMembers({ id: event_id }),
            TE.map((members) => {
                return {
                    host: host,
                    members: members
                }
            })
        )))

    return pipe(
        getEventInfo({ id: event_id }),
        TE.chain((event_info) => pipe(
            getEventInfo({ id: event_info.id }), 
            TE.chain((info) => pipe(
                getHostAndMembers({ id: info.id }),
                TE.map((hm) => { 
                    const e: Event = {
                        event_id: info.id,
                        event_name: info.name,
                        detail: info.detail || "",
                        location: info.location || "",
                        host: { user_id: hm.host.id },
                        participants: hm.members.map((member) => (
                            {
                                participant_name: member.name,
                                commemt: ""
                            }
                        ))
                    }
                    return e 
                }) 
            ))
        ))
    )
}

export const getEventInfo = (
    okHandler: (event: Event) => void,
    errorHandler: (e: Error) => void,
) => flow(
    tryGetEventInfo,
    TE.match(
        errorHandler,
        okHandler
    ),
    (task) => task().then(() => { })
)