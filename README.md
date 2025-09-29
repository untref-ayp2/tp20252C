# Trabajo Práctico Grupal

El trabajo práctico consiste en desarrollar la aplicación que se describe más adelante. Los grupos deben ser de 4 alumnos.

Cada grupo tendrá asignado un docente que los guiará en el desarrollo del trabajo. El docente asignado será el encargado de evaluar el trabajo práctico. Deben realizar entregas parciales del trabajo práctico y realizar una presentación oral del mismo.

Cada estudiante debe asumir un rol específico.

Todos los integrantes deben participar en el informe y en la presentación. Es importante que cada integrante realice commits identificables en el repositorio del proyecto.

La fechas de las presentaciones orales previstas son:

Comisión Lunes y Miercoles
: 5 y 10 de noviembre de 2025

Comisión Martes y Jueves
: 6 y 11 de noviembre de 2025

Para poder acceder a la presentación oral deben tener el trabajo práctico aprobado por el docente a cargo.

## Red Social de Nicho

### Objetivo general

Desarrollar una red social de nicho (por ejemplo para estudiantes de ingeniería en computación), que gestione usuarios, contenidos, relaciones entre usuarios y grupos, y que permita ejecutar ciertas operaciones complejas (búsquedas, recomendaciones, uniones de comunidades, estadísticas). El proyecto debe integrar diversas estructuras de datos vistas en clase, y además incluir al menos una estructura adicional no abordada (por ejemplo, conjuntos disjuntos / union-find). Debe priorizarse el diseño algorítmico (eficiencia), la corrección y claridad (documentación) y una presentación oral clara del trabajo.

### Requisitos mínimos del sistema

1. **Módulo de usuarios**
   - Registro de usuarios (con atributos como identificador único, nombre, email, fecha de alta, intereses)  
   - Eliminación / desactivación de usuarios  
   - Consultas: obtener datos de un usuario, listar usuarios por criterio

2. **Módulo de relaciones / amistad / seguidores**
   - Cada usuario puede “seguir” a otros usuarios (relaciones dirigidas)  
   - Soporte para eliminar seguidos / dejar de seguir  
   - Consulta de seguidores / seguidos de un usuario  
   - Consulta mutua: ¿quién me sigue que yo también sigo?

3. **Comunidades / grupos temáticos**
   - Crear grupos (identificador, nombre, tema, listado de miembros)  
   - Un usuario puede unirse / abandonar grupos  
   - Fusión automática de grupos (cuando dos grupos de intereses similares se detectan como “cercanos”) usando **conjuntos disjuntos (union-find)**  
   - Consulta de miembros de un grupo, grupos a los que participa un usuario

4. **Publicaciones / contenidos**
   - Los usuarios pueden publicar contenidos (texto corto), con etiquetas (“tags”)  
   - Mantenimiento de índice de tags → publicaciones que contienen cada tag  
   - Búsqueda de publicaciones por tag, por usuario, por rango temporal

5. **Recomendaciones / sugerencias**
   - Recomendar usuarios a seguir: “amigos de amigos”  
   - Recomendar grupos para un usuario: por coincidencia de tags  
   - Estas recomendaciones deben calcularse eficientemente

6. **Estadísticas y métricas**
   - Cantidad de publicaciones por usuario  
   - Usuarios con más seguidores  
   - Grupos más activos  
   - Distribución de tamaño de grupos (histograma)  

7. **Interfaz mínima / consola / API**
   - Interfaz de consola o API simple  
   - Documentación de los comandos con ejemplos de uso


### Estructuras de datos y técnicas que deben emplearse

- Listas enlazadas  
- Pilas o colas  
- Tablas hash / diccionarios  
- Árboles binarios de búsqueda / árboles balanceados  
- Montículo / cola de prioridad  
- Conjuntos disjuntos (union-find)  
- Alguna técnica algorítmica: recursividad, división y conquista, programación dinámica, greedy o backtracking  
- Análisis de complejidad de las operaciones principales

### Entregables

1. Código fuente organizado y comentado  
2. Documentación técnica (README, análisis de complejidad)  
3. Casos de prueba (mínimo 10)  
4. Presentación oral (15 a 20 minutos)  

### Criterios de evaluación

| Aspecto                           | Peso |
| --------------------------------- | ---- |
| Funcionalidad completa y correcta | 30 % |
| Uso de estructuras de datos       | 20 % |
| Diseño algorítmico y análisis     | 15 % |
| Calidad del código                | 10 % |
| Casos de prueba y robustez        | 10 % |
| Presentación oral                 | 10 % |
| Innovación / extras               | 5 %  |


### Ampliaciones opcionales

- “Me gusta” en publicaciones y ranking por popularidad  
- “Feed” personalizado  
- Bloqueo de usuarios  
- Visualización gráfica o interfaz web  
- Persistencia de datos en archivos 