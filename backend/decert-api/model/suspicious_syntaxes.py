from pydantic import BaseModel, Field
from typing import List, Optional

class SuspiciousSyntaxInDB(BaseModel):
    id: str=Field(alias="_id")
    name: str
    pattern: str

class SuspiciousSyntax(BaseModel):
    results: List[SuspiciousSyntaxInDB]
    total: int

class SuspiciousSyntaxInCreate(BaseModel):
    name: str
    pattern: str

class SuspiciousSyntaxInUpdate(BaseModel):
    name: Optional[str]
    pattern: Optional[str]
    
class SuspiciousSyntaxInDelete(BaseModel):
    id: str=Field(alias="_id")