## Documentación del Modelo Hibernate para CPUStat

### Clase CPUStat

Representa estadísticas de la CPU.

#### Propiedades

- **infoStatList**: `List<InfoStat>`

  Lista de información sobre la CPU.

- **timesStatList**: `List<TimesStat>`

  Lista de estadísticas de tiempo de la CPU.

```java
@Entity
@Table(name = "cpu_stat")
public class CPUStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToMany(mappedBy = "cpuStat", cascade = CascadeType.ALL)
    private List<InfoStat> infoStatList;

    @OneToMany(mappedBy = "cpuStat", cascade = CascadeType.ALL)
    private List<TimesStat> timesStatList;

    // getters y setters
}
```

### Clase InfoStat

Representa información sobre la CPU.

#### Propiedades

- **CPU**: `int`

  El número de CPU.

- **VendorID**: `String`

  El ID del fabricante de la CPU.

- **Family**: `String`

  La familia de la CPU.

- **Model**: `String`

  El modelo de la CPU.

- **Stepping**: `int`

  El stepping de la CPU.

- **PhysicalID**: `String`

  El ID físico de la CPU.

- **CoreID**: `String`

  El ID del núcleo de la CPU.

- **Cores**: `int`

  El número de núcleos de la CPU.

- **ModelName**: `String`

  El nombre del modelo de la CPU.

- **Mhz**: `double`

  La frecuencia de la CPU en MHz.

- **CacheSize**: `int`

  El tamaño de la caché de la CPU.

- **Flags**: `List<String>`

  Una lista de flags de la CPU.

- **Microcode**: `String`

  El microcódigo de la CPU.

```java
@Entity
@Table(name = "info_stat")
public class InfoStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "cpu")
    private int CPU;

    @Column(name = "vendor_id")
    private String vendorID;

    @Column(name = "family")
    private String family;

    @Column(name = "model")
    private String model;

    @Column(name = "stepping")
    private int stepping;

    @Column(name = "physical_id")
    private String physicalID;

    @Column(name = "core_id")
    private String coreID;

    @Column(name = "cores")
    private int cores;

    @Column(name = "model_name")
    private String modelName;

    @Column(name = "mhz")
    private double mhz;

    @Column(name = "cache_size")
    private int cacheSize;

    @ElementCollection
    @CollectionTable(name = "info_stat_flags", joinColumns = @JoinColumn(name = "info_stat_id"))
    @Column(name = "flag")
    private List<String> flags;

    @Column(name = "microcode")
    private String microcode;

    // getters y setters
}
```

### Clase TimesStat

Representa estadísticas de tiempo de la CPU.

#### Propiedades

- **CPU**: `String`

  El identificador de la CPU.

- **User**: `double`

  Tiempo en modo de usuario.

- **System**: `double`

  Tiempo en modo sistema.

- **Idle**: `double`

  Tiempo en estado de inactividad.

- **Nice**: `double`

  Tiempo en modo "nice".

- **Iowait**: `double`

  Tiempo de espera de E/S.

- **Irq**: `double`

  Tiempo en servicio de interrupción.

- **Softirq**: `double`

  Tiempo en servicio de interrupción suave.

- **Steal**: `double`

  Tiempo robado de la máquina virtual.

- **Guest**: `double`

  Tiempo en modo "guest".

- **GuestNice**: `double`

  Tiempo en modo "guest nice".

```java
@Entity
@Table(name = "times_stat")
public class TimesStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "cpu")
    private String CPU;

    @Column(name = "user")
    private double user;

    @Column(name = "system")
    private double system;

    @Column(name = "idle")
    private double idle;

    @Column(name = "nice")
    private double nice;

    @Column(name = "iowait")
    private double iowait;

    @Column(name = "irq")
    private double irq;

    @Column(name = "softirq")
    private double softirq;

    @Column(name = "steal")
    private double steal;

    @Column(name = "guest")
    private double guest;

    @Column(name = "guest_nice")
    private double guestNice;

    // getters y setters
}
```
