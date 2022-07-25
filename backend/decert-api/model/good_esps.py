from typing import Optional
from pydantic import BaseModel

class GoodEspsInCreate(BaseModel):
    name: str
    domain: str

class GoodEspsInGetAll(BaseModel):
    offset: int
    size: int

class GoodEspsInGetOne(BaseModel):
    domain: str

class GoodEspsInUpdate(BaseModel):
    name: str
    pattern: Optional[str]

class GoodEspsInDelete(BaseModel):
    domain: str
