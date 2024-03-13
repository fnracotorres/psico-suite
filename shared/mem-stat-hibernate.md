## Documentación del Modelo Hibernate para MemStat

### Clase MemStat

Representa estadísticas de memoria del sistema.

#### Propiedades

- **swapDeviceList**: `List<SwapDevice>`

  Lista de dispositivos de intercambio de memoria.

- **swapMemoryStat**: `SwapMemoryStat`

  Estadísticas de memoria de intercambio.

- **virtualMemoryExStat**: `VirtualMemoryExStat`

  Estadísticas extendidas de memoria virtual.

- **virtualMemoryStat**: `VirtualMemoryStat`

  Estadísticas de memoria virtual.

```java
@Entity
@Table(name = "mem_stat")
public class MemStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToMany(mappedBy = "memStat", cascade = CascadeType.ALL)
    private List<SwapDevice> swapDeviceList;

    @OneToOne(mappedBy = "memStat", cascade = CascadeType.ALL)
    private SwapMemoryStat swapMemoryStat;

    @OneToOne(mappedBy = "memStat", cascade = CascadeType.ALL)
    private VirtualMemoryExStat virtualMemoryExStat;

    @OneToOne(mappedBy = "memStat", cascade = CascadeType.ALL)
    private VirtualMemoryStat virtualMemoryStat;

    // getters y setters
}
```

### Clase SwapDevice

Representa un dispositivo de intercambio de memoria.

#### Propiedades

- **name**: `String`

  Nombre del dispositivo de intercambio.

- **usedBytes**: `long`

  Cantidad de bytes utilizados en el dispositivo de intercambio.

- **freeBytes**: `long`

  Cantidad de bytes libres en el dispositivo de intercambio.

```java
@Entity
@Table(name = "swap_device")
public class SwapDevice {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "mem_stat_id")
    private MemStat memStat;

    @Column(name = "name")
    private String name;

    @Column(name = "used_bytes")
    private long usedBytes;

    @Column(name = "free_bytes")
    private long freeBytes;

    // getters y setters
}
```

### Clase SwapMemoryStat

Representa estadísticas de memoria de intercambio.

#### Propiedades

- **total**: `long`

  Tamaño total de la memoria de intercambio en bytes.

- **used**: `long`

  Cantidad de memoria de intercambio utilizada en bytes.

- **free**: `long`

  Cantidad de memoria de intercambio libre en bytes.

- **usedPercent**: `double`

  Porcentaje de memoria de intercambio utilizada.

- **sin**: `long`

  Cantidad de bytes leídos desde el espacio de intercambio en la última medición.

- **sout**: `long`

  Cantidad de bytes escritos en el espacio de intercambio en la última medición.

- **pgin**: `long`

  Número de páginas de memoria de intercambio leídas desde el disco.

- **pgout**: `long`

  Número de páginas de memoria de intercambio escritas en el disco.

- **pgfault**: `long`

  Número de fallos de página de memoria de intercambio.

- **pgmajfault**: `long`

  Número de fallos mayores de página de memoria de intercambio (específico de Linux).

```java
@Entity
@Table(name = "swap_memory_stat")
public class SwapMemoryStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne
    @JoinColumn(name = "mem_stat_id")
    private MemStat memStat;

    @Column(name = "total")
    private long total;

    @Column(name = "used")
    private long used;

    @Column(name = "free")
    private long free;

    @Column(name = "used_percent")
    private double usedPercent;

    @Column(name = "sin")
    private long sin;

    @Column(name = "sout")
    private long sout;

    @Column(name = "pgin")
    private long pgin;

    @Column(name = "pgout")
    private long pgout;

    @Column(name = "pgfault")
    private long pgfault;

    @Column(name = "pgmajfault")
    private long pgmajfault;

    // getters y setters
}
```

### Clase VirtualMemoryExStat

Representa estadísticas extendidas de memoria virtual.

#### Propiedades

- **activeFile**: `long`

  Tamaño de la memoria activa en archivos.

- **inactiveFile**: `long`

  Tamaño de la memoria inactiva en archivos.

- **activeAnon**: `long`

  Tamaño de la memoria activa anónima.

- **inactiveAnon**: `long`

  Tamaño de la memoria inactiva anónima.

- **unevictable**: `long`

  Tamaño de la memoria no evacuable.

```java
@Entity
@Table(name = "virtual_memory_ex_stat")
public class VirtualMemoryExStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne
    @JoinColumn(name = "mem_stat_id")
    private MemStat memStat;

    @Column(name = "active_file")
    private long activeFile;

    @Column(name = "inactive_file")
    private long inactiveFile;

    @Column(name = "active_anon")
    private long activeAnon;

    @Column(name = "inactive_anon")
    private long inactiveAnon;

    @Column(name = "unevictable")
    private long unevictable;

    // getters y setters
}
```

### Clase VirtualMemoryStat

Representa estadísticas de memoria virtual.

#### Propiedades

- **total**: `long`

  Tamaño total de la memoria RAM en bytes.

- **available**: `long`

  Tamaño de la memoria RAM disponible para asignación de programas en bytes.

- **used**: `long`

  Tamaño de la memoria RAM utilizada por programas en bytes.

- **usedPercent**: `double`

  Porcentaje de la memoria RAM utilizada por programas.

- **free**: `long`

  Tamaño de la memoria RAM libre en bytes.

- **active**: `long`

  Tamaño de la memoria RAM activa en bytes.

