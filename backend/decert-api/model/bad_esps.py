from typing import Optional
from pydantic import BaseModel

class BadEspsInCreate(BaseModel):
    name: str
    domain: str

class BadEspsInGetAll(BaseModel):
    offset: int
    size: int

class BadEspsInGetOne(BaseModel):
    domain: str

class BadEspsInUpdate(BaseModel):
    name: str
    pattern: Optional[str]

class BadEspsInDelete(BaseModel):
    domain: str
