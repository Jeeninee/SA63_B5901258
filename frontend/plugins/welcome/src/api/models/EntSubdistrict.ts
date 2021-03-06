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
    EntSubdistrictEdges,
    EntSubdistrictEdgesFromJSON,
    EntSubdistrictEdgesFromJSONTyped,
    EntSubdistrictEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSubdistrict
 */
export interface EntSubdistrict {
    /**
     * 
     * @type {EntSubdistrictEdges}
     * @memberof EntSubdistrict
     */
    edges?: EntSubdistrictEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSubdistrict
     */
    id?: number;
    /**
     * Subdistrict holds the value of the "subdistrict" field.
     * @type {string}
     * @memberof EntSubdistrict
     */
    subdistrict?: string;
}

export function EntSubdistrictFromJSON(json: any): EntSubdistrict {
    return EntSubdistrictFromJSONTyped(json, false);
}

export function EntSubdistrictFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSubdistrict {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntSubdistrictEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'subdistrict': !exists(json, 'subdistrict') ? undefined : json['subdistrict'],
    };
}

export function EntSubdistrictToJSON(value?: EntSubdistrict | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntSubdistrictEdgesToJSON(value.edges),
        'id': value.id,
        'subdistrict': value.subdistrict,
    };
}


