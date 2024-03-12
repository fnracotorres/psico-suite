## Documentación de API para DiskStat

### Estructura de DiskStat

Representa las estadísticas de un disco.

#### Propiedades

- **IOCountersStats**: `map[string]IOCountersStat`

  Un mapa que contiene estadísticas para varios contadores de E/S indexados por sus nombres.

- **PartitionStat**: `PartitionStat`

  Estadísticas sobre la partición del disco.

- **UsageStat**: `UsageStat`

  Estadísticas sobre el uso del disco.

### Estructura de UsageStat

Estadísticas sobre el uso del disco.

#### Propiedades

- **Path**: `string`

  La ruta del disco.

- **Fstype**: `string`

  El tipo de sistema de archivos del disco.

- **Total**: `uint64`

  Espacio total en disco en bytes.

- **Free**: `uint64`

  Espacio libre en disco en bytes.

- **Used**: `uint64`

  Espacio utilizado en disco en bytes.

- **UsedPercent**: `float64`

  El porcentaje de espacio en disco utilizado.

- **InodesTotal**: `uint64`

  Número total de inodos.

- **InodesUsed**: `uint64`

  Número de inodos utilizados.

- **InodesFree**: `uint64`

  Número de inodos libres.

- **InodesUsedPercent**: `float64`

  El porcentaje de inodos utilizados.

### Estructura de PartitionStat

Estadísticas sobre la partición del disco.

#### Propiedades

- **Device**: `string`

  El nombre del dispositivo de la partición.

- **Mountpoint**: `string`

  El punto de montaje de la partición.

- **Fstype**: `string`

  El tipo de sistema de archivos de la partición.

- **Opts**: `string`

  Opciones aplicadas a la partición.

### Estructura de IOCountersStat

Estadísticas sobre contadores de E/S.

#### Propiedades

- **ReadCount**: `uint64`

  Número total de operaciones de lectura.

- **MergedReadCount**: `uint64`

  Número total de operaciones de lectura fusionadas.

- **WriteCount**: `uint64`

  Número total de operaciones de escritura.

- **MergedWriteCount**: `uint64`

  Número total de operaciones de escritura fusionadas.

- **ReadBytes**: `uint64`

  Número total de bytes leídos.

- **WriteBytes**: `uint64`

  Número total de bytes escritos.

- **ReadTime**: `uint64`

  Tiempo total dedicado a la lectura en milisegundos.

- **WriteTime**: `uint64`

  Tiempo total dedicado a la escritura en milisegundos.

- **IopsInProgress**: `uint64`

  Número de operaciones de E/S en progreso.

- **IoTime**: `uint64`

  Tiempo total dedicado a las operaciones de E/S en milisegundos.

- **WeightedIO**: `uint64`

  Tiempo ponderado dedicado a las operaciones de E/S en milisegundos.

- **Name**: `string`

  Nombre del disco.

- **SerialNumber**: `string`

  Número de serie del disco.

- **Label**: `string`

  Etiqueta del disco.
