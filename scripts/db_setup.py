import os

if os.path.isfile('./local-dev-env/.env-local'):
  print('.env-local file already exists. Skipping env creation step. To modify env, modify it manually or delete it and rerun this script.')
  exit()

username = input('Enter a username. Leave blank for default. (default postgres)')
if not len(username):
  username = 'postgres'
password = input('Enter a password.')
if not len(password):
  raise Exception('Password is required')
database = input('Enter a database name. Leave blank for default. (default pokedex)')
if not len(database):
  database = 'pokedex'

with open('./local-dev-env/.env-local') as f:
  f.write(f'POSTGRES_USER={username}\nPOSTGRES_PASSWORD={password}\nPOSTGRES_DB={database}')

print('Successfully created .env-local file!')