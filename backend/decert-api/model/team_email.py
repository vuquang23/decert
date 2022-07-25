from pydantic import BaseModel, Field
from typing import List, Optional

class TeamEmailInCreate(BaseModel):
    id: str = Field(..., alias='_id')
    pattern: str

class TeamEmailInUpdate(BaseModel):
    id: str = Field(..., alias='_id')
    pattern: Optional[str]

class TeamEmailInDB(BaseModel):
    id: str = Field(..., alias='_id')
    pattern: str

class TeamEmails(BaseModel):
    results: List[TeamEmailInDB]
    total: int

class TeamEmailInDelete(BaseModel):
    id: str = Field(..., alias='_id')