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
    EntDistrict,
    EntDistrictFromJSON,
    EntDistrictFromJSONTyped,
    EntDistrictToJSON,
    EntStudent,
    EntStudentFromJSON,
    EntStudentFromJSONTyped,
    EntStudentToJSON,
    EntSubdistrict,
    EntSubdistrictFromJSON,
    EntSubdistrictFromJSONTyped,
    EntSubdistrictToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProvinceEdges
 */
export interface EntProvinceEdges {
    /**
     * 
     * @type {EntDistrict}
     * @memberof EntProvinceEdges
     */
    district?: EntDistrict;
    /**
     * 
     * @type {EntStudent}
     * @memberof EntProvinceEdges
     */
    student?: EntStudent;
    /**
     * 
     * @type {EntSubdistrict}
     * @memberof EntProvinceEdges
     */
    subdistrict?: EntSubdistrict;
}

export function EntProvinceEdgesFromJSON(json: any): EntProvinceEdges {
    return EntProvinceEdgesFromJSONTyped(json, false);
}

export function EntProvinceEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProvinceEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'district': !exists(json, 'District') ? undefined : EntDistrictFromJSON(json['District']),
        'student': !exists(json, 'Student') ? undefined : EntStudentFromJSON(json['Student']),
        'subdistrict': !exists(json, 'Subdistrict') ? undefined : EntSubdistrictFromJSON(json['Subdistrict']),
    };
}

export function EntProvinceEdgesToJSON(value?: EntProvinceEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'district': EntDistrictToJSON(value.district),
        'student': EntStudentToJSON(value.student),
        'subdistrict': EntSubdistrictToJSON(value.subdistrict),
    };
}


