from litestar import Litestar, Router
from litestar.openapi import OpenAPIConfig

from recc.routes import health
from recc.settings import settings

base_router = Router(path='/api', route_handlers=[health])

app = Litestar([base_router], openapi_config=OpenAPIConfig(title='Recc API', version=settings.version))
