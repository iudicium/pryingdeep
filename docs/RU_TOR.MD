# Контейнер Тор

## Cборка

*Каталог, в котором находится контейнер Docker:*


- ```sh
docker build -t tor .
```

*Корневой каталог:*

- ```sh
docker build -t tor -f build/package/tor/Dockerfile .
```

## Запуск

Это разместит контейнер Docker на localhost, чтобы контейнеры могли взаимодействовать. Подстройте порт соответствующим образом.

```sh
docker run --network host --publish 9050:9050 tor
```
