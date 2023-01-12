import { fptsHelper } from '../../utils/fptsHelper'
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as R from "fp-ts/lib/Record";
import * as A from "fp-ts/lib/Array";
import { EventApi } from '../../utils/gcChallengeApi'
import { Event } from '../../types/event'
  
 
export const tryGetEventInfo = (event_id: string) => { 
    const getEventInfo = flow(EventApi.getEvent, fptsHelper.TE.ofApiResponse)
    const getEventHost = flow(EventApi.getEventHost, fptsHelper.TE.ofApiResponse,)
    const getEventMembers = flow(EventApi.getEventMembers, fptsHelper.TE.ofApiResponse,) 
    const getComments = flow (EventApi.getEventComments, fptsHelper.TE.ofApiResponse)
    
    const getEventInfoAndComments = flow(
        getEventInfo,
        TE.chain((eventInfo) => pipe(
            getComments({id: eventInfo.id}),
            TE.map(flow(
                A.map((e) => {
                    const r : [string, string] = [e.id, e.body]
                    return r
                }),
                R.fromEntries, 
                (comment) => (
                    {
                        eventInfo : eventInfo,
                        commentDic : comment
                    }
                )
            )) 
        ))
    )

    const getHostAndMembers = flow(
        getEventHost,
        TE.chain((host) => pipe(
            getEventMembers({ id: event_id }),
            TE.map((members) => {
                return {
                    host: host,
                    members: members.filter((member) => member.id !== host.id)
                }
            })
        )))

    return pipe(
        getEventInfo({ id: event_id }),
        TE.chain((event_info) => pipe(
            getEventInfoAndComments({ id: event_info.id }), 
            TE.chain((info) => pipe(
                getHostAndMembers({ id: info.eventInfo.id }),
                TE.map((hm) => { 
                    const e: Event = {
                        event_id: info.eventInfo.id,
                        event_name: info.eventInfo.name,
                        detail: info.eventInfo.detail || "",
                        location: info.eventInfo.location || "",
                        host: { user_id: hm.host.id },
                        /**
                         * host comment
                         * R.has(hm.host.id,info.commentDic) 
                            ? info.commentDic[hm.host.id]
                            : ""
                         */
                        participants: 
                            hm.members.map((member) => ( 
                                {
                                    participant_name: member.name,
                                    comment: 
                                        R.has(member.id,info.commentDic) 
                                        ? info.commentDic[member.id]
                                        : ""
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
    (task) => task().then(() => { }),
    () => {}
)