import os
import pdfplumber

# Folder where the PDFs are stored
PDF_FOLDER = "../../results/FILTERTT/2022/"
OUTPUT_FOLDER = "../../results/FILTERTT/txt_files/"  # Save extracted text here

# Ensure output directory exists
os.makedirs(OUTPUT_FOLDER, exist_ok=True)

def extract_text_from_pdf(pdf_path, txt_path):
    """Extracts text from a PDF and saves it as a .txt file."""
    text_output = ""
    with pdfplumber.open(pdf_path) as pdf:
        for page in pdf.pages:
            extracted_text = page.extract_text()
            if extracted_text:  # Avoid NoneType issues
                text_output += extracted_text + "\n"

    with open(txt_path, "w", encoding="utf-8") as f:
        f.write(text_output)

    print(f"Extracted: {pdf_path} → {txt_path}")

def process_all_pdfs():
    """Processes all PDFs in the folder and extracts text to .txt files."""
    pdf_files = [f for f in os.listdir(PDF_FOLDER) if f.endswith(".pdf")]

    if not pdf_files:
        print("No PDFs found in the folder!")
        return

    for pdf_file in pdf_files:
        pdf_path = os.path.join(PDF_FOLDER, pdf_file)
        txt_file = os.path.splitext(pdf_file)[0] + ".txt"  # Change .pdf to .txt
        txt_path = os.path.join(OUTPUT_FOLDER, txt_file)

        extract_text_from_pdf(pdf_path, txt_path)

if __name__ == "__main__":
    process_all_pdfs()
