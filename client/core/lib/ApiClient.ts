import { paths } from "../openapi/openapi"  
import { Fetcher } from "openapi-typescript-fetch";
/**
 * 
 *  http://localhost:8080 => gc api, 
 *  http://localhost:8003 => swagger api
 */
export const createApiClient = (
    baseUrl : "http://localhost:8080" | "http://localhost:8003"
) => { 
    const fetcher = Fetcher.for<paths>()
    fetcher.configure({
        baseUrl: baseUrl,
    }) 
    return fetcher
}