import re
import os

# List of team names to remove
team_names = [
    "Sparta ProTeam", "Sparta ratturid", "Redbike - Actual", "Sportland Bottecchia Team",
    "CFC", "CFC Spordiklubi", "Arctic Sport Club", "Rein Taaramäe Rattaklubi",
    "Spordipartner", "velo.clubbers", "Viimsi Rattaklubi", "REKORD CC",
    "Actual Print", "Filter", "Alecoq Team", "Rae Rattaklubi", "21CC",
    "Arctic Sport", "Team Vänt", "VIIMSI I", "VIIMSI II", "Tartu 2024", 
    "Tartu Ülikool", "Cycling Tartu", "Tartu Velo", "TÜASK", "Soudal Quick-Step",
    "UAE Team Emirates", "Jumbo-Visma", "Ineos Grenadiers", "Movistar Team",
    "EF Education-EasyPost", "BORA-hansgrohe", "AG2R Citroën Team",
    "Groupama-FDJ", "Trek-Segafredo", "Lotto Soudal", "Israel Start-Up Nation",
    "Astana Qazaqstan Team", "BikeExchange-Jayco", "Intermarché-Wanty-Gobert",
    "Cofidis", "Burgos-BH", "Caja Rural-Seguros RGA", "Euskaltel-Euskadi",
    "Human Powered Health", "Uno-X Pro Cycling Team",   "Aave Spordiklubi",
    "Ampler Development Team",
    "Arctic Sport Club",
    "AT Sport Team",
    "BalticChainCycling",
    "Bike Fanatics CC",
    "BORA-hansgrohe",
    "CFC",
    "CFC Spordiklubi",
    "CFC Suusakool",
    "Eneicat-RBH Global",
    "EF Education-EasyPost",
    "Filter Temposari",
    "HAWAII EXPRESS",
    "Ineos Grenadiers",
    "Intermarché-Wanty-Gobert",
    "Israel Start-Up Nation",
    "Järva-Jaani RSK",
    "Järva-Jaani Ratta- ja Suusaklubi",
    "Jumbo-Visma",
    "Kalevi Jalgrattakool",
    "Kalevi Spordikool",
    "Kuusalu Rattaklubi",
    "Movistar Team",
    "Nõmme Rattaklubi",
    "Pärnu Kalevi Spordikool",
    "Peloton",
    "ProShop Team",
    "Raplamaa Rattaklubi KoMo",
    "Rae Rattaklubi",
    "Rein Taaramäe Rattaklubi",
    "Sparta Spordiklubi",
    "Sportland Bottecchia Team",
    "Spordipartner",
    "Tartu 2024",
    "Tartu Ülikool",
    "Team Arkea-Samsic",
    "Team Kodar",
    "Team Vänt",
    "Trek-Segafredo",
    "TriPassion Triatloniklubi",
    "TÜASK",
    "Uno-X Pro Cycling Team",
    "VELO CLUBBERS",
    "VeloHunt ProTeam",
    "Viimsi Rattaklubi",
    "WallleniumSPORT",
    "CC Rota Mobilis",
    "DOLTCINI Latvia Cycling Team",
    "FNT Spordiklubi",
    "RMIT SK",
    "SK Rakke",
    "Spordiklubi",
    "Swedbank Spordiklubi",
    "Tabasalu Triatloniklubi",
    "Team Zoot Europe",
    "Velo+ Bottari Baltic",
    "Viimsi Sport",
      "CC Rota Mobilis",
    "DOLTCINI Latvia Cycling Team",
    "FNT Spordiklubi",
    "RMIT SK",
    "SK Rakke",
    "Spordiklubi",
    "Swedbank Spordiklubi",
    "Tabasalu Triatloniklubi",
    "Team Zoot Europe",
    "Velo+ Bottari Baltic",
    "Viimsi Sport",
    "El Pr",
    "Rietumu - Delfin",
    "Gaismas Magija - EMU",
    "Rein Taaramäe ratta4k5l.u6b ki",
    "Triatloniklubi UP43",
    "Triatloniklubi", 
    "REKORD",
    "KoMo",
    "Freesport",
    "ImpulseStore",
    "Big Express",
    "/KodarSport",
    "CC Rota Mobiis",
    "JK paralepa",
    "Jooksupartner",
    "UP Sport",
    "SK Keremeister",
    "SJK Viiking",
    "RedBike",
    "R.U.S.T",
    "VH Team",
    "MTÜ Trismile",
    "Suusakool",
    "Triathlon Estonia",
    "Steel Athletic",
    "Trismile",
    "Maardu",
    "KJK",
    "Koidu Suusaklubi",
    "Chapo Endurance Team",
    
]

# Create a regex pattern that matches any team name (with word boundaries)
team_pattern = re.compile(
    r'\b(?:' + '|'.join(map(re.escape, team_names)) + r')\b',
    re.IGNORECASE
)

# Regex pattern to match valid lines (starting with rank number)
valid_line_pattern = re.compile(r'^\d+\s+\d+')  # Rank + bib number

# Define paths
input_folder = "../../results/FILTERTT/txt_files/2022/"
output_folder = "../../results/FILTERTT/txt_files/2022_cleaned/"

# Ensure output directory exists
os.makedirs(output_folder, exist_ok=True)

# Process files
for filename in os.listdir(input_folder):
    input_path = os.path.join(input_folder, filename)
    if not os.path.isfile(input_path) or not filename.endswith('.txt'):
        continue

    output_path = os.path.join(output_folder, f"cleaned_{filename}")

    try:
        with open(input_path, 'r', encoding='utf-8') as f:
            content = f.read()

        # Remove team names
        cleaned_content = team_pattern.sub('', content)

        # Keep only valid lines and strip whitespace
        cleaned_lines = [
            line.strip() for line in cleaned_content.split('\n') 
            if valid_line_pattern.match(line)
        ]

        # Save cleaned content
        with open(output_path, 'w', encoding='utf-8') as f:
            f.write('\n'.join(cleaned_lines))

        print(f"Processed: {filename} → Saved to {output_path}")

    except Exception as e:
        print(f"Error processing {filename}: {str(e)}")