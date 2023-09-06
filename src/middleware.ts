import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { pb } from './lib/pocketbase'
 
export function middleware(request: NextRequest) {
  // Assume a "Cookie:nextjs=fast" header to be present on the incoming request
  // Getting cookies from the request using the `RequestCookies` API
  let cookie = request.cookies.get('pb_auth')
  console.log("HELLO MY DEAR") // => { name: 'nextjs', value: 'fast', Path: '/' }

  pb.authStore.loadFromCookie(cookie?.value.toString() || "");

  const response = NextResponse.next()
  
  return response
}