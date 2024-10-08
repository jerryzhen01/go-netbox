#!/usr/bin/env python3

import yaml

SPEC_PATH = 'api/openapi.yaml'


# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
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

            # Make *_count fields not required
            if 'required' in schema:
                updated_required = [
                    r for r in schema['required']
                    if not r.endswith('_count')
                ]
                schema['required'] = updated_required

# Save the updated spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)