- **inactive**: `long`

  Tamaño de la memoria RAM inactiva en bytes.

- **wired**: `long`

  Tamaño de la memoria RAM cableada en bytes.

- **laundry**: `long`

  Tamaño de la memoria RAM en proceso de limpieza en bytes (específico de FreeBSD).

- **buffers**: `long`

  Tamaño de la memoria RAM utilizada como buffers en bytes (específico de Linux).

- **cached**: `long`

  Tamaño de la memoria RAM utilizada como caché en bytes (específico de Linux).

- **writeback**: `long`

  Tamaño de la memoria RAM en proceso de escritura en bytes (específico de Linux).

- **dirty**: `long`

  Tamaño de la memoria RAM con datos sin guardar en bytes (específico de Linux).

- **writebackTmp**: `long`

  Tamaño de la memoria RAM en proceso de escritura temporal en bytes (específico de Linux).

- **shared**: `long`

  Tamaño de la memoria RAM compartida en bytes (específico de Linux).

- **slab**: `long`

  Tamaño de la memoria RAM utilizada para el control de kernel en bytes (específico de Linux).

- **sreclaimable**: `long`

  Tamaño de la memoria RAM recuperable en bytes (específico de Linux).

- **sunreclaim**: `long`

  Tamaño de la memoria RAM no recuperable en bytes (específico de Linux).

- **pageTables**: `long`

  Tamaño de la memoria RAM utilizada por tablas de páginas en bytes (específico de Linux).

- **swapCached**: `long`

  Tamaño de la memoria RAM utilizada como caché de intercambio en bytes (específico de Linux).

- **commitLimit**: `long`

  Límite de compromiso de memoria en bytes (específico de Linux).

- **committedAS**: `long`

  Tamaño total de la memoria virtual comprometida en bytes (específico de Linux).

- **highTotal**: `long`

  Tamaño total de la memoria RAM de alta memoria en bytes (específico de Linux).

- **highFree**: `long`

  Tamaño de la memoria RAM libre de alta memoria en bytes (específico de Linux).

- **lowTotal**: `long`

  Tamaño total de la memoria RAM de baja memoria en bytes (específico de Linux).

- **lowFree**: `long`

  Tamaño de la memoria RAM libre de baja memoria en bytes (específico de Linux).

- **swapTotal**: `long`

  Tamaño total de la memoria de intercambio en bytes (específico de Linux).

- **swapFree**: `long`

  Tamaño de la memoria de intercambio libre en bytes (específico de Linux).

- **mapped**: `long`

  Tamaño de la memoria RAM mapeada en bytes (específico de Linux).

- **vmallocTotal**: `long`

  Tamaño total de la memoria virtual asignada en bytes (específico de Linux).

- **vmallocUsed**: `long`

  Tamaño de la memoria virtual utilizada en bytes (específico de Linux).

- **vmallocChunk**: `long`

  Tamaño del fragmento de memoria virtual en bytes (específico de Linux).

- **hugePagesTotal**: `long`

  Número total de páginas grandes de memoria (específico de Linux).

- **hugePagesFree**: `long`

  Número de páginas grandes de memoria libres (específico de Linux).

- **hugePageSize**: `long`

  Tamaño de página grande de memoria en bytes (específico de Linux).

```java
@Entity
@Table(name = "virtual_memory_stat")
public class VirtualMemoryStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne
    @JoinColumn(name = "mem_stat_id")
    private MemStat memStat;

    @Column(name = "total")
    private long total;

    @Column(name = "available")
    private long available;

    @Column(name = "used")
    private long used;

    @Column(name = "used_percent")
    private double usedPercent;

    @Column(name = "free")
    private long free;

    @Column(name = "active")
    private long active;

    @Column(name = "inactive")
    private long inactive;

    @Column(name = "wired")
    private long wired;

    @Column(name = "laundry")
    private long laundry;

    @Column(name = "buffers")
    private long buffers;

    @Column(name = "cached")
    private long cached;

    @Column(name = "writeback")
    private long writeback

    @Column(name = "dirty")
    private long dirty;

    @Column(name = "writebacktmp")
    private long writebackTmp;

    @Column(name = "shared")
    private long shared;

    @Column(name = "slab")
    private long slab;

    @Column(name = "sreclaimable")
    private long sreclaimable;

    @Column(name = "sunreclaim")
    private long sunreclaim;

    @Column(name = "pagetables")
    private long pageTables;

    @Column(name = "swapcached")
    private long swapCached;

    @Column(name = "commitlimit")
    private long commitLimit;

    @Column(name = "committedas")
    private long committedAS;

    @Column(name = "hightotal")
    private long highTotal;

    @Column(name = "highfree")
    private long highFree;

    @Column(name = "lowtotal")
    private long lowTotal;

    @Column(name = "lowfree")
    private long lowFree;

    @Column(name = "swaptotal")
    private long swapTotal;

    @Column(name = "swapfree")
    private long swapFree;

    @Column(name = "mapped")
    private long mapped;

    @Column(name = "vmalloctotal")
    private long vmallocTotal;

    @Column(name = "vmallocused")
    private long vmallocUsed;

    @Column(name = "vmallocchunk")
    private long vmallocChunk;

    @Column(name = "hugepagestotal")
    private long hugePagesTotal;

    @Column(name = "hugepagesfree")
    private long hugePagesFree;

    @Column(name = "hugepagesize")
    private long hugePageSize;

    // getters y setters
}
```
