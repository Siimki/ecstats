#!/bin/bash

# Define years and their corresponding max stage numbers
years=(2015 2016 2017 2018 2019 2020 2021)
stages=(6 7 7 7 7 5 6)

# Loop through each year and stage
for i in "${!years[@]}"; do
  year="${years[i]}"
  max_stage="${stages[i]}"
  
  echo "Downloading files for year: $year (stages 1-$max_stage)"
  
  for ((stage=1; stage<=max_stage; stage++)); do
    # Build the URL
    url="https://tolknet.ee/results/${year}/filter_tt/${stage}_etapp/res_all.pdf"
    
    # Define the output filename
    filename="${year}_stage_${stage}.pdf"
    
    # Download with wget
    wget -q --show-progress -O "$filename" "$url"
    
    # Optional: Add a small delay to avoid server overload
    sleep 1
  done
done

echo "Download complete!"