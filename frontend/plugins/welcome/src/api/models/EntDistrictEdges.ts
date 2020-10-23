/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntProvince,
    EntProvinceFromJSON,
    EntProvinceFromJSONTyped,
    EntProvinceToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDistrictEdges
 */
export interface EntDistrictEdges {
    /**
     * Province holds the value of the province edge.
     * @type {Array<EntProvince>}
     * @memberof EntDistrictEdges
     */
    province?: Array<EntProvince>;
}

export function EntDistrictEdgesFromJSON(json: any): EntDistrictEdges {
    return EntDistrictEdgesFromJSONTyped(json, false);
}

export function EntDistrictEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDistrictEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'province': !exists(json, 'province') ? undefined : ((json['province'] as Array<any>).map(EntProvinceFromJSON)),
    };
}

export function EntDistrictEdgesToJSON(value?: EntDistrictEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'province': value.province === undefined ? undefined : ((value.province as Array<any>).map(EntProvinceToJSON)),
    };
}


