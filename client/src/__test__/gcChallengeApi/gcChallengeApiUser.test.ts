import { describe, it, expect, test } from 'vitest' 
import { flow, pipe } from 'fp-ts/lib/function' 
import * as E from 'fp-ts/Either'
import * as TE from 'fp-ts/TaskEither'
import { UserApi } from '../../core/utils/gcChallengeApi'  
import { fptsHelper } from '../../core/utils/fptsHelper'
 
const createUser = () => pipe( 
  UserApi.createUser({ 
      name: "aoki",   
      authenticated: true,
      icon: "string"
  }),
  fptsHelper.TE.ofApiResponse, 
)

describe('[gccAPI] UserApi.createUser', () => {
  it('正常系：Userを作成', async () => {  
    const resp = await createUser()()
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
}) 


describe('[gccAPI] UserApi.getUser', () => {
  it('正常系：Userを取得', async () => {  
    const user = createUser()
    const resp = await pipe( 
        user,
        TE.chain((user) => pipe(
          UserApi.getUser({id: user.id}),
          fptsHelper.TE.ofApiResponse,
          TE.map((user2) => 
            user.id === user2.id
            && user.name === user2.name
          )
        )) 
    )()
    expect(  
      E.isRight(resp) && resp.right
    )
    .toBe(
        true 
    ) 
  })
}) 


describe('[gccAPI] UserApi.deleteById', () => {
  it('正常系：Userを削除', async () => {   
    const user = createUser()
    const resp = await pipe( 
        user,
        TE.chain((user) => pipe( 
          UserApi.deleteById({ 
              id: user.id
          }) 
        )) 
    )() 
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
}) 

describe('[gccAPI] UserApi.updateById', () => {
  it('正常系：Userを更新', async () => {   
    // const user1 = await createUser()()
    // if (user1._tag === "Right") {
    //   const u = user1.right
    //   const t = await UserApi.updateById({
    //     id: u.id,
    //     name: `updated ${u.name}`
    //     //, icon : "aaaa" //一文字以上ないとError！スキーマ定義ではオプション
    //   })
    //   console.log(t)
    // }  
    const user = createUser()
    const resp = await pipe( 
        user,
        TE.chain((user) => pipe( 
          UserApi.updateById({ 
              id: user.id,
              name: `updated ${user.name}`
              , icon : "notnull"
          }),
          fptsHelper.TE.ofApiResponse,
          TE.map((user2) => 
            user.id === user2.id
            && user2.name === `updated ${user.name}`
          )
        )) 
    )()
    expect( 
      E.isRight(resp) && resp.right
    ).toBe(true) 
  })
}) 


describe('[gccAPI] UserApi.getUsersEvents', () => {
  it('正常系：Userが主催したイベントを取得。（リストは空）', async () => {  
    const user = createUser()
    const resp = await pipe( 
        user,
        TE.chain((user) => pipe( 
          UserApi.getUsersEvents({ 
              id: user.id
          }),
          fptsHelper.TE.ofApiResponse,
          TE.map((events) =>  events.length === 0)
        )) 
    )() 
    expect( 
        E.isRight(resp)
    )
    .toBe(
        true 
    ) 
  })
}) 

