# GoCron

GoCron like crontab service, but it can running on Windows use crontab like script.

#Futures

1.Cron job service like crontab service

#Script

  sec min hour  dom mon dow   command
  
  1 * * * * * C:\wamp\bin\php\php5.4.3\php.exe C:\wamp\www\phpinfo.php > php.log
  
  1-5 * * * * * C:\wamp\bin\php\php5.4.3\php.exe C:\wamp\www\phpinfo.php > php.log
  
#How to run

## it's will load default path:cron.cs

<code>
  GoCron
</code>
  
## You can write your cronjob script and load it

<code>
  GoCron -load my.cs 
</code>
