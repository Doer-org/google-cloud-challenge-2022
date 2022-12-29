import { paths } from "../openapi/openapi"  
import { Fetcher } from "openapi-typescript-fetch";

export const MockApiClient = () => { 
    const fetcher = Fetcher.for<paths>()
    fetcher.configure({
        baseUrl: 'http://localhost:8003', 
    }) 
    return fetcher
}