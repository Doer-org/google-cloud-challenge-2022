import { createApiClient } from '../lib/ApiClient';
import * as TE from 'fp-ts/TaskEither';
import { pipe } from 'fp-ts/lib/function';

const baseUrl =
    process.env.NEXT_PUBLIC_SERVER_URL ?? "http://localhost:8080" 
     //process.env.NEXT_PUBLIC_SERVER_URL ?? "https://gc-api-qgai5lo5hq-an.a.run.app" 
    // // "http://localhost:8080" 
const apiClient = createApiClient(baseUrl)

export module EventApi {
  /**作成したイベントがresponseとして返ってきます。 注 : adminはこのAPIをたたいたタイミングでイベントに参加したとみなされます。 つまり、他のAPIを用いてadminをeventに参加させる処理を行う必要はありません。 */
  export const createEvent = apiClient.path('/events').method('post').create();

  /** idでイベントを取得します。このidが共有するURLに使われます。 */
  export const getEvent = apiClient.path('/events/{id}').method('get').create();
  export const deleteEvent = (param: { id: string }) =>
    //apiClient.path("/events/{id}").method("delete").create() // "Unable to find content for application/json",
    pipe(
      TE.tryCatch(
        () =>
          fetch(`${baseUrl}/events/${param.id}`, {
            method: 'DELETE',
            credentials: 'include',
          }),
        (e: any) => Error(e)
      ),
      TE.chain((resp) =>
        resp.ok
          ? TE.right(resp)
          : TE.left(Error(`${resp.status} : ${resp.statusText}`))
      )
    );

  export const updateEvent = apiClient
    .path('/events/{id}')
    .method('patch')
    .create();
  export const getEventHost = apiClient
    .path('/events/{id}/admin')
    .method('get')
    .create();
  /**  このAPIを同時にたたいて、コメント一覧も取得してください。 */
  export const getEventComments = apiClient
    .path('/events/{id}/comments')
    .method('get')
    .create();

  /** nameのみが必須です。commentがある場合は、以下のようにrequest bodyに入れてください。 */
  // export const join = apiClient.path("/events/{id}/participants").method("post").create()

  /**
   * No response in the range 200-299 defined
   * nameのみが必須です。commentがある場合は、以下のようにrequest bodyに入れてください。
   * */
  export const join = (id: string, body: { name: string; comment: string }) =>
    pipe(
      TE.tryCatch(
        () =>
          fetch(`${baseUrl}/events/${id}/participants`, {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify(body),
          }),
        (e: any) => Error(e)
      ),
      TE.chain((resp) =>
        resp.ok
          ? TE.right(resp)
          : TE.left(Error(`${resp.status} : ${resp.statusText}`))
      )
    );

  // export const updateEventState = apiClient.path("/events/{id}/state").method("patch").create()
  /** No response in the range 200-299 defined */
  export const updateEventState = (
    param: { id: string },
    state: 'close' | 'cancel' | 'open'
  ) =>
    pipe(
      TE.tryCatch(
        () =>
          fetch(`${baseUrl}/events/${param.id}/state`, {
            method: 'PATCH',
            credentials: 'include',
            body: JSON.stringify({
                state : state
            })
          }),
        (e: any) => Error(e)
      ),
      TE.chain((resp) =>
        resp.ok
          ? TE.right(resp)
          : TE.left(Error(`${resp.status} : ${resp.statusText}`))
      )
    );

  /** このAPIで参加者を取得できますが、同時にコメントは取得できません。ホストを含む*/
  export const getEventMembers = apiClient
    .path('/events/{id}/users')
    .method('get')
    .create();
}

export module UserApi {
  export const createUser = apiClient.path('/users').method('post').create();
  export const getUser = apiClient.path('/users/{id}').method('get').create();
  //export const deleteById = apiClient.path("/users/{id}").method("delete").create()
  export const deleteById = (param: { id: string }) =>
    pipe(
      TE.tryCatch(
        () =>
          fetch(`${baseUrl}/users/${param.id}`, {
            method: 'DELETE',
            credentials: 'include',
          }),
        (e: any) => Error(e)
      ),
      TE.chain((resp) =>
        resp.ok
          ? TE.right(resp)
          : TE.left(Error(`${resp.status} : ${resp.statusText}`))
      )
    );
  export const updateById = apiClient
    .path('/users/{id}')
    .method('patch')
    .create();
  export const getUsersEvents = apiClient
    .path('/users/{id}/events')
    .method('get')
    .create();
}
