from fastapi.security import HTTPBearer
from fastapi import Depends, HTTPException
from datetime import datetime, timedelta
from pydantic import ValidationError
from typing import Optional
import jwt
from .config import SECRET_KEY, ALGORITHM
from ..model.jwt import JWTUser

access_token_jwt_subject = "access"

reusable_oauth2 = HTTPBearer(scheme_name="Authorization")


def generate_token(
    username: str, role: str, expires_delta: Optional[timedelta] = None
) -> str:
    if expires_delta:
        expire = datetime.utcnow() + timedelta(minutes=expires_delta)
    else:
        expire = datetime.utcnow() + timedelta(minutes=15)

    to_encode = {"exp": expire, "username": username, "role": role}
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
    return encoded_jwt


def get_payload_from_token(token: str, secret_key: str) -> JWTUser:
    try:
        return JWTUser(**jwt.decode(token, secret_key, algorithms=[ALGORITHM]))
    except jwt.PyJWTError as decode_error:
        raise ValueError("unable to decode JWT token") from decode_error
    except ValidationError as validation_error:
        raise ValueError("malformed payload in token") from validation_error
