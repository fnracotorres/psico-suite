## Documentación del Modelo Hibernate para ProcessStat

### Clase ProcessStat

Representa estadísticas de un proceso.

#### Propiedades

- **memoryInfoExStat**: `MemoryInfoExStat`

  Estadísticas extendidas de información de memoria del proceso.

- **memoryInfoStat**: `MemoryInfoStat`

  Estadísticas de información de memoria del proceso.

- **numCtxSwitchesStat**: `NumCtxSwitchesStat`

  Estadísticas de cambios de contexto del proceso.

```java
@Entity
@Table(name = "process_stat")
public class ProcessStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne(mappedBy = "processStat", cascade = CascadeType.ALL)
    private MemoryInfoExStat memoryInfoExStat;

    @OneToOne(mappedBy = "processStat", cascade = CascadeType.ALL)
    private MemoryInfoStat memoryInfoStat;

    @OneToOne(mappedBy = "processStat", cascade = CascadeType.ALL)
    private NumCtxSwitchesStat numCtxSwitchesStat;

    // getters y setters
}
```

### Clase MemoryInfoExStat

Representa estadísticas extendidas de información de memoria.

#### Propiedades

- **RSS**: `long`

  Tamaño de la porción de memoria residente (RSS) en bytes.

- **VMS**: `long`

  Tamaño de la memoria virtual (VMS) en bytes.

- **Shared**: `long`

  Tamaño de la memoria compartida en bytes.

- **Text**: `long`

  Tamaño de la memoria de texto en bytes.

- **Lib**: `long`

  Tamaño de la memoria de la biblioteca en bytes.

- **Data**: `long`

  Tamaño de la memoria de datos en bytes.

- **Dirty**: `long`

  Tamaño de la memoria sucia en bytes.

```java
@Entity
@Table(name = "memory_info_ex_stat")
public class MemoryInfoExStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "rss")
    private long rss;

    @Column(name = "vms")
    private long vms;

    @Column(name = "shared")
    private long shared;

    @Column(name = "text")
    private long text;

    @Column(name = "lib")
    private long lib;

    @Column(name = "data")
    private long data;

    @Column(name = "dirty")
    private long dirty;

    // getters y setters
}
```

### Clase MemoryInfoStat

Representa estadísticas de información de memoria.

#### Propiedades

- **rss**: `long`

  Tamaño de la memoria residente (RSS) en bytes.

- **vms**: `long`

  Tamaño de la memoria virtual (VMS) en bytes.

- **hwm**: `long`

  Tamaño máximo de memoria residente (HWM) en bytes.

- **data**: `long`

  Tamaño del segmento de datos en bytes.

- **stack**: `long`

  Tamaño de la pila en bytes.

- **locked**: `long`

  Tamaño de memoria bloqueada en bytes.

- **swap**: `long`

  Tamaño de la memoria swap en bytes.

```java
@Entity
@Table(name = "memory_info_stat")
public class MemoryInfoStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "rss")
    private long rss;

    @Column(name = "vms")
    private long vms;

    @Column(name = "hwm")
    private long hwm;

    @Column(name = "data")
    private long data;

    @Column(name = "stack")
    private long stack;

    @Column(name = "locked")
    private long locked;

    @Column(name = "swap")
    private long swap;

    // getters y setters
}
```

### Clase NumCtxSwitchesStat

Representa estadísticas de cambios de contexto.

#### Propiedades

- **voluntary**: `long`

  Número de cambios de contexto voluntarios.

- **involuntary**: `long`

  Número de cambios de contexto involuntarios.

```java
@Entity
@Table(name = "num_ctx_switches_stat")
public class NumCtxSwitchesStat {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "voluntary")
    private long voluntary;

    @Column(name = "involuntary")
    private long involuntary;

    // getters y setters
}
```
