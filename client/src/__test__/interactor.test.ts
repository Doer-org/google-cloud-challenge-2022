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
import { right } from 'fp-ts/lib/EitherT'
  

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
    const event = await createNewEvent()()
    const eventid = event._tag === "Right" ? event.right.id : "" 
    await tryJoinEvent({
      event_id : eventid,
      participant_name : "A",
      comment : "aaaa"
    })()
    
    await tryJoinEvent({
      event_id : eventid,
      participant_name : "B",
      comment : "bbbb"
    })() 
    await tryJoinEvent({
      event_id : eventid,
      participant_name : "C",
      comment : "cccc"
    })()
         
    const ret =  await tryGetEventInfo(eventid)()
    console.log(ret)
    if(ret._tag === "Right") {
      ret.right.participants.map((v, i) => {
        console.log(v)
      })
    }
    expect(  
      E.isRight(event) && E.isRight(ret)
    )
    .toBe(
      true 
    ) 
  })
})
