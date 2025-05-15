# AetherStack Development Guide

This guide provides instructions for developers looking to contribute to AetherStack or set up the system for their own use. It covers cloning the repository, setting up the development environment for each module (CLI, Backend, VS Code Extension, Web Dashboard), how modules communicate, and available development scripts.

## Table of Contents

- [1. Prerequisites](#1-prerequisites)
- [2. Cloning the Repository](#2-cloning-the-repository)
- [3. Overall Project Structure](#3-overall-project-structure)
- [4. Setting Up Core Tools](#4-setting-up-core-tools)
  - [4.1. Go (for CLI)](#41-go-for-cli)
  - [4.2. Python (for Backend)](#42-python-for-backend)
  - [4.3. Node.js & npm/yarn (for VS Code Extension & Dashboard)](#43-nodejs--npmyarn-for-vs-code-extension--dashboard)
- [5. Module-Specific Setup](#5-module-specific-setup)
  - [5.1. CLI (Go + Cobra)](#51-cli-go--cobra)
    - [5.1.1. Environment Setup](#511-environment-setup)
    - [5.1.2. Building and Running](#512-building-and-running)
    - [5.1.3. Development Scripts](#513-development-scripts)
  - [5.2. Backend API Server (Python + FastAPI)](#52-backend-api-server-python--fastapi)
    - [5.2.1. Environment Setup (Poetry/PDM/venv)](#521-environment-setup-poetrypdmvenv)
    - [5.2.2. Database Setup (SQLite & PostgreSQL)](#522-database-setup-sqlite--postgresql)
    - [5.2.3. Running the Server](#523-running-the-server)
    - [5.2.4. Development Scripts](#524-development-scripts)
  - [5.3. VS Code Extension (TypeScript)](#53-vs-code-extension-typescript)
    - [5.3.1. Environment Setup](#531-environment-setup)
    - [5.3.2. Building and Running in Development Mode](#532-building-and-running-in-development-mode)
    - [5.3.3. Development Scripts](#533-development-scripts)
  - [5.4. Web Dashboard (Next.js + TailwindCSS)](#54-web-dashboard-nextjs--tailwindcss)
    - [5.4.1. Environment Setup](#541-environment-setup)
    - [5.4.2. Running the Development Server](#542-running-the-development-server)
    - [5.4.3. Development Scripts](#543-development-scripts)
- [6. Inter-Module Communication](#6-inter-module-communication)
  - [6.1. CLI to Backend](#61-cli-to-backend)
  - [6.2. VS Code Extension to Backend/CLI](#62-vs-code-extension-to-backendcli)
  - [6.3. Web Dashboard to Backend](#63-web-dashboard-to-backend)
  - [6.4. Ports and Sockets](#64-ports-and-sockets)
- [7. Documentation Setup (MkDocs)](#7-documentation-setup-mkdocs)
  - [7.1. Installation](#71-installation)
  - [7.2. Building and Serving Docs Locally](#72-building-and-serving-docs-locally)
- [8. Coding Standards and Conventions](#8-coding-standards-and-conventions)
  - [8.1. Linting and Formatting](#81-linting-and-formatting)
  - [8.2. Commit Messages](#82-commit-messages)
- [9. Debugging](#9-debugging)
- [10. Sample Development Workflows](#10-sample-development-workflows)
  - [10.1. Adding a New CLI Command](#101-adding-a-new-cli-command)
  - [10.2. Adding a New Backend API Endpoint](#102-adding-a-new-backend-api-endpoint)
  - [10.3. Developing a VS Code Extension Feature](#103-developing-a-vs-code-extension-feature)
- [11. Contribution Guidelines](#11-contribution-guidelines)

---

## 1. Prerequisites

*(TODO: List essential software that developers must have installed before starting, e.g., Git, Docker (optional), text editor/IDE like VS Code.)*

## 2. Cloning the Repository

```bash
git clone https://github.com/your-org/aetherstack.git # Replace with actual repo URL
cd aetherstack
```

## 3. Overall Project Structure

*(TODO: Briefly describe the mono-repo structure, referencing the main directories: `cli/`, `backend/`, `vscode-extension/`, `dashboard/`, `docs/`.)*

## 4. Setting Up Core Tools

### 4.1. Go (for CLI)
*(TODO: Instructions for installing Go, recommended version. Link to official Go download page.)*
- Recommended Go version: 1.x (latest stable)

### 4.2. Python (for Backend)
*(TODO: Instructions for installing Python, recommended version. Mention virtual environments.)*
- Recommended Python version: 3.9+

### 4.3. Node.js & npm/yarn (for VS Code Extension & Dashboard)
*(TODO: Instructions for installing Node.js and npm/yarn, recommended versions. Link to official Node.js download page.)*
- Recommended Node.js version: 18.x or 20.x (LTS)
- Recommended package manager: `npm` (comes with Node.js) or `yarn` (optional, install separately)

## 5. Module-Specific Setup

### 5.1. CLI (Go + Cobra)
Located in the `/cli` directory.

#### 5.1.1. Environment Setup
*(TODO: Steps to prepare the Go environment, `go mod tidy`.)*
```bash
cd cli
go mod tidy
# Any other setup steps
```

#### 5.1.2. Building and Running
*(TODO: How to build the CLI binary and run it.)*
```bash
# Build
go build -o aether ./main.go

# Run
./aether --help
./aether init --name my-project # Example command
```

#### 5.1.3. Development Scripts
*(TODO: Any `Makefile` targets or helper scripts for common dev tasks like linting, testing, building.)*

### 5.2. Backend API Server (Python + FastAPI)
Located in the `/backend` directory.

#### 5.2.1. Environment Setup (Poetry/PDM/venv)
*(TODO: Instructions for setting up a Python virtual environment and installing dependencies using `requirements.txt`, Poetry, or PDM. SQLModel/SQLAlchemy setup.)*

**Using `venv` and `pip` (default):**
```bash
cd backend
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -r requirements.txt
```
*(TODO: Add instructions for Poetry or PDM if chosen as primary.)*

#### 5.2.2. Database Setup (SQLite & PostgreSQL)
*(TODO: How to set up the initial SQLite database. Instructions or pointers for setting up PostgreSQL for development if needed. Migrations using Alembic if SQLAlchemy is chosen.)*
- SQLite is used by default for ease of setup. The database file will be created automatically (e.g., `aetherstack.db`).
- For PostgreSQL, ensure a running instance and update connection strings in configuration.

#### 5.2.3. Running the Server
*(TODO: Command to start the FastAPI server using Uvicorn.)*
```bash
cd backend
source venv/bin/activate # If not already active
uvicorn src.main:app --reload --port 8000
```
Access the API at `http://localhost:8000` and interactive docs at `http://localhost:8000/docs`.

#### 5.2.4. Development Scripts
*(TODO: Scripts for linting, testing, running migrations.)*

### 5.3. VS Code Extension (TypeScript)
Located in the `/vscode-extension` directory.

#### 5.3.1. Environment Setup
*(TODO: `npm install` or `yarn install`.)*
```bash
cd vscode-extension
npm install
# or yarn install
```

#### 5.3.2. Building and Running in Development Mode
*(TODO: How to compile TypeScript and launch the extension in a new VS Code Extension Development Host window (usually F5).)*
- Open the `vscode-extension` folder in VS Code.
- Press `F5` or go to `Run > Start Debugging`.
- This will compile the extension and open a new VS Code window (Extension Development Host) with the AetherStack extension loaded.
- Test the `Hello from AetherStack` command from the command palette.

#### 5.3.3. Development Scripts
*(TODO: `package.json` scripts for linting, testing, compiling, packaging.)*

### 5.4. Web Dashboard (Next.js + TailwindCSS)
Located in the `/dashboard` directory.

#### 5.4.1. Environment Setup
*(TODO: `npm install` or `yarn install`. Setup for TailwindCSS and shadcn/ui.)*
```bash
cd dashboard
npm install
# or yarn install
# npx shadcn-ui@latest init (if not already done)
```

#### 5.4.2. Running the Development Server
*(TODO: Command to start the Next.js development server.)*
```bash
npm run dev
# or yarn dev
```
Access the dashboard at `http://localhost:3000`.

#### 5.4.3. Development Scripts
*(TODO: `package.json` scripts for linting, testing, building.)*

## 6. Inter-Module Communication

*(TODO: Overview of how the different parts of AetherStack talk to each other.)*

### 6.1. CLI to Backend
*(TODO: Typically HTTP requests to the Backend API. Configuration for backend URL.)*

### 6.2. VS Code Extension to Backend/CLI
*(TODO: Can make HTTP requests to Backend API or execute CLI commands as child processes.)*

### 6.3. Web Dashboard to Backend
*(TODO: HTTP requests from Next.js (client-side or server-side) to the Backend API.)*

### 6.4. Ports and Sockets
*(TODO: Default ports for backend (e.g., 8000), dashboard (e.g., 3000). Any other relevant network details.)*
- Backend API: `http://localhost:8000`
- Web Dashboard: `http://localhost:3000`

## 7. Documentation Setup (MkDocs)

The project documentation is located in the `/docs` directory and built using MkDocs with the Material theme.

### 7.1. Installation

To work with the documentation, you need to install MkDocs and the Material theme, along with any plugins specified in `mkdocs.yml`.

```bash
# Ensure Python and pip are installed
pip install mkdocs mkdocs-material pymdown-extensions mermaid2
# Or, if a requirements_docs.txt is provided:
# pip install -r docs/requirements_docs.txt
```

### 7.2. Building and Serving Docs Locally

Navigate to the root of the repository (`aetherstack/`) where the `mkdocs.yml` file is located (or to `docs/` if `mkdocs.yml` is there - adjust path as needed, current setup has `mkdocs.yml` in `docs/`).

```bash
cd docs  # If mkdocs.yml is in the docs directory
mkdocs serve
```
This will start a local development server, typically at `http://127.0.0.1:8001/`. The site will auto-reload when you save changes to the documentation files.

To build the static site (e.g., for deployment):
```bash
cd docs # If mkdocs.yml is in the docs directory
mkdocs build
```
The static files will be generated in the `docs/site/` directory.

## 8. Coding Standards and Conventions

*(TODO: Links to specific style guides, linting configurations (e.g., ESLint, Prettier, Pylint, GoFmt).)*

### 8.1. Linting and Formatting
*(TODO: Instructions on how to run linters and formatters for each module.)*

### 8.2. Commit Messages
*(TODO: Preferred commit message format, e.g., Conventional Commits.)*

## 9. Debugging

*(TODO: Tips and tools for debugging each component.)*

## 10. Sample Development Workflows

### 10.1. Adding a New CLI Command
*(TODO: High-level steps: define command in Cobra, implement logic, add tests.)*

### 10.2. Adding a New Backend API Endpoint
*(TODO: High-level steps: define Pydantic models, create FastAPI router/endpoint, implement service logic, add tests.)*

### 10.3. Developing a VS Code Extension Feature
*(TODO: High-level steps: define command/UI contribution in `package.json`, implement TypeScript logic, test in Extension Development Host.)*

## 11. Contribution Guidelines

Please refer to `CONTRIBUTING.md` (to be created) for detailed information on how to contribute, our code of conduct, and the process for submitting changes.

---
*Happy developing! If you encounter any issues or have suggestions, please open an issue on GitHub.*
