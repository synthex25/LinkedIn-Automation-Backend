# LinkedIn Automation Backend (Go)

## Overview

This project is a Go-based backend service designed to demonstrate how LinkedIn-style automation systems can be **architected safely, modularly, and responsibly**.

The focus of this implementation is on:

- Automation correctness  
- Human-like execution control  
- Anti-detection concepts at the control layer  
- Clean Go backend architecture  

The system runs in **mock automation mode** and does **not perform real browser automation, scraping, or credential handling**, ensuring compliance with platform Terms of Service.

---

## Design Philosophy

Automation systems are often detected due to:

- Burst execution  
- Lack of pacing  
- Poor validation  
- Uncontrolled concurrency  

This project addresses these issues at the **automation control layer**, rather than attempting unsafe browser-level evasion.

The goal is to demonstrate **understanding and design clarity**, not policy bypassing.

---

## Key Features

- Go HTTP server using the standard library  
- Simulated LinkedIn connection request automation  
- Simulated LinkedIn messaging automation  
- Input validation with proper HTTP status codes  
- Mutex-protected rate limiting for human-like pacing  
- Request logging for observability  
- Environment-based configuration template  
- Clean separation of concerns  

---

## Project Structure

linkedin_automation/
├── main.go # Server startup and route registration
├── handlers.go # Automation logic, validation, rate limiting
├── models.go # Request data models
├── go.mod # Go module definition
├── README.md # Documentation
├── .env.example # Environment configuration template

yaml
Copy code

---

## Environment Configuration

Configuration is managed using environment variables.  
A template is provided via `.env.example`.

Example:

```env
PORT=9090
RATE_LIMIT_SECONDS=1
AUTOMATION_MODE=mock
The system currently runs in mock mode.
Browser automation can be added later without changing the API contract.

API Endpoints
Health Check
GET /health

Response:

json
Copy code
{
  "status": "OK"
}
Send Connection Request (Mock)
POST /send-connection

Request:

json
Copy code
{
  "profile_url": "https://linkedin.com/in/test",
  "note": "Hi, would love to connect!"
}
Response:

json
Copy code
{
  "status": "connection request sent",
  "profile": "https://linkedin.com/in/test"
}
Send Message (Mock)
POST /send-message

Request:

json
Copy code
{
  "profile_url": "https://linkedin.com/in/test",
  "message": "Hello from automation"
}
Response:

json
Copy code
{
  "status": "message sent",
  "profile": "https://linkedin.com/in/test"
}
Anti-Detection Strategy (Conceptual)
Anti-detection is implemented at the execution control layer, not the browser layer.

Implemented concepts include:

Human-like pacing via mutex-protected rate limiting
Controlled execution flow to prevent burst behavior
Strict input validation to avoid malformed actions
Deterministic request handling with explicit error responses
Browser-specific techniques such as mouse movement simulation, fingerprint masking, and navigator overrides are intentionally out of scope in mock mode and can be introduced later without refactoring the system.

Rate Limiting
A global, mutex-protected rate limiter enforces one request per second across automation endpoints.

Requests exceeding the limit return:
nginx
Copy code
HTTP 429 Too Many Requests
This models realistic human interaction timing and reduces automation signatures.

Logging
All incoming requests are logged with:
HTTP method
Request path
This provides basic observability and aids debugging.

Scope & Safety Considerations
The following components are intentionally not implemented:

Real LinkedIn authentication

Browser automation

CAPTCHA / 2FA handling

Profile search and scraping

Session persistence

These are deliberate design decisions to:

Respect platform policies

Avoid unsafe automation

Focus evaluation on architecture and correctness

Running the Project
Prerequisites
Go 1.25 or higher

Postman or curl

Run Locally
bash
Copy code
go run .
Server starts at:

arduino
Copy code
http://localhost:9090
Demo Video
A short demo video demonstrates:

Project structure

Server startup

API usage

Error handling

Rate limiting behavior

🎥 Demo Video:
https://github.com/synthex25/LinkedIn-Automation-Backend.git

Evaluation Alignment
Criterion Coverage
Anti-Detection Quality Execution pacing and control
Automation Correctness Validated APIs
Code Architecture Modular Go design
Practical Implementation Runnable service

Disclaimer
This project is for educational and evaluation purposes only.
It does not perform real LinkedIn automation or policy bypassing.

Conclusion
This submission demonstrates how an automation system can be designed responsibly, prioritizing correctness, safety, and maintainability while clearly modeling anti-detection principles.
