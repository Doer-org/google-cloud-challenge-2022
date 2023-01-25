import { describe, it, expect, test } from 'vitest' 
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as E from 'fp-ts/Either'
import { EventApi } from '../../core/utils/swaggerApi'  
import { fptsHelper } from '../../core/utils/fptsHelper'
 
describe('[swaggerApi] EventApi.createEvent', () => {
  it('正常系：Eventを作成', async () => {  
    const resp = await pipe( 
        EventApi.createEvent({
          name: "string",
          type: "string",
          state: "string",
          size: 2
        }),
        fptsHelper.TE.ofApiResponse, 
    )()
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
})


describe('[swaggerApi] EventApi.getEvent', () => {
  it('正常系：Eventを取得', async () => {  
    const resp = await pipe( 
        EventApi.getEvent({ 
            id: "3fa85f64-5717-4562-b3fc-2c963f66afa6"
        }),
        fptsHelper.TE.ofApiResponse, 
    )()
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
})


describe('[swaggerApi] EventApi.deleteEvent', () => {
  it('正常系：Eventを取得', async () => {  
    const resp = await EventApi.deleteEvent({ 
        id: "3fa85f64-5717-4562-b3fc-2c963f66afa6"
    } )()  
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
})


describe('[swaggerApi] EventApi.deleteEvent', () => {
  it('異常系：Eventを取得(idがGuidでない場合)', async () => {  
    const resp = await EventApi.deleteEvent({ 
        id: "abc"
    } )()  
    expect( 
        E.isLeft(resp)
    )
    .toBe(
        true 
    ) 
  })
}) 

describe('[swaggerApi] EventApi.updateEvent', () => {
  it('正常系：Eventを更新', async () => {  
    const resp = await pipe( 
      EventApi.updateEvent({ 
          id: "3fa85f64-5717-4562-b3fc-2c963f66afa6", 
          name: "updated name"
        }
      ),
      fptsHelper.TE.ofApiResponse, 
    )()  
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
})


describe('[swaggerApi] EventApi.getEventHost', () => {
  it('正常系：Eventの主催者を取得', async () => {  
    const resp = await pipe( 
      EventApi.getEventHost({ 
          id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",  
        }
      ),
      fptsHelper.TE.ofApiResponse, 
    )()  
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
}) 

describe('[swaggerApi] EventApi.getEventComments', () => {
  it('正常系：Eventのコメント取得', async () => {  
    const resp = await pipe( 
      EventApi.getEventComments({ 
          id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",  
        }
      ),
      fptsHelper.TE.ofApiResponse, 
    )()  
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
})
 
// No response in the range 200-299 defined

// describe('[swaggerApi] EventApi.join', () => {
//   it('正常系：Eventに参加', async () => {  
//     const resp = await EventApi.join(
//       "123",//"3fa85f64-5717-4562-b3fc-2c963f66afa6",
//       { 
//         name: "aoki",
//         comment : "viiiiiiitest"
//       }
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

// No response in the range 200-299 defined

// describe('[swaggerApi] EventApi.updateEventState', () => {
//   it('正常系：Eventのstate変更', async () => {  
//     const resp = await 
//       EventApi.updateEventState({ 
//           id: 111, //"3fa85f64-5717-4562-b3fc-2c963f66afa6",  
//         }
//         , "close"
//       )()  
//     console.log(resp)
//     expect( 
//         E.isRight(resp)
//     )
//     .toBe(
//         true 
//     ) 
//   })
// })

 
describe('[swaggerApi] EventApi.getEventComments', () => {
  it('正常系：Eventの参加者を取得（ホストを含む）', async () => {  
    const resp = await pipe( 
      EventApi.getEventMembers({ 
          id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",  
        }
      ),
      fptsHelper.TE.ofApiResponse, 
    )()  
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
})








