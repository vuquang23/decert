from pydantic import BaseModel, Field
from typing import List

class CountryInCreate(BaseModel):
    id: str = Field(..., alias='_id')
    value: int
    code: str

class CountryInUpdate(BaseModel):
    value: int
    code: str

class CountryInDB(BaseModel):
    id: str = Field(..., alias='_id')
    value: int
    code: str

class Countries(BaseModel):
    results: List[CountryInDB]
    total: int

class CountryInDelete(BaseModel):
    id: str = Field(..., alias='_id')