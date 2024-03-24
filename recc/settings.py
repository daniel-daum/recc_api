import os

from dotenv import load_dotenv
from pydantic_settings import BaseSettings

load_dotenv()


class Settings(BaseSettings):
	"""Loads Recc API settings and envrionment variables."""

	port: int = int(os.getenv('PORT'))
	version: str = os.getenv('VERSION')


settings = Settings()
