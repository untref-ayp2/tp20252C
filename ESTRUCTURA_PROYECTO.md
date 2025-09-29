# Estructura del Proyecto - Red Social de Nicho

## 📁 Organización Completa del Espacio de Trabajo
A continuación se detalla la estructura completa del proyecto, incluyendo código fuente, documentación, casos de prueba y organización del trabajo en equipo.

Esta estructura está diseñada para facilitar la colaboración entre los integrantes del grupo, permitiendo que cada uno pueda trabajar en su módulo asignado y contribuir al proyecto de manera organizada.

Pueden modificarla según sus necesidades, pero se recomienda mantener una estructura clara y coherente.

```
tp20252C/
├── README.md                    # Descripción general del trabajo práctico
├── .gitignore                   # Archivos ignorados por git
│
├── 📁 src/                      # Código fuente principal
│   ├── README.md               
│   ├── usuarios/               # Módulo de gestión de usuarios
│   ├── relaciones/             # Módulo de seguimientos y relaciones
│   ├── comunidades/            # Módulo de grupos y comunidades
│   ├── publicaciones/          # Módulo de contenidos y publicaciones
│   ├── recomendaciones/        # Algoritmos de recomendación
│   ├── estadisticas/           # Cálculo de métricas
│   ├── estructuras/            # Estructuras de datos personalizadas
│   ├── interfaz/               # Interfaz de usuario/consola/API
│   └── utils/                  # Utilidades y funciones auxiliares
│
├── 📁 docs/                     # Documentación técnica
│   ├── README.md
│   ├── analisis_complejidad/   # Análisis de algoritmos
│   ├── diagramas/              # Diagramas UML y arquitectura
│   ├── presentacion/           # Material para presentación oral
│   └── informes/               # Informes técnicos
│
├── 📁 tests/                    # Casos de prueba
│   ├── README.md
│   ├── casos_prueba/           # Casos de prueba documentados (mín. 10)
│   ├── unitarios/              # Pruebas unitarias por módulo
│   └── integracion/            # Pruebas de integración
│
├── 📁 equipo/                   # Organización del trabajo grupal
│   ├── README.md               # Roles y responsabilidades
│   ├── roles/                  # Definición de roles específicos
│   ├── reuniones/              # Actas de reuniones
│   │   └── plantilla_acta.md
│   └── planificacion/          # Cronogramas y planificación
│       └── cronograma.md
│
├── 📁 examples/                 # Ejemplos de uso y demostraciones
├── └── 📁 build/               # Archivos compilados y ejecutables
```

## 🎯 Módulos del Sistema

### Módulos Principales

1. **`usuarios/`** - Registro, eliminación, consultas de usuarios
2. **`relaciones/`** - Seguimientos, seguidores, consultas mutuas  
3. **`comunidades/`** - Grupos, fusión automática con union-find
4. **`publicaciones/`** - Contenidos, tags, búsquedas
5. **`recomendaciones/`** - Sugerencias de usuarios y grupos
6. **`estadisticas/`** - Métricas y estadísticas del sistema

### Módulos de Soporte

7. **`estructuras/`** - Listas, árboles, hash, montículos, union-find
8. **`interfaz/`** - Consola/API para interactuar con el sistema
9. **`utils/`** - Funciones auxiliares y utilidades

## 📋 Entregables Organizados

- **📄 Código fuente**: Organizado por módulos en `src/`
- **📚 Documentación técnica**: En `docs/` con análisis de complejidad
- **🧪 Casos de prueba**: Mínimo 10 casos en `tests/`
- **🎤 Presentación oral**: Material en `docs/presentacion/`

## 👥 Trabajo en Equipo

- **Roles definidos**: Cada integrante tiene responsabilidades específicas
- **Seguimiento**: Plantillas para reuniones y planificación
- **Colaboración**: Estructura que facilita el trabajo distribuido
- **Commits identificables**: Cada integrante puede trabajar en su módulo

## 🏗️ Próximos Pasos

1. **Asignar roles** específicos a cada integrante del grupo
2. **Completar información** en `equipo/README.md`
3. **Elegir lenguaje** de programación y configurar herramientas
4. **Definir cronograma** detallado en `equipo/planificacion/`
5. **Comenzar implementación** con módulos básicos

---

**✅ El entorno está configurado y listo para el desarrollo colaborativo del trabajo práctico grupal.**