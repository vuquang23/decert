from fastapi.security import SecurityScopes, HTTPBearer
from starlette import status
from fastapi import HTTPException, Security
from ..core.config import SECRET_KEY
from ..core import jwt

reusable_oauth2 = HTTPBearer(scheme_name="Authorization")


def get_curent_user(
    security_scopes: SecurityScopes,
    token: str = Security(reusable_oauth2),
):
    try:
        token_data = jwt.get_payload_from_token(
            token.credentials,
            str(SECRET_KEY),
        )
    except ValueError as e:
        raise HTTPException(
            status_code=status.HTTP_403_FORBIDDEN,
            detail="could not validate credentials",
        ) from e

    if security_scopes.scopes and not token_data.role:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Not enough permissions",
        )
    if security_scopes.scopes and token_data.role not in security_scopes.scopes:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Not enough permissions",
        )
