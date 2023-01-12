import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
 
export const config = {
  matcher: ['/event/new', '/event/:eventId*'],
}

export const middleware = (req: NextRequest) => {
  return NextResponse.next()


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