import { readFileSync, writeFileSync } from 'fs';
import { join } from 'path';
import { parse } from 'protobufjs';

const protoPath = join(process.cwd(), 'idl.proto');
const outputPath = join(process.cwd(), 'src/lib/apiConfig.json');

async function generateConfig() {
  try {
    const root = await parse(readFileSync(protoPath, 'utf8'));
    const services = root.services;
    
    const config = Object.keys(services).map(serviceName => {
      return services[serviceName].methods.map(method => ({
        name: method.name,
        path: method.options['(api.post)'] || method.options['(api.get)'],
        method: method.options['(api.post)'] ? 'POST' : 'GET',
        requestFields: method.resolvedRequestType.fieldsArray.map(field => ({
          name: field.name,
          type: field.type,
          rule: field.rule
        }))
      }));
    }).flat();

    writeFileSync(outputPath, JSON.stringify(config, null, 2));
    console.log('API config generated successfully');
  } catch (err) {
    console.error('Error generating API config:', err);
    process.exit(1);
  }
}

generateConfig();