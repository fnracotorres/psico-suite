## Documentación del Modelo Hibernate para LoadStat

### Clase LoadStat

Representa estadísticas de carga del sistema.

#### Propiedades

- **avgStat**: `AvgStat`

  Estadísticas promedio de carga.

- **miscStat**: `MiscStat`

  Otras estadísticas de carga.

```java
@Entity
@Table(name = "load_stat")
public class LoadStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne(mappedBy = "loadStat", cascade = CascadeType.ALL)
    private AvgStat avgStat;

    @OneToOne(mappedBy = "loadStat", cascade = CascadeType.ALL)
    private MiscStat miscStat;

    // getters y setters
}
```

### Clase AvgStat

Representa estadísticas promedio de carga del sistema.

#### Propiedades

- **load1**: `double`

  Promedio de carga en el último minuto.

- **load5**: `double`

  Promedio de carga en los últimos cinco minutos.

- **load15**: `double`

  Promedio de carga en los últimos quince minutos.

```java
@Entity
@Table(name = "avg_stat")
public class AvgStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "load1")
    private double load1;

    @Column(name = "load5")
    private double load5;

    @Column(name = "load15")
    private double load15;

    // getters y setters
}
```

### Clase MiscStat

Representa estadísticas misceláneas del sistema.

#### Propiedades

- **procsTotal**: `int`

  Número total de procesos en el sistema.

- **procsCreated**: `int`

  Número de procesos creados desde el inicio del sistema.

- **procsRunning**: `int`

  Número de procesos en ejecución.

- **procsBlocked**: `int`

  Número de procesos bloqueados.

- **ctxt**: `int`

  Número de cambios de contexto realizados desde el inicio del sistema.

```java
@Entity
@Table(name = "misc_stat")
public class MiscStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "procs_total")
    private int procsTotal;

    @Column(name = "procs_created")
    private int procsCreated;

    @Column(name = "procs_running")
    private int procsRunning;

    @Column(name = "procs_blocked")
    private int procsBlocked;

    @Column(name = "ctxt")
    private int ctxt;

    // getters y setters
}
```
