# ecstats

## Current workflow 

 - Download results. Prefer live timing website for known format. 
 - If needed, do data cleaning with python/GO. 
 - Make Regex at regex101.com to group(name, time, place, etc..) the results. 
 - Update config.yaml with regex and race_id. 
 - If we first run main.go the program shows possible problematic names. Take a look at them manually from terminal before continuing. 
 - Continue inserting results if previous steps are done. 
 - Query check the amount of results manually from DB.
 - Make backup from DB.


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
if name is "";  name = DUMMY Data.


