#!/usr/bin/env python3

import yaml

SPEC_PATH = 'api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

new_status_values = [
    "racked", "validating", "inspecting", "archive", "available", "provisioning"
]
new_status_labels = [
    "Racked", "Validating", "Inspecting", "Archive", "Available", "Provisioning"
]
objects_to_modify = [
    "Device", "DeviceWithConfigContext", "DeviceWithConfigContextRequest", "WritableDeviceWithConfigContextRequest"
] 

if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():
        if 'properties' in schema:
            # Remove "null" item from nullable enums
            for prop_name, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]
            for ntype in nullable_types:
                if ntype in schema['properties']:
                    schema['properties'][ntype]['nullable'] = True

            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]
            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    # Only attempt to pop 'nullable' if it exists
                    if 'nullable' in schema['properties'][ntype]:
                        schema['properties'][ntype].pop('nullable')

            ################################################################################
            # ignore the required fields. (this is for fixing API Spec and actual data schema mismatch issues)
            ################################################################################
            # Make all *_count fields not required
            if 'required' in schema:
                updated_required = [
                    r for r in schema['required']
                    if not r.endswith('_count')
                ]
                schema['required'] = updated_required


            # Specific fix for 'Device'
            if name == "Device" and 'required' in schema:
                schema['required'] = [
                    r for r in schema['required']
                    if r != 'device_type' and r != 'site' and r != 'role' and r != 'parent_device' and r != 'primary_ip' and r != 'created' and r != 'last_updated'
                ]
   
            # Specific fix for 'DeviceWithConfigContext'
            # if name == "Device" or name == "DeviceWithConfigContext" and 'required' in schema:
            if name == "DeviceWithConfigContext" and 'required' in schema:
                schema['required'] = [
                    r for r in schema['required']
                    if r != 'role' and r != 'parent_device'
                ]
                
            ################################################################################
            # add custom field values.
            ################################################################################

            # Specific fix: Add 'inspecting' and 'racked' to the Device status enum
            if name in objects_to_modify and 'properties' in schema and 'status' in schema['properties']:
                status = schema['properties']['status']
                if 'properties' in status and 'value' in status['properties'] and 'enum' in status['properties']['value']:
                    for new_value in new_status_values:
                        if new_value not in status['properties']['value']['enum']:
                            status['properties']['value']['enum'].append(new_value)

                if 'properties' in status and 'label' in status['properties'] and 'enum' in status['properties']['label']:
                    for new_label in new_status_labels:
                        if new_label not in status['properties']['label']['enum']:
                            status['properties']['label']['enum'].append(new_label)
                    

# Save the updated spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)

