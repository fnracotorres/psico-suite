## Documentación de API para DiskStat

### Clase DiskStat

Representa las estadísticas de un disco.

#### Propiedades

- **IOCountersStats**: `Map<String, IOCountersStat>`

  Un mapa que contiene estadísticas para varios contadores de E/S indexados por sus nombres.

- **PartitionStat**: `PartitionStat`

  Estadísticas sobre la partición del disco.

- **UsageStat**: `UsageStat`

  Estadísticas sobre el uso del disco.

### Clase UsageStat

Estadísticas sobre el uso del disco.

#### Propiedades

- **path**: `String`

  La ruta del disco.

- **fstype**: `String`

  El tipo de sistema de archivos del disco.

- **total**: `long`

  Espacio total en disco en bytes.

- **free**: `long`

  Espacio libre en disco en bytes.

- **used**: `long`

  Espacio utilizado en disco en bytes.

- **usedPercent**: `double`

  El porcentaje de espacio en disco utilizado.

- **inodesTotal**: `long`

  Número total de inodos.

- **inodesUsed**: `long`

  Número de inodos utilizados.

- **inodesFree**: `long`

  Número de inodos libres.

- **inodesUsedPercent**: `double`

  El porcentaje de inodos utilizados.

### Clase PartitionStat

Estadísticas sobre la partición del disco.

#### Propiedades

- **device**: `String`

  El nombre del dispositivo de la partición.

- **mountpoint**: `String`

  El punto de montaje de la partición.

- **fstype**: `String`

  El tipo de sistema de archivos de la partición.

- **opts**: `String`

  Opciones aplicadas a la partición.

### Clase IOCountersStat

Estadísticas sobre contadores de E/S.

#### Propiedades

- **readCount**: `long`

  Número total de operaciones de lectura.

- **mergedReadCount**: `long`

  Número total de operaciones de lectura fusionadas.

- **writeCount**: `long`

  Número total de operaciones de escritura.

- **mergedWriteCount**: `long`

  Número total de operaciones de escritura fusionadas.

- **readBytes**: `long`

  Número total de bytes leídos.

- **writeBytes**: `long`

  Número total de bytes escritos.

- **readTime**: `long`

  Tiempo total dedicado a la lectura en milisegundos.

- **writeTime**: `long`

  Tiempo total dedicado a la escritura en milisegundos.

- **iopsInProgress**: `long`

  Número de operaciones de E/S en progreso.

- **ioTime**: `long`

  Tiempo total dedicado a las operaciones de E/S en milisegundos.

- **weightedIO**: `long`

  Tiempo ponderado dedicado a las operaciones de E/S en milisegundos.

- **name**: `String`

  Nombre del disco.

- **serialNumber**: `String`

  Número de serie del disco.

- **label**: `String`

  Etiqueta del disco.
