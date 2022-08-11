from fastapi import APIRouter, Depends, Query, Path, Security

router = APIRouter()

SAMPLE_COLLECTION = {
    "collectionName": "quangdz",
}


@router.get(
    "/collection/list", 
    tags = ["Collection"]
)
def collection_list(*, collection_id: str, offset: int, limit: int):
    return [SAMPLE_COLLECTION, SAMPLE_COLLECTION]
    
@router.get(
    "/collection/info", 
    tags = ["Collection"]
)
def cert_verify(*, cert_id: str):
    return SAMPLE_COLLECTION