SELECT * FROM web_pages
INNER JOIN cryptos ON web_pages.id = cryptos.web_page_id
INNER JOIN emails ON web_pages.id = emails.web_page_id
INNER JOIN phone_numbers ON web_pages.id = phone_numbers.web_page_id
INNER JOIN wordpress ON web_pages.id = wordpress.web_page_id
