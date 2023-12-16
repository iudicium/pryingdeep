<details id="table-of-contents">
    <summary style="font-size: larger;">Содержание</summary>
    <ol>
        <li>
            <a href="#crawler-configuration">Настройка краулера</a>
        </li>
        <li>
            <a href="#exporter">Экспорт</a>
        </li>
        <li>
            <a href="#global-flags">Глобальные флаги</a>
        </li>
    </ol>
</details>




# Crawler Configuration

Следующие разделы поясняют каждую настройку в конфигурации `crawler`:


<details>
    <summary><strong>allowed-domains</strong></summary>

    - **Описание:** Белый список разрешенных доменов
    - **Значение по умолчанию:** `пустой список`
    - ***Пример:*** `old.reddit.com`  &rarr; посещение только доменов old.reddit.com
</details>


<details>
    <summary><strong>body-size</strong></summary>

    - **Описание:** Максимальный размер тела HTTP-ответа в байтах.
    - **Значение по умолчанию:** `0` -> значит без ограничений
</details>

<details>
    <summary><strong>cache-dir</strong></summary>

    - **Описание:** Путь каталога для кэширования. Оставьте пустым для отключения кэширования.
    - **Значение по умолчанию:** `""` (пустая строка)
</details>


<details>
    <summary><strong>crypto</strong></summary>

    - **Описание:** Включение или отключение функций, связанных с крипто-модулем.
    - **Значение по умолчанию:** `false`
</details>

<details>
    <summary><strong>debug</strong></summary>

    - **Описание:** Включение или отключение режима отладки для GoColly.
    - **Значение по умолчанию:** `false`
</details>


<details>
    <summary><strong>disallowed-domains</strong></summary>

    - **Описание:** Черный список доменов для краулера.
    - **Значение по умолчанию:** `[]` (пустой список)
    - ***Пример:*** reddit.com &rarr; краулер не будет посещать ни один URL reddit
</details>

<details>
    <summary><strong>disallowed-url-filters</strong></summary>

    - **Описание:** Список регулярных выражений для фильтрации запрещенных URL.
    - **Значение по умолчанию:** `[]` (пустой список)
    - ***Пример:*** `http://httpbin\.org/h.+"`
</details>

<details>
    <summary><strong>email</strong></summary>

    - **Описание:** Включение или отключение функций, связанных с электронной почтой.
    - **Значение по умолчанию:** `false`
</details>

<details>
    <summary><strong>ignore-robots-txt</strong></summary>

    - **Описание:** Включение или отключение игнорирования файла robots.txt.
    - **Значение по умолчанию:** `false`
</details>


<details>
    <summary><strong>limit-delay</strong></summary>

    - **Описание:** Задержка в секундах между запросами.
    - **Значение по умолчанию:** `0`
</details>

<details>
    <summary><strong>limit-random-delay</strong></summary>

    - **Описание:** Случайная задержка в секундах, добавляемая к фиксированной задержке.
    - **Значение по умолчанию:** `0`
</details>


<details>
  <summary><strong>max-depth</strong></summary>

  - **Description:** Maximum depth for crawling links.
  - **Default Value:** `0` &rarr; unlimited depth
</details>

<details>
    <summary><strong>max-depth</strong></summary>

    - **Описание:** Максимальная глубина для обхода ссылок.
    - **Значение по умолчанию:** `0` &rarr; неограниченная глубина
</details>

<details>
    <summary><strong>phone</strong></summary>

    - **Описание:** Список стран для извлечения номеров телефонов.
    - **Значение по умолчанию:** `[]` (пустой список)
    - ***Пример:*** "RU, NL, DE, US" &rarr; можно выбрать, из каких стран извлекать номера
</details>


<details>
    <summary><strong>queue-max-size</strong></summary>

    - **Описание:** Максимальный размер очереди краулера.
    - **Значение по умолчанию:** `50000`
</details>

<details>
    <summary><strong>queue-maxsize</strong></summary>

    - **Описание:** Количество потоков, используемых для краулинга.
    - **Значение по умолчанию:** `4`
</details>


<details>
    <summary><strong>tor</strong></summary>

    - **Описание:** Запуск краулера через прокси `tor` и разрешение на краулинг .onion-ссылок.
    - **Значение по умолчанию:** `false`
</details>

<details>
    <summary><strong>url-filters</strong></summary>

    - **Описание:** Список регулярных выражений для фильтрации URL.
    - **Значение по умолчанию:** `[]` (пустой список)
    - ***Пример:*** `http://httpbin\.org/h.+` `(?:https?://)?(?:www)?(\\S*?\\.onion)\\b` -> ограничит только домены .onion
</details>

<details>
    <summary><strong>url-revisit</strong></summary>

    - **Описание:** Включение или отключение повторного посещения URL.
    - **Значение по умолчанию:** `false`
</details>

