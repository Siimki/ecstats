# ecstats

## SQL Commands 

Backup command for terminal. 

    psql -U postgres -p 5433 -d ecstats < ecstats_backup.sql

Making backup file:

    pg_dump -U postgres -p 5433 -d ecstats -f ecstats_backup18.sql


# There is yaml file with important DB data
    
    Need to have it locally. 

info about DB:

Windspeed is in km/h
weather info is taken from: wunderground.com for TRR/TRM
weather info after that is taken from estonian forecasts


Race types:
ROAD = "Road Race"
MTB = "MTB"
Gravel = "Gravel" 
TT = "TT" || "TTP" for pair TT || "TTT" for Team time trial

if no nationality. Nat = XXX
if DNF then pos = 0 // change to nothing perhaps. 
if name is "";  name = DUMMY Data


