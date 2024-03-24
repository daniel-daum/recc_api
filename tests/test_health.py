from litestar.status_codes import HTTP_200_OK
from litestar.testing import TestClient

def pytest_health() -> str:
    return "pytest is healthy!"

def test_pytest_health():
    got = pytest_health()
    want = "pytest is healthy!"

    assert got == want

def test_health_check_with_fixture(test_client: TestClient) -> None:
    with test_client as client:
        response = client.get("/api/health")
        assert response.status_code == HTTP_200_OK
        assert response.json() == {'status': 'healthy'}
