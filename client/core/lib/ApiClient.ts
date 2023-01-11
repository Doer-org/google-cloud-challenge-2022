import { paths } from "../openapi/openapi"  
import { Fetcher } from "openapi-typescript-fetch";

export const SwaggerApiClient = () => { 
    const fetcher = Fetcher.for<paths>()
    fetcher.configure({
        baseUrl: 'http://localhost:8003',  
    }) 
    return fetcher
}

export const ApiClient = () => { 
    const fetcher = Fetcher.for<paths>()
    fetcher.configure({
        baseUrl: 'http://localhost:8080',
    }) 
    return fetcher
}