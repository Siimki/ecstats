import csv

# File "results273.tsv" should contain your raw rows, tab-separated, columns:
# position, bib, gender, first_name, last_name, team, time, category, race_number, points, status
# Ensure fully-capitalized last names in the file (e.g. MUIZNIEKS, TARVIS, etc.)

# Manual overrides for duplicate names
manual_overrides = {
    ('Peeter',    'TARVIS'):    39,
    ('Virgo',     'KARU'):     241,
    ('Margus',    'KIVI'):     450,
    ('Marek',     'SAAR'):   11951,
    ('Mari-Liis', 'MÕTTUS'): 13750
}

in_path  = "../../results/to_sql_insert.txt"
out_path = "../../results/to_sql_insert_final.txt"
values   = []

with open(in_path, encoding="utf-8") as f:
    reader = csv.reader(f, delimiter="\t")
    for row in reader:
        position    = int(row[0])
        race_number = int(row[1])              # bib number is the race_number now
        first_name  = row[3]
        last_name   = row[4].upper()
        time_str    = row[6]
        points      = int(row[9])

        # Determine rider_id via override or lookup
        key = (first_name, last_name)
        if key in manual_overrides:
            rider_id = manual_overrides[key]
        else:
            rider_id = f"(SELECT id FROM riders WHERE first_name='{first_name}' AND last_name='{last_name}')"

        values.append(
            f"(273, {rider_id}, {position}, '{time_str}'::interval, {points}, {race_number})"
        )

# Build full INSERT SQL
sql = "INSERT INTO results (race_id, rider_id, position, time, points, race_number) VALUES\n"
sql += ",\n".join(values) + ";"

with open(out_path, 'w', encoding="utf-8") as f:
    f.write(sql)

print(f"Generated SQL with {len(values)} rows at {out_path}")