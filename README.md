# Docker Monitoring for Zabbix

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
