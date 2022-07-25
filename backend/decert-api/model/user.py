from pydantic import BaseModel
from typing import Optional, List


class UserInCreate(BaseModel):
    username: str
    password: str
    role: str


class UserInUpdate(BaseModel):
    username: str
    password: Optional[str]
    role: Optional[str]


class UserInLogin(BaseModel):
    username: str
    password: str


class UserInDelete(BaseModel):
    username: str


class UserInDB(BaseModel):
    username: str
    password: str
    role: str


class Users(BaseModel):
    data: List[UserInDB]
