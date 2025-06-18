import os
import re

# Folder containing cleaned TXT files
input_folder = "../../results/FILTERTT/Final_Cleaned_TXT/2022/"
output_folder = "../../results/FILTERTT/Final_Cleaned_TXT/"

# Ensure output directory exists
os.makedirs(output_folder, exist_ok=True)

# Pattern to match valid race result lines (start with a number)
valid_line_pattern = re.compile(r'^\d+\s')

# Pattern to remove " I ", " II ", " III " when surrounded by spaces
roman_numeral_pattern = re.compile(r'\s(I{1,3})\s')

# Process all cleaned TXT files
for filename in os.listdir(input_folder):
    if filename.endswith("_cleaned.txt"):  # Process only previously cleaned files
        input_file_path = os.path.join(input_folder, filename)
        output_file_path = os.path.join(output_folder, filename.replace("_cleaned.txt", "_final.txt"))

        with open(input_file_path, "r", encoding="utf-8") as infile, open(output_file_path, "w", encoding="utf-8") as outfile:
            for line in infile:
                if valid_line_pattern.match(line):  # Keep only valid race result lines
                    cleaned_line = roman_numeral_pattern.sub(" ", line)  # Remove " I ", " II ", " III "
                    outfile.write(cleaned_line.strip() + "\n")  # Strip leading/trailing spaces

        print(f"Processed and saved final cleaned file: {output_file_path}")

print("Final cleaning completed for all files.")
