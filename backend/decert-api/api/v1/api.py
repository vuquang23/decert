from fastapi import APIRouter

from .endpoints.cert import router as cert_router
from .endpoints.collection import router as collection_router

api_router = APIRouter()
api_router.include_router(cert_router)
api_router.include_router(collection_router)