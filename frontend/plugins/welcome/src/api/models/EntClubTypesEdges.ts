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
    EntClub,
    EntClubFromJSON,
    EntClubFromJSONTyped,
    EntClubToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClubTypesEdges
 */
export interface EntClubTypesEdges {
    /**
     * Club holds the value of the club edge.
     * @type {Array<EntClub>}
     * @memberof EntClubTypesEdges
     */
    club?: Array<EntClub>;
}

export function EntClubTypesEdgesFromJSON(json: any): EntClubTypesEdges {
    return EntClubTypesEdgesFromJSONTyped(json, false);
}

export function EntClubTypesEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClubTypesEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'club': !exists(json, 'club') ? undefined : ((json['club'] as Array<any>).map(EntClubFromJSON)),
    };
}

export function EntClubTypesEdgesToJSON(value?: EntClubTypesEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'club': value.club === undefined ? undefined : ((value.club as Array<any>).map(EntClubToJSON)),
    };
}

