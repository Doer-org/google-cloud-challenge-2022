import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
 
export const config = {
  matcher: [
    '/event/new', 
    '/event/:eventId/admin*'
  ],
}

export const middleware = async (req: NextRequest) => {
  const resp = await fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/validate`, {
    method: 'GET',
    headers: req.headers
  })
  if (resp.ok) {
    return NextResponse.next()
  } else {
    return NextResponse.rewrite(req.nextUrl.basePath)
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