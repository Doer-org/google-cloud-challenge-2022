import {MockApiClient} from "../lib/ApiClient"

const apiClient = MockApiClient()   

export module EStateApi {
    export const getEStates = apiClient.path("/e-states").method("get").create()
    export const createEState = apiClient.path("/e-states").method("post").create()
    export const findById = apiClient.path("/e-states/{id}").method("get").create()
    export const deleteById = apiClient.path("/e-states/{id}").method("delete").create()
    export const updateById = apiClient.path("/e-states/{id}").method("patch").create()
    export const getEventsFilterByEState = apiClient.path("/e-states/{id}/event").method("get").create() 
}
export module ETypeApi {
    export const getETypes= apiClient.path("/e-types").method("get").create()
    export const createEType = apiClient.path("/e-types").method("post").create()
    export const findById = apiClient.path("/e-types/{id}").method("get").create()
    export const deleteById = apiClient.path("/e-types/{id}").method("delete").create()
    export const updateById = apiClient.path("/e-types/{id}").method("patch").create()
    export const getEventsFilterByEType = apiClient.path("/e-types/{id}/event").method("get").create() 
}

export module EventApi {
    export const getEventList = apiClient.path("/events").method("get").create()
    export const create = apiClient.path("/events").method("post").create()
    export const findById = apiClient.path("/events/{id}").method("get").create()
    export const deleteById = apiClient.path("/events/{id}").method("delete").create()
    export const updateById = apiClient.path("/events/{id}").method("patch").create()
    export const getEventStateById = apiClient.path("/events/{id}/state").method("get").create()
    export const getEventTypeById = apiClient.path("/events/{id}/type").method("get").create()
    export const getEventMembers = apiClient.path("/events/{id}/users").method("get").create() 
}

export module UserApi {   
    export const getUsers = apiClient.path("/users").method("get").create()
    export const createUser = apiClient.path("/users").method("post").create()
    export const findById = apiClient.path("/users/{id}").method("get").create()
    export const deleteById = apiClient.path("/users/{id}").method("delete").create()
    export const updateById = apiClient.path("/users/{id}").method("patch").create()
    export const getUsersEvents = apiClient.path("/users/{id}/events").method("get").create() 
} 
 