import { describe, it, expect, test } from 'vitest'
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as E from 'fp-ts/Either'
import { EventApi, UserApi } from '../../core/utils/gcChallengeApi'
import { fptsHelper } from '../../core/utils/fptsHelper'

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


describe('[gccAPI] EventApi.createEvent', () => {
  it('正常系：Eventを作成', async () => {
    const resp = await createNewEvent()()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})


describe('[gccAPI] EventApi.getEvent', () => {
  it('正常系：Eventを取得', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) => pipe(
        EventApi.getEvent({
          id: event.id
        }),
        fptsHelper.TE.ofApiResponse
      )
      )
    )()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})


describe('[gccAPI] EventApi.deleteEvent', () => {
  it('正常系：Eventを削除', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) =>
        EventApi.deleteEvent({
          id: event.id
        })
      )
    )()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})

describe('[gccAPI] EventApi.updateEvent', () => {
  it('正常系：Eventを更新', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) => pipe(
        EventApi.updateEvent({
          id: event.id,
          name: `updated ${event.name}`
        }
        ),
        fptsHelper.TE.ofApiResponse,
      )
      )
    )()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})


describe('[gccAPI] EventApi.getEventHost', () => {
  it('正常系：Eventの主催者を取得', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) => pipe(
        EventApi.getEventHost({
          id: event.id,
        }
        ),
        fptsHelper.TE.ofApiResponse,
      )
      )
    )()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})

describe('[gccAPI] EventApi.getEventComments', () => {
  it('正常系：Eventのコメント取得', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) => pipe(
        EventApi.getEventComments({
          id: event.id,
        }
        ),
        fptsHelper.TE.ofApiResponse,
      )
      )
    )()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})

// // No response in the range 200-299 defined

describe('[gccAPI] EventApi.join', () => {
  it('正常系：Eventに参加', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) =>
        EventApi.join(
          event.id,
          {
            name: "aoki v2",
            comment: "doer saikooooo"
          }
        )
      )
    )()
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
}) 

// describe('[gccAPI] EventApi.updateEventState', () => {
//   it('正常系：Eventのstate変更', async () => { 
//     const resp = await pipe(
//       createNewEvent(), 
//       TE.chain((event)=>  
//         EventApi.updateEventState({ 
//             id: event.id, 
//           }
//           , "close"
//         )
//       ) 
//     )()    
//     console.log(resp)
//     expect( 
//         E.isRight(resp)
//     )
//     .toBe(
//         true 
//     ) 
//   })
// })


describe('[gccAPI] EventApi.getEventComments', () => {
  it('正常系：Eventの参加者を取得（ホストを含む）', async () => {
    const resp = await pipe(
      createNewEvent(),
      TE.chain((event) => pipe(
        EventApi.getEventMembers({
          id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
        }
        ),
        fptsHelper.TE.ofApiResponse,
      )
      )
    )() 
    expect(
      E.isRight(resp)
    ).toBe(
      true
    )
  })
})








