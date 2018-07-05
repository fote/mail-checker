# mail-checker
Util to test MTA


# How to use

```
Usage of ./mail-checker:
  -body string
    	Body (default "This is a test. Ignore this message")
  -from string
    	From (default "noreply@mydomain.ru")
  -mailHost string
    	Mail server address (default "127.0.0.1")
  -mailPort string
    	Mail server port (default "25")
  -password string
    	Password (default "password")
  -subject string
    	Subject (default "Test message")
  -to string
    	To (default "mymail@yandex.ru")
  -username string
    	From (default "noreply")
 ```
 
 Example usage:
 ```
 ./mail-checker -mailHost 172.17.0.1 -mailPort 25 -username user1 -password user1password -to mymail@gmail.com
 ````
