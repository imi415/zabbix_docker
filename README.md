# Docker Monitoring for Zabbix

This small go plugin is for providing simple API wrapping for zabbix-agent's user parameter, and can be used with unmodified agent and server.
- We also have two other similar agent scripts written in Ruby and Node.JS, however duing our test, they consume more CPU resource than I can accept.

## Usage

- This agent caches resource data(JSON, named by container IDs) at $HOME(expected `/var/lib/zabbix` as user `zabbix`), please make sure this directory exists and we have write access to this directory.
- Upload the user parameters file to `/etc/zabbix/zabbix\_agentd.d/`.
- Upload the compiled binary to `/etc/zabbix/zabbix\_agentd.d/bin/`.
- Import zabbix template file.
- Add template to desired host.

## Template

This template file contains

- Docker daemon ID, server version, container amounts monitoring.
- Discovery rule for containers(by ID) and apps(by image).
- Container basic resources(CPU, memory, network I/O) monitoring.
- Several graphs.
- Several triggers.

## API Reference

|Method Calls|`$1`|`$2`|`$3`|`$4`|
|---|---|---|---|---|
|System-wide Informarion [1]|`Info`|<1>|<2>|`null`|
|Container Information [2]|`Container`|`$CONTAINER_ID`|<1>|<2>|
|Discovery[3]|`Discovery`|<1>|<2>|`null`|

- [1]: Wrapper for `docker info`
  - <1>: ~~Get available parameters by invoking `./zbx_docker Info Debug`~~
- [2]: Wrapper for `docker $CONTAINER_ID stats`
  - <1>: Container ID
  - <2>: ~~Get available parameters by invoking `./zbx_docker Container $CONTAINER_ID Debug`~~
- [3]:Need to add discovery in zabbix dashboard
  - <1>: Discovery Type, `Container` or `Image`
  - <2>: ~~Get debug information for selected discovery.~~

## Disclaimer
Since this is my first Go program(except "hello-world"),  it could be *really* unreilable. Although I use it in our production environment, please think twice before use. Also please consider the following awesome projects.
* [zabbix-docker-monitoring](https://github.com/monitoringartist/zabbix-docker-monitoring)
