import { describe, it, expect, test } from 'vitest' 
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as E from 'fp-ts/Either'
import { EventApi, UserApi } from '../core/utils/gcChallengeApi'  
import { fptsHelper } from '../core/utils/fptsHelper'
import { tryCreateNewEvent } from '../core/api/event/create'
import { m } from 'vitest/dist/index-2d10c3fd'
import { getEventInfo, tryGetEventInfo } from '../core/api/event/getInfo'
import { joinEvent, tryJoinEvent } from '../core/api/event/join'
  

const createUser = () => pipe(
  UserApi.createUser({
    name: "aoki",
    authenticated: true,
    icon: "string"
  }),
  fptsHelper.TE.ofApiResponse,
) 

const createNewEvent = () => {
  const user = createUser()
  const resp = pipe(
    user,
    TE.chain((user) => pipe(
      EventApi.createEvent({
        name: "aoki camp",
        type: "string",
        state: "string",
        admin: user.id
      }),
      fptsHelper.TE.ofApiResponse,
    ))
  )
  return resp
}


describe('[core/api/create] create', () => {
  it('正常系：Eventを作成', async () => { 
    const hostUser = createUser()
    const ret = await pipe( 
      hostUser,
      TE.chain((user) =>  
        tryCreateNewEvent(
          { user_id : user.id },
          {
            event_name : "aoki camp!",
            max_member: 4,
            detail: "detail detail detail",
            location: "earth",
            timestamp: Date.now(), 
          }
        ) 
      ), 
    )()

    expect(  
      E.isRight(ret)
    )
    .toBe(
      true 
    ) 
  })
})


describe('[core/api/create] getInfo', () => {
  it('正常系：Event関連情報を取得', async () => { 
    const event = createNewEvent() 
    const ret = await pipe( 
      event, 
      TE.chain((event) => {
         
        tryJoinEvent({
          event_id : event.id,
          participant_name : "A",
          comment : "aaaa"
        })
        
        tryJoinEvent({
          event_id : event.id,
          participant_name : "B",
          comment : "bbbb"
        })
        
        tryJoinEvent({
          event_id : event.id,
          participant_name : "C",
          comment : "cccc"
        })
 
        return tryGetEventInfo(event.id)
      }), 
    )()
    expect(  
      E.isRight(ret)
    )
    .toBe(
      true 
    ) 
  })
})
