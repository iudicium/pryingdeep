<details id="table-of-contents">
  <summary style="font-size: larger;">Table of Contents</summary>
  <ol>
    <li>
      <a href="#crawler-configuration">Crawler Configuration</a>
    </li>
    <li>
      <a href="#exporter">Exporter</a>
    </li>
    <li>
      <a href="#global-flags">Global flags</a>
    </li>
  </ol>
</details>



# Crawler Configuration

The following sections explain each setting in the `crawler` configuration:

<details>
  <summary><strong>allowed-domains</strong></summary>

  - **Description:** White list allowed domains
  - **Default Value:** `empty list`
  - ***Example:*** `old.reddit.com` &rarr; visit only old.reddit.com domains
</details>

<details>
  <summary><strong>body-size</strong></summary>

  - **Description:** Maximum size of the HTTP response body in bytes.
  - **Default Value:** `0` &rarr; unlimited
</details>

<details>
  <summary><strong>cache-dir</strong></summary>

  - **Description:** Directory path for caching. Leave empty for no caching.
  - **Default Value:** `""` (empty string)
</details>

<details>
  <summary><strong>crypto</strong></summary>

  - **Description:** Enable or disable crypto-related features.
  - **Default Value:** `false`
</details>

<details>
  <summary><strong>debug</strong></summary>

  - **Description:** Enable or disable debugging mode for GoColly.
  - **Default Value:** `false`
</details>

<details>
  <summary><strong>disallowed-domains</strong></summary>

  - **Description:** Domain black list for the crawler.
  - **Default Value:** `[]` (empty list)
  - ***Example:*** reddit.com &rarr; crawler will not visit any reddit urls
</details>

<details>
  <summary><strong>disallowed-url-filters</strong></summary>

  - **Description:** List of regular expressions to filter disallowed URLs.
  - **Default Value:** `[]` (empty list)
  - ***Example:*** `http://httpbin\.org/h.+"`
</details>

<details>
  <summary><strong>email</strong></summary>

  - **Description:** Enable or disable email-related features.
  - **Default Value:** `false`
</details>

<details>
  <summary><strong>ignore-robots-txt</strong></summary>

  - **Description:** Enable or disable ignoring the robots.txt file.
  - **Default Value:** `false`
</details>

<details>
  <summary><strong>limit-delay</strong></summary>

  - **Description:** Delay in seconds between requests.
  - **Default Value:** `0`
</details>

<details>
  <summary><strong>limit-random-delay</strong></summary>

  - **Description:** Random delay in seconds added to the fixed delay.
  - **Default Value:** `0`
</details>

<details>
  <summary><strong>max-depth</strong></summary>

  - **Description:** Maximum depth for crawling links.
  - **Default Value:** `0` &rarr; unlimited depth
</details>

<details>
  <summary><strong>phone</strong></summary>

  - **Description:** List of countries to parse phone numbers from.
  - **Default Value:** `[]` (empty list)
  - ***Example:*** "RU,NL,DE,GB,US" &rarr; You can choose which countries don't have to be every
</details>

<details>
  <summary><strong>queue-max-size</strong></summary>

  - **Description:** Maximum size of the crawler's queue.
  - **Default Value:** `50000`
</details>

<details>
  <summary><strong>queue-threads</strong></summary>

  - **Description:** Number of threads used for crawling.
  - **Default Value:** `4`
</details>

<details>
  <summary><strong>tor</strong></summary>

  - **Description:** Run the crawler through a tor proxy and allow crawling of .onion links
  - **Default Value:** `false`
</details>

<details>
  <summary><strong>url-filters</strong></summary>

  - **Description:** List of regular expressions to filter URLs.
  - **Default Value:** `[]` (empty list)
  - ***Example:*** `http://httpbin\.org/h.+` `(?:https?://)?(?:www)?(\\S*?\\.onion)\\b` -> will limit to .onion domains only
</details>

<details>
  <summary><strong>url-revisit</strong></summary>

  - **Description:** Enable or disable revisiting URLs.
  - **Default Value:** `false`
</details>

<details>
  <summary><strong>urls</strong></summary>

  - **Description:** List of starting URLs for the crawler.
  - **Default Value:** `[]` (empty list)
  - **Example:**
  ```yaml
  urls:
  - https://example.com
  - https://example2.com
  ```
</details>

