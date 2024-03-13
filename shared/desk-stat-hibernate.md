## Documentación del Modelo Hibernate para DiskStat

### Clase DiskStat

Representa las estadísticas de un disco.

#### Propiedades

- **ioCountersStatList**: `Map<String, IOCountersStat>`

  Un mapa que contiene estadísticas para varios contadores de E/S indexados por sus nombres.

- **partitionStat**: `PartitionStat`

  Estadísticas sobre la partición del disco.

- **usageStat**: `UsageStat`

  Estadísticas sobre el uso del disco.

```java
@Entity
@Table(name = "disk_stat")
public class DiskStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id")
    private Long id;

    @OneToMany(mappedBy = "diskStat", cascade = CascadeType.ALL)
    private Map<String, IOCountersStat> ioCountersStatList;

    @OneToOne(mappedBy = "diskStat", cascade = CascadeType.ALL)
    private PartitionStat partitionStat;

    @OneToOne(mappedBy = "diskStat", cascade = CascadeType.ALL)
    private UsageStat usageStat;

    // getters y setters
}
```

### Clase UsageStat

Estadísticas sobre el uso del disco.

#### Propiedades

- **path**: `String`

  La ruta del disco.

- **fsType**: `String`

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

```java
@Entity
@Table(name = "usage_stat")
public class UsageStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id")
    private Long id;

    @OneToOne
    @JoinColumn(name = "disk_stat_id")
    private DiskStat diskStat;

    @Column(name = "path")
    private String path;

    @Column(name = "fs_type")
    private String fsType;

    @Column(name = "total")
    private long total;

    @Column(name = "free")
    private long free;

    @Column(name = "used")
    private long used;

    @Column(name = "used_percent")
    private double usedPercent;

    @Column(name = "inodes_total")
    private long inodesTotal;

    @Column(name = "inodes_used")
    private long inodesUsed;

    @Column(name = "inodes_free")
    private long inodesFree;

    @Column(name = "inodes_used_percent")
    private double inodesUsedPercent;

    // getters y setters
}
```

### Clase PartitionStat

Estadísticas sobre la partición del disco.

#### Propiedades

- **device**: `String`

  El nombre del dispositivo de la partición.

- **mountPoint**: `String`

  El punto de montaje de la partición.

- **fsType**: `String`

  El tipo de sistema de archivos de la partición.

- **opts**: `String`

  Opciones aplicadas a la partición.

```java
@Entity
@Table(name = "partition_stat")
public class PartitionStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id")
    private Long id;

    @OneToOne
    @JoinColumn(name = "disk_stat_id")
    private DiskStat diskStat;

    @Column(name = "device")
    private String device;

    @Column(name = "mount_point")
    private String mountPoint;

    @Column(name = "fs_type")
    private String fsType;

    @Column(name = "opts")
    private String opts;

    // getters y setters
}
```

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

```java
@Entity
@Table(name = "io_counters_stat")
public class IOCountersStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id")
    private Long id;

    @ManyToOne
    @JoinColumn(name = "disk_stat_id")
    private DiskStat diskStat;

    @Column(name = "read_count")
    private long readCount;

    @Column(name = "merged_read_count")
    private long mergedReadCount;

    @Column(name = "write_count")
    private long writeCount;

    @Column(name = "merged_write_count")
    private long mergedWriteCount;

    @Column(name = "read_bytes")
    private long readBytes;

    @Column(name = "write_bytes")
    private long writeBytes;

    @Column(name = "read_time")
    private long readTime;

    @Column(name = "write_time")
    private long writeTime;

    @Column(name = "iops_in_progress")
    private long iopsInProgress;

    @Column(name = "io_time")
    private long ioTime;

    @Column(name = "weighted_io")
    private long weightedIO;

    @Column(name = "name")
    private String name;

    @Column(name = "serial_number")
    private String serialNumber;

    @Column(name = "label")
    private String label;

    // getters y setters
}
```
