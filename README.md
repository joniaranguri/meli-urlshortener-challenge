# meli-urlshortener-challenge

This repo contains the design and implementation for MercadoLibre Backend Challenge.

### Statement (in Spanish)

Necesitamos hacer un urlshortener tipo goo.gl o bitly para publicar promociones en twitter.
Armar la arquitectura de la solución que satisfaga los siguientes requisitos:

- Las urls tienen una vigencia indefinida.
- La plataforma debe soportar picos de tráfico alrededor de 1M RPM.
- Debemos contar con estadísticas de acceso nearly real time.
- Las URLs publicadas deben poder ser administradas:
- Habilitarse o deshabilitarse
- Modificar la url de destino (cualquiera de sus componentes)
- A su vez la resolución de urls debe realizarse lo más rápido y con el menor costo
  posible.
- La funcionalidad debe ser operativa en un 99,98% del tiempo (uptime)

A tener en cuenta:

- Planteo de distintos componentes explicando su responsabilidad y el por qué de su
  inclusión.
- Explicación de la infraestructura, herramientas/tecnologías preexistentes y cuales
  son los motivos de la elección de cada una (por ejemplo si se incluye una ddbb en
  particular, comentar si se evaluaron otras alternativas y cuál fue el racional de la
  decisión final).
- Es deseable incluir gráficos y explicación breve escrita para que la propuesta sea
  correctamente entendida.
- Es deseable entregar una API Rest funcional corriendo en algún cloud público, con
  una capacidad aproximada de 5000 rpm (verificable) a modo de demo (no hace falta
  gastar dinero en la prueba, mockear lo que no pueda obtenerse gratis).
- El código debe ser compartido a través de un repositorio o bien en un zip.