<details>
    <summary><strong>urls</strong></summary>

    - **Описание:** Список начальных URL для краулера.
    - **Значение по умолчанию:** `[]` (пустой список)
    - **Пример:**
    ```yaml
    urls:
    - https://example.com
    - https://example2.com

</details>

<details>
    <summary><strong>user-agent</strong></summary>

    - **Описание:** Строка пользовательского агента для HTTP-запросов.
    - ***Пример:*** `Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0`
    - ***Источник:*** [useragents.me](https://www.useragents.me/)
</details>


<details>
    <summary><strong>keywords</strong></summary>

    - **Описание:** Ключевое слово, предложение, список ключевых слов.
    - **Значение по умолчанию:** `[]`
    - ***Пример:*** `search -k owasp -k hacking -k "Please hack the box!"`
</details>

---

# Exporter
<!-- associations -->
<details>
    <summary><strong>associations</strong></summary>
    <ul>
        <li>
            <strong>Описание:</strong> Укажите различные SQL таблицы, которые вы хотите экспортировать из базы данных.
        </li>
        <li>
            <strong>Значение по умолчанию:</strong> <code>all</code>
        </li>
        <li>
            <strong>Возможные значения:</strong>
            <ul>
                <li><code>"WP" - WordPress</code></li>
                <li><code>"E" - Email</code></li>
                <li><code>"P" - PhoneNumbers</code></li>
                <li><code>"C" - Crypto</code></li>
            </ul>
        </li>
    </ul>
</details>


<!-- criteria -->
<details>
    <summary><strong>criteria</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> <code>{}</code> - (пустой JSON)
        </li>
        <li>
            <strong>Описание:</strong> Критерии для экспортера.
        </li>
        <li>
            <strong>Объяснение:</strong> Если вы используете ключевое слово LIKE, оно автоматически выполняет оператор SQL <code>LIKE</code>. Нет необходимости добавлять дополнительные <code>%</code> внутри критерии.
        </li>
        <li>
            <strong>Использование:</strong>
            <pre><code>pryingdeep -q 'title=test,"url=LIKE example.com"'</code></pre>
        </li>
    </ul>
</details>


<!-- filepath -->
<details>
    <summary><strong>filepath</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> <code>data.json</code>
        </li>
        <li>
            <strong>Описание:</strong> Путь к файлу для вывода экспорта.
        </li>
        <li>
            <strong>Значение по умолчанию:</strong> <code>data.json</code>
        </li>
    </ul>
</details>


<!-- limit -->
<details>
    <summary><strong>limit</strong></summary>
    <ul>
        <li>
            <strong>Описание:</strong> Ограничивает экспортер до определенного количества элементов. 0 означает каждую строку в базе данных.
        </li>
        <li>
            <strong>Значение по умолчанию:</strong> <code>0</code>
        </li>
    </ul>
</details>


<!-- raw-sql -->
<details>
    <summary><strong>raw-sql</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> <code>false</code>
        </li>
        <li>
            <strong>Описание:</strong> Включение или отключение выполнения запросов сырого SQL.
        </li>
        <li>
            <strong>Значение по умолчанию:</strong> <code>false</code>
        </li>
    </ul>
</details>

<!-- raw-sql-filepath -->
<details>
    <summary><strong>raw-sql-filepath</strong></summary>
    <ul>
        <li>
            <strong>Значение по умолчанию:</strong> <code>pkg/querybuilder/queries/select.sql</code>
        </li>
        <li>
            <strong>Описание:</strong> Путь к файлу с сырыми SQL-запросами.
        </li>
    </ul>
</details>

<!-- sort-by -->
<details>
    <summary><strong>sort-by</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> <code>url</code>
        </li>
        <li>
            <strong>Описание:</strong> Поле, используемое для сортировки. Обычный <code>ORDER BY</code>.
        </li>
        <li>
            <strong>Значение по умолчанию:</strong> <code>status_code</code>
        </li>
    </ul>
</details>

<!-- sort-order -->
<details>
    <summary><strong>sort-order</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> <code>asc</code>
        </li>
        <li>
            <strong>Описание:</strong> Порядок сортировки для экспортера.
        </li>
    </ul>
</details>

<details>
    <summary><strong>offset</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> <code>0</code>
        </li>
        <li>
            <strong>Описание:</strong> Количество записей для пропуска во время экспорта.
            Имейте в виду, что если вы хотите, чтобы id начинался с 1, установите `sort-by` в `id`, а `sort-order` в `asc`.
            В противном случае фильтрация может быть странной, и вы получите записи, начиная с 50, когда запросили
            пропуск от 1.
        </li>
    </ul>
</details>



# Global Flags

<details>
    <summary><strong>-s, --silent</strong></summary>
    <ul>
        <li>
            <strong>Значение по умолчанию:</strong> <code>false</code>
        </li>
        <li>
            <strong>Описание:</strong> Используйте этот флаг, чтобы отключить логгинг  и запустить инструмент в тихом режиме.
        </li>
    </ul>
</details>

<details>
    <summary><strong>-z, --save-config</strong></summary>
    <ul>
        <li>
            <strong>Значение по умолчанию:</strong> <code>false</code>
        </li>
        <li>
            <strong>Описание:</strong> Используйте этот флаг, чтобы сохранить выбранные параметры в ваш файл конфигурации `.yaml.`
        </li>
    </ul>
</details>

<details>
    <summary><strong>-c, --config &lt;path&gt;</strong></summary>
    <ul>
        <li>
            <strong>Значение:</strong> Путь к файлу конфигурации .yaml. Пожалуйста, также сохраните имя файла как <code>pryingdeep</code>, в противном случае программа сломается.
        </li>
        <li>
            <strong>Описание:</strong> Используйте этот флаг, чтобы указать путь к файлу конфигурации `.yaml.`
        </li>
    </ul>
</details>
