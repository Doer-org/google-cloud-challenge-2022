import { paths } from "../openapi/openapi"  
import { Fetcher } from "openapi-typescript-fetch";
import {MockApiClient} from "../lib/ApiClient"

export default  () => {
    const test = () => { 
        const f = MockApiClient() 
         
        const getUserFromId = f.path("/users/{id}").method("get").create()
        const r = getUserFromId({id: 123}) 
        return r.then((r) => { 
            console.log(r.data)
            return r 
        }).catch((e) => e)
    } 
    return {
        test
    }
}