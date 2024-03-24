from litestar import MediaType, get


@get(path='/health', tags=['health'], media_type=MediaType.JSON, description='API health check endpoint.')
async def health() -> dict[str, str]:
	return {'status': 'healthy'}
