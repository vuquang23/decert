from typing import Optional
from pydantic import BaseModel

class DisposableInCreate(BaseModel):
    name: str
    domain: str

class DisposableInGetAll(BaseModel):
    offset: int=0
    size: int=10

class DisposableInGetOne(BaseModel):
    domain: str

class DisposableInUpdate(BaseModel):
    name: str
    pattern: Optional[str]

class DisposableInDelete(BaseModel):
    domain: str

