from lib2to3.pytree import Base
from optparse import Option
from typing import Optional, List
from pydantic import BaseModel, Field

class CompanyInDB(BaseModel):
    id: str=Field(alias="_id")
    name: str
    pattern: str

class ListCompanies(BaseModel):
    results: List[CompanyInDB]
    total: int

class CompanyInCreate(BaseModel):
    domain: str
    name: str

class CompanyInUpdate(BaseModel):
    name: Optional[str]
    pattern: Optional[str]

class CompanyInDelete(BaseModel):
    domain: str


