import os
from decouple import config

API_V1_STR = "/api/v1"
ROOT_PATH_CMS_API = config("ROOT_PATH_CMS_API")
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES=60*24*7

MONGODB_URI = config("MONGODB_CONNECTION_STR")
PROJECT_NAME = config("MONGODB_METADATA_COLLECTION")
SECRET_KEY = config("SECRET_KEY")
TRUST_RULE_API_BASE_PATH = config("TRUST_RULE_API_BASE_PATH")

ACCESS_ID = config("ACCESS_ID")
SPACE_ACCESS_KEY = config("SPACE_SECRET_KEY")
ENDPOINT_URL = config("ENDPOINT_URL")
REGION_NAME = config("REGION_NAME")

admin_permission = ["admin.read", "admin.write"]
roles = ["admin", "NORMAL"]
NORMAL_permission = ["NORMAL.read"]
allow_create_resource=None
allow_read_resource=None

database_name = PROJECT_NAME
email_database_name = "ScamAdviser"
anonymous_collection = "anonymous_email_providers"
disposable_collection = "disposable_email_providers"
free_collection = "free_email_providers"
good_esps = "good_reputation_esps"
bad_esps = "bad_reputation_esps"
countries_collection = "spam_countries"
suspicius_syntaxes_collection = "suspicious_syntaxes"
team_email_collections = "team_email_patterns"
well_known_companies_collection = "well_known_companies"
label_email_collection = "labelEmail"
unspam_email_collection = "unspam"
tld_collection = "tld"

users_database_name = "Users"
users_collection = "Users"
email_display_collection = "scamEmail"