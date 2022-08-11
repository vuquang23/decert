from pydantic import BaseModel

class EmailGroupByModel(BaseModel):
    legit: int
    phishing: int

# class EmailGroupByHuman(BaseModel):


class StatisticsInResponse(BaseModel):
    results: EmailGroupByModel