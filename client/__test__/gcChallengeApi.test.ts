import { describe, it, expect, test } from 'vitest' 
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import * as E from 'fp-ts/Either'
import { EventApi } from '../core/utils/gcChallengeApi'  
import { fptsHelper } from '../core/utils/fptsHelper'
 
describe('[gcChallengeApi] EventApi.createEvent', () => {
  it('正常系：Eventを作成', async () => {  
    const resp = await pipe( 
        EventApi.createEvent({ 
            name: "string", 
            type: "string",
            state:  "string",
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
