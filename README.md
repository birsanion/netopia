# Netopia project

Aplicația curentă include:
- endpoint-ul **/payments** primește un request HTTP cu un payload pentru inițierea unei plăți, îl salveaza intr-o baza de date MySQL și trimite un mesaj într-o coadă RabbitMQ
- endpoint-ul **/health** verifică dacă aplicația rulează
- endpoint-ul **/metrics** oferă date despre performanța, utilizarea sau starea internă a aplicației (integrat **Prometheus**)

## Instrucțiuni pentru rularea aplicației local

### Dependențe

- [Docker compose](https://docs.docker.com/compose/install/)


Comandă pentru a rula aplicația local

    $ docker compose up --build

Aplicația va fi accesibilă la adresa [localhost:8888](http://localhost:8888)

Pentru a testa endpoint-urile puteți folosi:

- Documentație (doar pentru **/payments** și **/health**) accesibilă la adresa: [localhost:8070](http://localhost:8070)

- Interfața de administrare a bazei de date accesibilă la adresa: [localhost:9000](http://localhost:9000) (server: db, username: root, password: secret)

- Interfata de administrare RabbitMQ accesibilă la adresa:: [localhost:15672](http://localhost:15672) (user: guest, password: guest)
