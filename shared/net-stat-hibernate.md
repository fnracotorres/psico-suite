## Documentación del Modelo Hibernate para NetStat

### Clase NetStat

Representa estadísticas de red.

#### Propiedades

- **connectionStatList**: `List<ConnectionStat>`

  Lista de estadísticas de conexiones de red.

- **conntrackStatListAggregation**: `List<ConntrackStat>`

  Lista de estadísticas de seguimiento de conexiones agregadas.

- **conntrackStatListPerCPU**: `List<ConntrackStat>`

  Lista de estadísticas de seguimiento de conexiones por CPU.

- **filterStatList**: `List<FilterStat>`

  Lista de estadísticas de filtro de red.

- **ioCountersStatListAggregation**: `List<IOCountersStat>`

  Lista de estadísticas de contadores de E/S agregadas.

- **ioCountersStatListPerNIC**: `List<IOCountersStat>`

  Lista de estadísticas de contadores de E/S por interfaz de red.

- **interfaceStatList**: `List<InterfaceStat>`

  Lista de estadísticas de interfaz de red.

- **protoCountersStatList**: `List<ProtoCountersStat>`

  Lista de estadísticas de contadores de protocolo de red.

```java
@Entity
@Table(name = "net_stat")
public class NetStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<ConnectionStat> connectionStatList;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<ConntrackStat> conntrackStatListAggregation;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<ConntrackStat> conntrackStatListPerCPU;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<FilterStat> filterStatList;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<IOCountersStat> ioCountersStatListAggregation;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<IOCountersStat> ioCountersStatListPerNIC;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<InterfaceStat> interfaceStatList;

    @OneToMany(mappedBy = "netStat", cascade = CascadeType.ALL)
    private List<ProtoCountersStat> protoCountersStatList;

    // getters y setters
}
```

### Clase ConnectionStat

Representa estadísticas de conexión de red.

#### Propiedades

- **fd**: `int`

  Descriptor de archivo asociado con la conexión.

- **family**: `int`

  Familia de protocolo de la conexión.

- **type**: `int`

  Tipo de conexión.

- **laddr**: `Addr`

  Dirección local de la conexión.

- **raddr**: `Addr`

  Dirección remota de la conexión.

- **status**: `String`

  Estado de la conexión.

- **uids**: `List<Integer>`

  Lista de identificadores de usuario asociados con la conexión.

- **pid**: `int`

  Identificador del proceso asociado con la conexión.

```java
@Entity
@Table(name = "connection_stat")
public class ConnectionStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "fd")
    private int fd;

    @Column(name = "family")
    private int family;

    @Column(name = "type")
    private int type;

    @OneToOne(mappedBy = "connectionStat", cascade = CascadeType.ALL)
    private Addr laddr;

    @OneToOne(mappedBy = "connectionStat", cascade = CascadeType.ALL)
    private Addr raddr;

    @Column(name = "status")
    private String status;

    @ElementCollection
    @CollectionTable(name = "connection_stat_uids")
    @Column(name = "uid")
    private List<Integer> uids;

    @Column(name = "pid")
    private int pid;

    // getters y setters
}
```

### Clase Addr

Representa una dirección IP y un puerto.

#### Propiedades

- **ip**: `String`

  Dirección IP.

- **port**: `int`

  Puerto.

```java
@Entity
@Table(name = "addr")
public class Addr {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "ip")
    private String ip;

    @Column(name = "port")
    private int port;

    // getters y setters
}
```

### Clase ConntrackStat

Representa estadísticas de seguimiento de conexiones.

#### Propiedades

- **entries**: `int`

  Número de entradas en la tabla de seguimiento de conexiones.

- **searched**: `int`

  Número de búsquedas realizadas en la tabla de seguimiento de conexiones.

- **found**: `int`

  Número de búsquedas que resultaron exitosas.

- **new**: `int`

  Número de entradas nuevas agregadas que no se esperaban antes.

- **invalid**: `int`

  Número de paquetes vistos que no pueden ser rastreados.

- **ignore**: `int`

  Paquetes vistos que ya están conectados a una entrada.

- **delete**: `int`

  Número de entradas que fueron eliminadas.

- **deleteList**: `int`

  Número de entradas que fueron agregadas a la lista de eliminación.

- **insert**: `int`

  Número de entradas insertadas en la lista.

- **insertFailed**: `int`

  Número de intentos de inserción fallidos (la misma entrada ya existe).

- **drop**: `int`

  Número de paquetes eliminados debido a un error de seguimiento de conexiones.

- **earlyDrop**: `int`

  Entradas eliminadas para dejar espacio para nuevas, si se alcanza el tamaño máximo.

- **icmpError**: `int`

  Subconjunto de inválido. Paquetes que no se pueden rastrear debido a un error.

- **expectNew**: `int`

  Entradas agregadas después de que ya estuviera presente una expectativa.

- **expectCreate**: `int`

  Expectativas agregadas.

- **expectDelete**: `int`

  Expectativas eliminadas.

- **searchRestart**: `int`

  Búsquedas reiniciadas debido a cambios en el tamaño de la tabla.

