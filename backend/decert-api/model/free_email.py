from typing import Optional
from pydantic import BaseModel

class FreeInCreate(BaseModel):
    name: str
    domain: str

class FreeInGetAll(BaseModel):
    offset: int
    size: int

class FreeInGetOne(BaseModel):
    domain: str

class FreeInUpdate(BaseModel):
    name: str
    pattern: Optional[str]

class FreeInDelete(BaseModel):
    domain: str
