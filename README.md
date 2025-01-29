# ecstats

## SQL Commands 

Backup command for terminal. 

    psql -U postgres -p 5433 -d ecstats < ecstats_backup.sql

Making backup file:

    pg_dump -U postgres -p 5433 -d ecstats -f ecstats_backup18.sql


info about DB:

Windspeed is in km/h
weather info is taken from: wunderground.com 

Race types:
ROAD = "Road Race"
MTB = "MTB"
Gravel = "Gravel" 
TT = "TT"

if no nationality. Nat = XXX
if DNF then pos = 0 
if name is "";  name = DUMMY Data


