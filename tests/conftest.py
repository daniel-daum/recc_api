import pytest

from litestar.testing import TestClient

from recc.server import app


@pytest.fixture(scope="function")
def test_client() -> TestClient:
    return TestClient(app=app)
