import webapi from "./gocliRequest"
import * as components from "./demoComponents"
export * from "./demoComponents"

/**
 * @description 
 * @param params
 */
export function demo(params: components.RequestParams, name: string) {
	return webapi.get<components.Response>(`/from/${name}`, params)
}
