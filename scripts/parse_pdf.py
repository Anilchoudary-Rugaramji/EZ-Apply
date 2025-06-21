from fastapi import FastAPI, HTTPException, Request
from fastapi.responses import JSONResponse
from pydantic import BaseModel
import fitz  # PyMuPDF
import re
from io import BytesIO

app = FastAPI()

# Regular expressions for various fields
name_regex = r"([A-Z][a-z]+(?: [A-Z][a-z]+)+)"
email_regex = r"[\w\.-]+@[\w\.-]+\.\w+"
phone_regex = r"(\+?\d{1,3}[-.\s]?)?(\(?\d{3}\)?[-.\s]?)?\d{3}[-.\s]?\d{4}"
designation_regex = r"(Software Engineer|Data Scientist|Developer|Engineer|Manager|Architect|Lead)"
experience_regex = r"(\d+)\s*(year|yr|yrs|experience)"
education_regex = r"(Bachelor|Master|PhD|Diploma|Associate)[\w\s]*"
location_regex = r"(New York|San Francisco|Los Angeles|Seattle|Boston|Chicago|Remote|[A-Za-z\s]+)"
skills_keywords = ["Go", "Python", "Java", "AWS", "Docker", "Kubernetes",
                   "React", "PostgreSQL", "SQL", "Machine Learning", "Data Science"]

# Model to return JSON response


class ResumeMetaData(BaseModel):
    name: str
    email: str
    phone: str
    designation: str
    experience: int
    highest_education: str
    location: str
    skills: list[str]

# Route to handle PDF parsing


@app.post("/parse_pdf", response_model=ResumeMetaData)
async def parse_pdf(request: Request):
    # Get the raw PDF bytes from the request
    pdf_data = await request.body()

    if not pdf_data:
        raise HTTPException(status_code=400, detail="No PDF data provided")

    # Verify PDF signature
    # if not pdf_data.startswith(b'%PDF-'):
    #     raise HTTPException(
    #         status_code=415, detail="Invalid PDF format: File does not start with PDF signature")

    # Create BytesIO directly from request data
    pdf_bytes = BytesIO(pdf_data)

    try:
        # Create PyMuPDF document from bytes
        doc = fitz.open(stream=pdf_bytes, filetype="pdf")

        # Extract text from PDF document
        text = ""
        for page_num in range(doc.page_count):
            page = doc.load_page(page_num)
            text += page.get_text("text")

        # Extract metadata from the extracted text
        metadata = extract_resume_metadata(text)

        return metadata

    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# Function to extract resume metadata


def extract_resume_metadata(text):
    name = extract_first_match(name_regex, text)
    email = extract_first_match(email_regex, text)
    phone = extract_first_match(phone_regex, text)
    designation = extract_first_match(designation_regex, text)
    experience = extract_experience(text)
    highest_education = extract_first_match(education_regex, text)
    location = extract_first_match(location_regex, text)
    skills = extract_skills(text)

    return {
        "name": name or "Not Found",  # Handle None case
        "email": email or "Not Found",
        "phone": phone or "Not Found",
        "designation": designation or "Not Found",
        # If experience is None, default to 0
        "experience": experience if experience else 0,
        "highest_education": highest_education or "Not Found",
        "location": location or "Not Found",
        # Return empty list if no skills
        "skills": skills if skills else ["Not Found"]
    }


def extract_first_match(regex, text):
    # Find the first match of the regex in the text
    match = re.search(regex, text)
    if match:
        return match.group(0)
    return None


def extract_experience(text):
    # Extract experience in years
    match = re.search(experience_regex, text)
    if match:
        return int(match.group(1))  # Return experience as integer
    return None  # Return None if not found


def extract_skills(text):
    # Extract skills based on predefined keywords
    skills_found = []
    for skill in skills_keywords:
        if re.search(r"\b" + re.escape(skill) + r"\b", text, re.IGNORECASE):
            skills_found.append(skill)
    return skills_found


# Run the app with Uvicorn (ASGI server)
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5000)
