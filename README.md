# picapuento
job to automate time registry


**AWS

CRON Format
-----------------------------------
minute hour day month dayOfweek Year


Example
-----------------------------------
Every minute - cron(1 * * ? *)
Every day at midnight - cron(0 0 * * ? *)

cron (0 9,14 ? * MON-FRIDAY *)- clock in
cron (0 13,17 ? * MON-FRIDAY *) - clock out 
