## Documentación del Modelo Hibernate para HostStat

### Clase HostStat

Representa estadísticas del host.

#### Propiedades

- **infoStat**: `InfoStat`

  Información sobre el host.

- **temperatureStatList**: `List<TemperatureStat>`

  Lista de estadísticas de temperatura del host.

- **userStatList**: `List<UserStat>`

  Lista de estadísticas de usuario del host.

```java
@Entity
@Table(name = "host_stat")
public class HostStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne(mappedBy = "hostStat", cascade = CascadeType.ALL)
    private InfoStat infoStat;

    @OneToMany(mappedBy = "hostStat", cascade = CascadeType.ALL)
    private List<TemperatureStat> temperatureStatList;

    @OneToMany(mappedBy = "hostStat", cascade = CascadeType.ALL)
    private List<UserStat> userStatList;

    // getters y setters
}
```

### Clase InfoStat

Representa información del sistema.

#### Propiedades

- **hostname**: `String`

  El nombre del host.

- **uptime**: `long`

  Tiempo de actividad del sistema en segundos.

- **bootTime**: `long`

  Tiempo de arranque del sistema en segundos desde la época de Unix.

- **procs**: `long`

  Número de procesos en ejecución.

- **os**: `String`

  El nombre del sistema operativo.

- **platform**: `String`

  La distribución o versión específica del sistema operativo.

- **platformFamily**: `String`

  La familia de plataformas del sistema operativo.

- **platformVersion**: `String`

  La versión del sistema operativo.

- **kernelVersion**: `String`

  La versión del kernel del sistema operativo.

- **kernelArch**: `String`

  La arquitectura del kernel del sistema operativo.

- **virtualizationSystem**: `String`

  El sistema de virtualización en uso.

- **virtualizationRole**: `String`

  El rol del sistema de virtualización (anfitrión o invitado).

- **hostID**: `String`

  El identificador único del host.

```java
@Entity
@Table(name = "info_stat")
public class InfoStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "hostname")
    private String hostname;

    @Column(name = "uptime")
    private long uptime;

    @Column(name = "boot_time")
    private long bootTime;

    @Column(name = "procs")
    private long procs;

    @Column(name = "os")
    private String os;

    @Column(name = "platform")
    private String platform;

    @Column(name = "platform_family")
    private String platformFamily;

    @Column(name = "platform_version")
    private String platformVersion;

    @Column(name = "kernel_version")
    private String kernelVersion;

    @Column(name = "kernel_arch")
    private String kernelArch;

    @Column(name = "virtualization_system")
    private String virtualizationSystem;

    @Column(name = "virtualization_role")
    private String virtualizationRole;

    @Column(name = "host_id")
    private String hostID;

    // getters y setters
}
```

### Clase TemperatureStat

Representa estadísticas de temperatura.

#### Propiedades

- **sensorKey**: `String`

  Clave del sensor de temperatura.

- **temperature**: `double`

  Temperatura del sensor en grados Celsius.

```java
@Entity
@Table(name = "temperature_stat")
public class TemperatureStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "sensor_key")
    private String sensorKey;

    @Column(name = "temperature")
    private double temperature;

    // getters y setters
}
```

### Clase UserStat

Representa estadísticas de usuario.

#### Propiedades

- **user**: `String`

  Nombre de usuario.

- **terminal**: `String`

  Terminal donde se inició la sesión.

- **host**: `String`

  Host donde se inició la sesión.

- **started**: `int`

  Marca de tiempo en que se inició la sesión.

```java
@Entity
@Table(name = "user_stat")
public class UserStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "user")
    private String user;

    @Column(name = "terminal")
    private String terminal;

    @Column(name = "host")
    private String host;

    @Column(name = "started")
    private int started;

    // getters y setters
}
```
