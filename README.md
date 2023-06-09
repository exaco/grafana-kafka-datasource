# Kafka Datasource for Grafana
[![License](https://img.shields.io/github/license/hoptical/grafana-kafka-datasource)](LICENSE)
[![CI](https://github.com/hoptical/grafana-kafka-datasource/actions/workflows/ci.yml/badge.svg)](https://github.com/hoptical/grafana-kafka-datasource/actions/workflows/ci.yml)
[![Release](https://github.com/hoptical/grafana-kafka-datasource/actions/workflows/release.yml/badge.svg)](https://github.com/hoptical/grafana-kafka-datasource/actions/workflows/release.yml)

The Kafka data source plugin allows you to visualize streaming Kafka data from within Grafana.

## Reqirements

- Apache Kafka v0.9+
- Grafana v8.0+

> Note: This is a backend plugin, so the Grafana server should've access to the Kafka broker.

## Getting started

### Installation via grafana-cli tool

Use the grafana-cli tool to install the plugin from the commandline:

```bash
grafana-cli plugins install hamedkarbasi93-kafka-datasource
```

The plugin will be installed into your grafana plugins directory; the default is `/var/lib/grafana/plugins`. [More information on the cli tool](https://grafana.com/docs/grafana/latest/administration/cli/#plugins-commands).

### Installation via zip file

Alternatively, you can manually download the [latest](https://github.com/hoptical/grafana-kafka-datasource/releases/latest) release .zip file and unpack it into your grafana plugins directory; the default is `/var/lib/grafana/plugins`.

## Configure the data source

[Add a data source](https://grafana.com/docs/grafana/latest/datasources/add-a-data-source/) by filling in the following fields:

### Basic fields

| Field | Description                                        |
| ----- | -------------------------------------------------- |
| Name  | A name for this particular AppDynamics data source |
| Servers  | The URL of the Kafka bootstrap servers separated by comma. E.g. `broker1:9092, broker2:9092`              |

### Query the Data source

To query the Kafka topic, you have to config the below items in the query editor.

| Field | Description                                        |
| ----- | -------------------------------------------------- |
| Topic  | Topic Name |
| Partition  | Partition Number |
| Auto offset reset | Starting offset to consume that can be from latest or last 100. |
| Timestamp Mode | Timestamp of the message value to visualize; It can be Now or Message Timestamp
> **Note**: Make sure to enable the `streaming` toggle.

![kafka dashboard](https://raw.githubusercontent.com/hoptical/grafana-kafka-datasource/86ea8d360bfd67cfed41004f80adc39219983210/src/img/graph.gif)

## Known limitations

- The plugin currently does not support any authorization and authentication method.
- The plugin currently does not support TLS.
- Plugin is based on [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go), hence it only supports Linux-based operating systems (not available on arm64 known as Apple Silicon) as discussed in [#6](https://github.com/hoptical/grafana-kafka-datasource/issues/6). However, we're cosidering changing the base package to support all operating systems.

This plugin supports topics publishing very simple JSON formatted messages. Note that only the following structure is supported as of now:

```json
{
    "value1": 1.0,
    "value2": 2,
    "value3": 3.33,
    ...
}
```

We plan to support more complex JSON data structures, Protobuf and AVRO in the upcoming releases. Contributions are highly encouraged!
## Compiling the data source by yourself

Note: there are build issues with Node v16.  Use Node v14.  Known working dev environments:

* Ubuntu 20.04 LTS
* node v14.19.3 (npm 6.14.17)
* yarn 1.22.19
* go 1.18

and

* Ubuntu 22.04 LTS
* node v14.20.0 (npm 6.14.17)
* yarn 1.22.19
* go 1.18.1

A data source backend plugin consists of both frontend and backend components.

### Frontend

1. Install dependencies

   ```bash
   yarn install
   ```

2. Build plugin in development mode or run in watch mode

   ```bash
   yarn dev
   ```

   or

   ```bash
   yarn watch
   ```

3. Build plugin in production mode

   ```bash
   yarn build
   ```

### Backend

1. Update [Grafana plugin SDK for Go](https://grafana.com/docs/grafana/latest/developers/plugins/backend/grafana-plugin-sdk-for-go/) dependency to the latest minor version:

   ```bash
   go get -u github.com/grafana/grafana-plugin-sdk-go
   go mod tidy
   ```

2. Build backend plugin binaries for Linux:

   ```bash
   mage build:backend
   ```

## Contributing

Thank you for considering contributing! If you find an issue or have a better way to do something, feel free to open an issue or a PR.

## License

This repository is open-sourced software licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).

## Learn more

- [Build a data source backend plugin tutorial](https://grafana.com/tutorials/build-a-data-source-backend-plugin)

- [Grafana plugin SDK for Go](https://grafana.com/docs/grafana/latest/developers/plugins/backend/grafana-plugin-sdk-for-go/)