<details>
  <summary><strong>user-agent</strong></summary>

  - **Description:** User agent string for HTTP requests.
  - **Default Value:** `colly - https://github.com/gocolly/colly/v2`
  - ***Example:*** `Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0`
  - ***Source:*** [useragents.me](https://www.useragents.me/)
</details>


<details>
  <summary><strong>keywords</strong></summary>

  - **Description:** A keyword, sentence, list of keywords
  - **Default Value:** `[]`
  - ***Example:*** `search -k owasp -k hacking -k "Please hack the box!"`

</details>
---

# Exporter
<!-- associations -->
<details>
  <summary><strong>associations</strong></summary>
  <ul>
    <li>
      <strong>Description:</strong> Specify the different tables you want to export from the database.
    </li>
    <li>
      <strong>Default:</strong> <code>all</code>
    </li>
    <li>
      <strong>Values:</strong>
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
      <strong>Value:</strong> <code>{}</code> - (empty json)
    </li>
    <li>
      <strong>Description:</strong> Criteria for the exporter.
    </li>
    <li>
      <strong>Explanation:</strong> If you use the LIKE keyword it will automatically perform the SQL <code>LIKE</code> statement. There's no need for adding extra <code>%</code> inside the criteria.
    </li>
    <li>
      <strong>Usage:</strong>
      <pre><code>pryingdeep -q 'title=test,"url=LIKE example.com"'</code></pre>
    </li>
  </ul>
</details>

<!-- filepath -->
<details>
  <summary><strong>filepath</strong></summary>
  <ul>
    <li>
      <strong>Value:</strong> <code>data.json</code>
    </li>
    <li>
      <strong>Description:</strong> Filepath for the exporter output.
    </li>
    <li>
      <strong>Default Value:</strong> <code>data.json</code>
    </li>
  </ul>
</details>

<!-- limit -->
<details>
  <summary><strong>limit</strong></summary>
  <ul>
    <li>
      <strong>Description:</strong> Limit the exporter to a certain number of items. 0 means every row inside the database.
    </li>
    <li>
      <strong>Default Value:</strong> <code>0</code>
    </li>
  </ul>
</details>

<!-- raw-sql -->
<details>
  <summary><strong>raw-sql</strong></summary>
  <ul>
    <li>
      <strong>Value:</strong> <code>false</code>
    </li>
    <li>
      <strong>Description:</strong> Enable or disable the use of performing raw SQL queries.
    </li>
    <li>
      <strong>Default Value:</strong> <code>false</code>
    </li>
  </ul>
</details>

<!-- raw-sql-filepath -->
<details>
  <summary><strong>raw-sql-filepath</strong></summary>
  <ul>
    <li>
      <strong>Default:</strong> <code>pkg/querybuilder/queries/select.sql</code>
    </li>
    <li>
      <strong>Description:</strong> Filepath for the raw SQL queries.
    </li>
  </ul>
</details>

<!-- sort-by -->
<details>
  <summary><strong>sort-by</strong></summary>
  <ul>
    <li>
      <strong>Value:</strong> <code>url</code>
    </li>
    <li>
      <strong>Description:</strong> Field to use for sorting. Just a generic <code>ORDER BY</code>.
    </li>
    <li>
      <strong>Default Value:</strong> <code>status_code</code>
    </li>
  </ul>
</details>

<!-- sort-order -->
<details>
  <summary><strong>sort-order</strong></summary>
  <ul>
    <li>
      <strong>Value:</strong> <code>asc</code>
    </li>
    <li>
      <strong>Description:</strong> Sort order for the exporter.
    </li>
  </ul>
</details>


<details>
  <summary><strong>offset</strong></summary>
  <ul>
    <li>
      <strong>Value:</strong> <code>0</code>
    </li>
    <li>
      <strong>Description:</strong> Number of records to skip during export.
      Keep in mind if you want to the id to start from 1, set `sort-by` to `id` and `sort-order` to `asc`
      Otherwise, the filtering might be weird, and you will get records starting from 50 when you asked
      for offset from 1.
    </li>
  </ul>
</details>


# Global Flags

<details>
  <summary><strong>-s, --silent</strong></summary>
  <ul>
    <li>
      <strong>Default:</strong> <code>false</code>
    </li>
    <li>
      <strong>Description:</strong> Use this flag to disable logging and run silently.
    </li>
  </ul>
</details>

<details>
  <summary><strong>-z, --save-config</strong></summary>
  <ul>
    <li>
      <strong>Default:</strong> <code>false</code>
    </li>
    <li>
      <strong>Description:</strong> Use this flag to save chosen options to your .yaml configuration.
    </li>
  </ul>
</details>

<details>
  <summary><strong>-c, --config &lt;path&gt;</strong></summary>
  <ul>
    <li>
      <strong>Value:</strong> The path to the .yaml configuration file. Please also keep the filename as  <code>pryingdeep</code>, otherwise the program will break.
    </li>
    <li>
      <strong>Description:</strong> Use this flag to specify the path to the .yaml configuration.
    </li>
  </ul>
</details>
