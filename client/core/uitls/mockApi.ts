import {MockApiClient} from "../lib/ApiClient"

const apiClient = MockApiClient()


export module EventApi {
    /**作成したイベントがresponseとして返ってきます。 注 : adminはこのAPIをたたいたタイミングでイベントに参加したとみなされます。 つまり、他のAPIを用いてadminをeventに参加させる処理を行う必要はありません。 */
    export const createEvent = apiClient.path("/events").method("post").create()

    /** idでイベントを取得します。このidが共有するURLに使われます。 */
    export const getEvent = apiClient.path("/events/{id}").method("get").create()
    export const deleteEvent = apiClient.path("/events/{id}").method("delete").create()
    export const updateEvent = apiClient.path("/events/{id}").method("patch").create()
    export const getEventHost = apiClient.path("/events/{id}/admin").method("get").create()
    /**  このAPIを同時にたたいて、コメント一覧も取得してください。 */
    export const getEventComments = apiClient.path("/events/{id}/comments").method("get").create() 

    /** nameのみが必須です。commentがある場合は、以下のようにrequest bodyに入れてください。 */
    export const join = apiClient.path("/events/{id}/participants").method("post").create()
    export const updateEventState = apiClient.path("/events/{id}/state").method("patch").create()

    /** このAPIで参加者を取得できますが、同時にコメントは取得できません。*/
    export const getEventMembers = apiClient.path("/events/{id}/users").method("get").create() 
}

export module UserApi {   
    export const createUser = apiClient.path("/users").method("post").create()
    export const getUser = apiClient.path("/users/{id}").method("get").create()
    export const deleteById = apiClient.path("/users/{id}").method("delete").create()
    export const updateById = apiClient.path("/users/{id}").method("patch").create()
    export const getUsersEvents = apiClient.path("/users/{id}/events").method("get").create()  
} 
 