```java
@Entity
@Table(name = "conntrack_stat")
public class ConntrackStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "entries")
    private int entries;

    @Column(name = "searched")
    private int searched;

    @Column(name = "found")
    private int found;

    @Column(name = "new")
    private int newEntries;

    @Column(name = "invalid")
    private int invalid;

    @Column(name = "ignore")
    private int ignore;

    @Column(name = "delete")
    private int deleteEntries;

    @Column(name = "delete_list")
    private int deleteList;

    @Column(name = "insert")
    private int insert;

    @Column(name = "insert_failed")
    private int insertFailed;

    @Column(name = "drop")
    private int drop;

    @Column(name = "early_drop")
    private int earlyDrop;

    @Column(name = "icmp_error")
    private int icmpError;

    @Column(name = "expect_new")
    private int expectNew;

    @Column(name = "expect_create")
    private int expectCreate;

    @Column(name = "expect_delete")
    private int expectDelete;

    @Column(name = "search_restart")
    private int searchRestart;

    // getters y setters
}
```

### Clase FilterStat

Representa estadísticas de filtro de red.

#### Propiedades

- **connTrackCount**: `long`

  Conteo actual de conexiones rastreadas por el filtro.

- **connTrackMax**: `long`

  Máximo de conexiones rastreadas admitidas por el filtro.

```java
@Entity
@Table(name = "filter_stat")
public class FilterStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "conntrack_count")
    private long connTrackCount;

    @Column(name = "conntrack_max")
    private long connTrackMax;

    // getters y setters
}
```

### Clase IOCountersStat

Representa estadísticas de contadores de E/S.

#### Propiedades

- **name**: `String`

  Nombre de la interfaz.

- **bytesSent**: `long`

  Número de bytes enviados.

- **bytesRecv**: `long`

  Número de bytes recibidos.

- **packetsSent**: `long`

  Número de paquetes enviados.

- **packetsRecv**: `long`

  Número de paquetes recibidos.

- **errin**: `long`

  Número total de errores mientras se reciben datos.

- **errout**: `long`

  Número total de errores mientras se envían datos.

- **dropin**: `long`

  Número total de paquetes entrantes que fueron descartados.

- **dropout**: `long`

  Número total de paquetes salientes que fueron descartados.

- **fifoin**: `long`

  Número total de errores de FIFO mientras se reciben datos.

- **fifoout**: `long`

  Número total de errores de FIFO mientras se envían datos.

```java
@Entity
@Table(name = "io_counters_stat")
public class IOCountersStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "name")
    private String name;

    @Column(name = "bytes_sent")
    private long bytesSent;

    @Column(name = "bytes_recv")
    private long bytesRecv;

    @Column(name = "packets_sent")
    private long packetsSent;

    @Column(name = "packets_recv")
    private long packetsRecv;

    @Column(name = "errin")
    private long errin;

    @Column(name = "errout")
    private long errout;

    @Column(name = "dropin")
    private long dropin;

    @Column(name = "dropout")
    private long dropout;

    @Column(name = "fifoin")
    private long fifoin;

    @Column(name = "fifoout")
    private long fifoout;

    // getters y setters
}
```

### Clase InterfaceStat

Representa estadísticas de una interfaz de red.

#### Propiedades

- **index**: `int`

  Índice de la interfaz.

- **mtu**: `int`

  Unidad máxima de transmisión.

- **name**: `String`

  Nombre de la interfaz, por ejemplo, "en0", "lo0", "eth0.100".

- **hardwareAddr**: `String`

  Dirección MAC de la interfaz en forma IEEE MAC-48, EUI-48 y EUI-64.

- **flags**: `List<String>`

  Lista de banderas de la interfaz, por ejemplo, FlagUp, FlagLoopback, FlagMulticast.

- **addrs**: `List<InterfaceAddr>`

  Lista de direcciones asociadas con la interfaz.

```java
@Entity
@Table(name = "interface_stat")
public class InterfaceStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "index")
    private int index;

    @Column(name = "mtu")
    private int mtu;

    @Column(name = "name")
    private String name;

    @Column(name = "hardware_addr")
    private String hardwareAddr;

    @ElementCollection
    @CollectionTable(name = "interface_stat_flags")
    @Column(name = "flag")
    private List<String> flags;

    @OneToMany(mappedBy = "interfaceStat", cascade = CascadeType.ALL)
    private List<InterfaceAddr> addrs;

    // getters y setters
}
```

### Clase InterfaceAddr

Representa una dirección asociada con una interfaz de red.

#### Propiedades

- **addr**: `String`

  Dirección IP asociada con la interfaz.

```java
@Entity
@Table(name = "interface_addr")
public class InterfaceAddr {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "addr")
    private String addr;

    @ManyToOne
    @JoinColumn(name = "interface_stat_id")
    private InterfaceStat interfaceStat;

    // getters y setters
}
```

### Clase ProtoCountersStat

Representa estadísticas de contadores de protocolo.

#### Propiedades

- **protocol**: `String`

  Protocolo asociado con las estadísticas.

- **stats**: `Map<String, Long>`

  Estadísticas específicas del protocolo, representadas como un mapa donde la clave es el nombre del contador y el valor es el valor del contador.

```java
@Entity
@Table(name = "proto_counters_stat")
public class ProtoCountersStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "protocol")
    private String protocol;

    @ElementCollection
    @CollectionTable(name = "proto_counters_stat_stats")
    @MapKeyColumn(name = "counter_name")
    @Column(name = "counter_value")
    private Map<String, Long> stats;

    // getters y setters
}
```
