# GoCron

GoCron like Corntab service, but it can running on Windows use same crontab script.

#Futures

1.Cron job service like crontab service

#Script

  sec min hour  dom mon dow   command
  
  1 * * * * * C:\wamp\bin\php\php5.4.3\php.exe C:\wamp\www\phpinfo.php > php.log
  
  1-5 * * * * * C:\wamp\bin\php\php5.4.3\php.exe C:\wamp\www\phpinfo.php > php.log
  
#How to run

  GoCron  //it's will load default path:cron.cs
  
  You can write your cronjob script and load it
  
  GoCron -load my.cs 
