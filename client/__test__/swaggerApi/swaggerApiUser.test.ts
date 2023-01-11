import { describe, it, expect, test } from 'vitest' 
import { flow, pipe } from 'fp-ts/lib/function' 
import * as E from 'fp-ts/Either'
import { UserApi } from '../../core/utils/swaggerApi'  
import { fptsHelper } from '../../core/utils/fptsHelper'
 
describe('[swaggerApi] UserApi.createUser', () => {
  it('正常系：Userを作成', async () => {  
    const resp = await pipe( 
        UserApi.createUser({ 
            name: "string",   
            authenticated: true,
            icon: "string"
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


describe('[swaggerApi] UserApi.getUser', () => {
  it('正常系：Userを取得', async () => {  
    const resp = await pipe( 
        UserApi.getUser({ 
            id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",  
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


describe('[swaggerApi] UserApi.deleteById', () => {
  it('正常系：Userを削除', async () => {  
    // const resp = await pipe( 
    //     UserApi.deleteById({ 
    //         id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",  
    //     }),
    //     fptsHelper.TE.ofApiResponse, 
    // )()
    const resp = 
      await UserApi.deleteById({ 
        id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",  
      })()
    // console.log(resp)
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
}) 

describe('[swaggerApi] UserApi.updateById', () => {
  it('正常系：Userを更新', async () => {  
    const resp = await pipe( 
        UserApi.updateById({ 
            id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
            name: "new aoki"
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


describe('[swaggerApi] UserApi.getUsersEvents', () => {
  it('正常系：Userが主催したイベントを取得', async () => {  
    const resp = await pipe( 
        UserApi.getUsersEvents({ 
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

