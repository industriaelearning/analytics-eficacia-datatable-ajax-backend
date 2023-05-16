import React, { ReactElement } from 'react';
import { FieldSet, InlineField, Input, LegacyForms } from '@grafana/ui';
import type { EditorProps } from './types';
import { useChangeOptions } from './useChangeOptions';
import { useChangeSecureOptions } from './useChangeSecureOptions';
import { useResetSecureOptions } from './useResetSecureOptions';
import { testIds } from '../testIds';

const { SecretFormField } = LegacyForms;

export function ConfigEditor(props: EditorProps): ReactElement {
  const { jsonData, secureJsonData, secureJsonFields } = props.options;
  const onPostgresHostChange = useChangeOptions(props, 'postgresHost');
  const onPostgresPortChange = useChangeOptions(props, 'postgresPort');
  const onPostgresUserChange = useChangeOptions(props, 'postgresUsername');
  const onPostgresDatabaseChange = useChangeOptions(props, 'postgresDatabase');
  const onPostgresPasswordChange = useChangeSecureOptions(props, 'postgresPassword');
  const onResetPostgresPassword = useResetSecureOptions(props, 'postgresPassword');

    return (
        <>

        <FieldSet label="Postgres Settings">
            <InlineField label="Postgresql host" tooltip="Postgresql host">
                <Input
                    onChange={onPostgresHostChange}
                    placeholder="host"
                    data-testid={testIds.configEditor.postgresHost}
                    value={jsonData?.postgresHost ?? ''}
                />
            </InlineField>

            <InlineField label="Postgresql port" tooltip="Postgresql port">
                <Input
                    onChange={onPostgresPortChange}
                    placeholder="5432"
                    data-testid={testIds.configEditor.postgresPort}
                    value={jsonData?.postgresPort ?? ''}
                />
            </InlineField>

            <InlineField label="Postgresql database" tooltip="Postgresql database">
                <Input
                    onChange={onPostgresDatabaseChange}
                    placeholder="Database name"
                    data-testid={testIds.configEditor.postgresDatabase}
                    value={jsonData?.postgresDatabase ?? ''}
                />
            </InlineField>

            <InlineField label="Postgresql user" tooltip="Postgresql user">
                <Input
                    onChange={onPostgresUserChange}
                    placeholder="user"
                    data-testid={testIds.configEditor.postgresUsername}
                    value={jsonData?.postgresUsername ?? ''}
                />
            </InlineField>

            <SecretFormField
                tooltip="Postgres Password"
                isConfigured={Boolean(secureJsonFields.postgresPassword)}
                value={secureJsonData?.postgresPassword || ''}
                label="Postgresql password"
                placeholder="postgresql password"
                labelWidth={12}
                inputWidth={20}
                data-testid={testIds.configEditor.postgresPassword}
                onReset={onResetPostgresPassword}
                onChange={onPostgresPasswordChange}
            />

        </FieldSet>
        </>
    );
}
