from fastapi import APIRouter, Depends, Query, Path, Security

router = APIRouter()

SAMPLE_CERT = {
        "certName": "quangdz",
        "issuer": {
            "name": "quangdz",
            "wallet": "quangdz",
        },
        "receiver": {
            "name": "quangdz",
            "wallet": "quangdz",
            "dateOfBirth": "hehe",
        },
        "description": "anh quang dz",
        "issuanceDate": "hom nay",
        "expiredAt": "ngay mai",
        "proof": {
            "nftAddress": "khong co dau",
            "nftID": "cung khong co dau"
        }
    }


@router.post(
    "/cert/list", 
    tags = ["Certification"]
)
def cert_list(*, collection_id: str, offset: int, limit: int):
    return [SAMPLE_CERT, SAMPLE_CERT]

@router.get(
    "/cert/verify", 
    tags = ["Certification"]
)
def cert_verify(*, cert_id: str):
    return SAMPLE_CERT

@router.post(
    "/cert/store", 
    tags = ["Certification"]
)
def cert_issue(*, issuer, receiver):
    return SAMPLE_CERT

@router.post(
    "/cert/update", 
    tags = ["Certification"]
)
def cert_update(*, tx_hash: str):
    return SAMPLE_CERT
