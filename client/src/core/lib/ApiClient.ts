import { paths } from '../openapi/openapi';
import { Fetcher } from 'openapi-typescript-fetch';
/**
 *
 *  http://localhost:8080 => gc api,
 *  http://localhost:8003 => swagger api
 *  https://gc-api-qgai5lo5hq-an.a.run.app => デプロイ先
 */
export const createApiClient = (
  baseUrl: string
    // | 'http://localhost:8080'
    // | 'http://localhost:8003'
    // | 'https://gc-api-qgai5lo5hq-an.a.run.app'
) => {
  const fetcher = Fetcher.for<paths>();
  fetcher.configure({
    baseUrl: baseUrl,
  });
  return fetcher;
};
