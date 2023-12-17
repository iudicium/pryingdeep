<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/iudicium/pryingdeep/web/static/logo.png">
    <img src="../web/static/logo.png" alt="Logo" width="350" height="200">
  </a>

  <h3 align="center">Prying Deep</h3>

  <p align="center">
    <b>OSINT инструмент для сбора информации с тёмного и обычного интернета.
    <br />
    <br />
    <a href="https://github.com/iudicium/pryingdeep/issues">Сообщить об ошибке</a>
    <a href="https://github.com/iudicium/pryingdeep/issues">Предложить улучшение</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Оглавление</summary>
  <ol>
    <li>
      <a href="#about-the-project">О проекте</a>
    </li>
    <li>
      <a href="#getting-started">Начало</a>
      <ul>
        <li><a href="#prerequisites">Предварительные требования</a></li>
        <li><a href="#manual-installation">Ручная установка</a></li>
        <li><a href="#binary-installation">Бинарная установка</a></li>
      </ul>
    </li>
    <li>
      <a href="#docker">Докер</a>
      <ul>
        <li><a href="#tor">Тор</a></li>
      </ul>
    </li>
    <li><a href="#usage">Использование</a></li>
    <li><a href="#roadmap">План развития</a></li>
    <li><a href="#contributing">Участие</a></li>
    <li><a href="#license">Лицензия</a></li>
    <li><a href="#preview">Предварительный просмотр</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Я видел много проектов на Гитхабе, которые тоже поисковые роботы для "черного интернета",
но большинство из них не соответствовали *моим* стандартам OSINT в глубокой паутине.
И я принял решение создать мой OSINT инструмент для глубокой паутины.


Этот инструмент служит напоминанием, что всегда следует соблюдать лучшие практики OPSEC(Операционная Безопасность) в черном интернете.


<u>Автор этого проекта не ответственность за любой возможный ущерб, причиненный использованием этого инструмента.</u>

## Что делает этот инструмент?


*Prying Deep Ползет по платформам тёмной/обычной сети и извлекает максимальное количество информации эффективно*


---

<!-- GETTING STARTED -->
## Getting Started


### Prerequisites


Прежде чем использовать наш OSINT инструмент, убедитесь, что у вас установлены следующие зависимости:

1. **Docker:(необязательно)**
    - Вы можете скачать и установить Docker, следуя официальным инструкциям по установке для вашей операционной системы:
    - [Руководство по установке Docker](https://docs.docker.com/get-docker/).

2. **Go: (обязательно)**
    - [Руководство по установке Golang.](https://go.dev/doc/install).

3. **PostgresSQL: (требуется, если вы не используете Docker)**
    - Убедитесь, что переменные окружения в вашем файле .yaml соответствуют среде в docker-compose.yaml
    - [Установка PostgreSQL](https://www.postgresql.org/download/)


### Бинарная установка:

1. Установите бинарный файл с помощью:

```Sh
go install -v github.com/iudicium/pryingdeep/cmd/pryingdeep@latest
```

2. Выполните следующую комманду:
```sh
pryingdeep install
```
3. Настройте значения внутри папки конфигурации под ваши нужды.

### Ручная установка:

1. Клонируйте репозиторий:
   ```sh
    git clone https://github.com/pryingbytez/pryingdeep.git
   ```

2.Настройте значения в файле конфигурации .yaml либо с использованием флагов, либо вручную.
  `Database`, `logger`, `tor` требуют ручной настройки. <br>


<u>Вам нужно будет прочитать:*</u> [Документацию Colly](https://github.com/gocolly/colly/blob/v1.2.0/colly.go).  Также обратитесь к [Обьяснению конфигурации](./CONFIG.MD#table-of-contents)


3. Соберите бинарный файл с помощью:
   - `go build` -> directory: `cmd/pryingdeep`
   - `go build cmd/pryingdeep/pryingdeep.go` -> c корневой директории, бинарный файл также будет там.

## Docker

Для запуска *pryingdeep* внутри контейнера Docker используйте эту команду:

```sh
docker-compose up
```


## Config
Прочитайте подробности о каждом параметре здесь:
[config](./CONFIG.MD#table-of-contents)


## Tor

Узнайте больше о сборке и запуске нашего контейнера tor здесь:
[Tor](../build/package/tor/README.MD#build)


<!-- USAGE EXAMPLES -->
## Использование:

Crawler:
   ```sh
   ./pryingdeep crawl
   ```

Exporter:
 ```sh
 ./pryingdeep export
  ```

Search:
 ```sh
 ./pryingdeep search
  ```


<!-- ROADMAP -->
## Roadmap

- [x] Добавить команду поиска, которая не требует использования onion-ссылок.
- [ ] Получить ключ для API Shodan для тестирования модуля иконок (favicon)
- [ ] Придумать способ получения IP-адреса сервера.
- [ ] Реализовать   идентификацию и поиск файлов.
- [ ] Найдите способ подавления ошибок дублирования уникальных записей в Gorm, так как они занимают половину экрана...


<!-- CONTRIBUTING -->
## Contributing




Если у вас есть предложение, которое сделает этот проект лучше, пожалуйста, форкните репозиторий и создайте pull-запрос. Вы также можете просто открыть issue с тегом "enhancement".
Не забудьте поставить звезду проекту! Спасибо еще раз!

1. Форкните проект
2. Создайте ветку для вашей новой функции (`git checkout -b feature/AmazingFeature`)
3. Зафиксируйте ваши изменения (`git commit -m 'Add some AmazingFeature'`)
4. Отправьте изменения в ветку (`git push origin feature/AmazingFeature`)
5. Откройте pull-запрос



<!-- LICENSE -->
## License


Распространяется под лицензией GPL-3.0. Дополнительную информацию см. в файле LICENSE.



## Video Preview

[![Video Title](https://img.youtube.com/vi/TvRT_lvixI0/maxresdefault.jpg)](https://www.youtube.com/watch?v=TvRT_lvixI0)

---

## Preview

![Preview](../web/static/preview.png)
![Data Preview](../web/static/data_preview.png)
![Data Preview 2](../web/static/data_preview2.png)

