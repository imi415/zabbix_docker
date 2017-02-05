# Docker Monitoring for Zabbix

This small go plugin is for providing simple API wrapping for zabbix-agent's user parameter, and can be used with unmodified agent and server.

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
