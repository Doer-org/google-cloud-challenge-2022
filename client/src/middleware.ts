// import { NextResponse } from 'next/server'
// import type { NextRequest } from 'next/server'
// // import { getEventList, tryGetEventList } from './core/api/user/getEventList'
// // import * as TE from 'fp-ts/TaskEither'
// // import * as O from 'fp-ts/Option'
// // import * as A from 'fp-ts/Array'
// // import { flow, pipe } from 'fp-ts/lib/function'

// export const config = {
//   matcher: [
//     '/',
//     '/event/new',
//     '/event/:eventId*/admin(.)*'
//   ],
// }

// export const middleware = async (req: NextRequest) => {
//   console.log("middleware")
//   // const resp = await fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/ping`, {
//   //   method: 'GET',
//   //   headers: req.headers
//   // })
//   // const body = await resp.json()
//   // if (body.code === 400) {
//   //   return NextResponse.redirect(`${req.nextUrl.origin}/auth`)
//   // }
//   // return NextResponse.next()
//   // const resp = await fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/validate`, {
//   //   method: 'GET',
//   //   headers: req.headers
//   // })
//   // const body = await resp.json()
//   // if (body.code === 400) {
//   //   return NextResponse.redirect(`${req.nextUrl.origin}/auth`)
//   // } else {
//   //   return NextResponse.next()
//   // }
//   return NextResponse.next()
//   // // const path = req.nextUrl.pathname
//   // // const regex = RegExp('/event/([^\/]*)/admin.*')
//   // // const eventId = regex.exec(path)?.[1]  
//   // // if (eventId) { 
//   // //   console.log(eventId)
//   // //   const user = await fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/user`, {
//   // //     method: 'GET',
//   // //     headers: req.headers
//   // //   }).then(async (ok) => ok.json()) //.catch((e) => {console.log(e)})
//   // //   const next = await pipe(
//   // //     tryGetEventList(user.id),
//   // //     TE.match(
//   // //       (e) => (NextResponse.redirect(`${req.nextUrl.origin}/event`)),
//   // //       (ok) => {
//   // //         return NextResponse.next()
//   // //         console.log(user)
//   // //         console.log(ok)
//   // //         // pipe(ok, A.map((a) => {console.log(a.id)}))
//   // //         return pipe(ok, A.findFirst((a) => (a.id === eventId)), O.isSome)
//   // //         ? NextResponse.next()
//   // //         : NextResponse.redirect(`${req.nextUrl.origin}/event`)
//   // //       }
//   // //     )
//   // //   )()
//   // //   return next
//   // // } 
//   // // return NextResponse.next()
// }

import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
 
export const config = {
  matcher: [
    '/',
    '/event/new', 
    '/event/:eventId*/admin'
  ],
}

export const middleware = async (req: NextRequest) => {
  // return NextResponse.next()
  const resp = await fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/validate`, {
    method: 'GET',
    headers: req.headers
  })
  const body = await resp.json()
  if (body.code === 400) {
    return NextResponse.redirect(`${req.nextUrl.origin}/auth`)
  } else {
    return NextResponse.next()
  }
  // return NextResponse.next()
  // const authorizationHeader = req.headers.get('authorization')

  // if (authorizationHeader) {
  //   const basicAuth = authorizationHeader.split(' ')[1]
  //   const [user, password] = atob(basicAuth).split(':') 
  //   if (
  //     user === "a" && //process.env.BASIC_AUTH_USER &&
  //     password === "c" // process.env.BASIC_AUTH_PASSWORD
  //   ) {
  //     return NextResponse.next()
  //   }
  // } 

  // const url = req.nextUrl
  // url.pathname = '/api/auth' 
  // return NextResponse.rewrite(url)
}