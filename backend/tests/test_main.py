from fastapi.testclient import TestClient
from sqlmodel import Session, SQLModel, create_engine
from ..src.main import app, AnalysisJobCreate # Adjust import based on your structure

# Use an in-memory SQLite database for testing
DATABASE_URL_TEST = "sqlite:///./test.db"
engine_test = create_engine(DATABASE_URL_TEST, echo=False)

def create_db_and_tables_test():
    SQLModel.metadata.create_all(engine_test)

def drop_db_and_tables_test():
    SQLModel.metadata.drop_all(engine_test)

client = TestClient(app)

def test_healthcheck():
    create_db_and_tables_test() # Ensure tables are created if startup event doesn't run for TestClient context
    response = client.get("/healthcheck")
    assert response.status_code == 200
    assert response.json() == {"status": "ok", "message": "AetherStack Backend API is healthy"}
    drop_db_and_tables_test()

def test_request_analysis_placeholder():
    create_db_and_tables_test()
    job_data = {"repo_url": "https://github.com/test/repo.git", "status": "pending"}
    response = client.post("/analyze", json=job_data)
    assert response.status_code == 202 # Accepted
    data = response.json()
    assert data["repo_url"] == job_data["repo_url"]
    assert data["status"] == "pending"
    assert "id" in data
    # TODO: Add more assertions once the endpoint is fully implemented
    drop_db_and_tables_test()

# TODO: Add more tests for other endpoints and functionalities

