import {MockApiClient} from "../lib/ApiClient"

const apiClient = MockApiClient()

export module EventApi {
    export const createEvent = apiClient.path("/events").method("post").create()
    export const getEvent = apiClient.path("/events/{id}").method("get").create()
    export const deleteEvent = apiClient.path("/events/{id}").method("delete").create()
    export const updateEvent = apiClient.path("/events/{id}").method("patch").create()
    export const getEventHost = apiClient.path("/events/{id}/admin").method("get").create()
    export const getEventComments = apiClient.path("/events/{id}/comments").method("get").create()
    export const updateEventState = apiClient.path("/events/{id}/state").method("patch").create()
    export const getEventType = apiClient.path("/events/{id}/type").method("patch").create()
    export const getEventMembers = apiClient.path("/events/{id}/users").method("get").create() 
}

export module UserApi {   
    export const createUser = apiClient.path("/users").method("post").create()
    export const getUser = apiClient.path("/users/{id}").method("get").create()
    export const deleteById = apiClient.path("/users/{id}").method("delete").create()
    export const updateById = apiClient.path("/users/{id}").method("patch").create()
    export const getUsersEvents = apiClient.path("/users/{id}/events").method("get").create()  
} 
 