# Usage of DBeaver Team Edition GraphQL API with Python 3

How to start:
1. Change the current directory to the directory with this README file.
2. Initialize a virtual environment with
```sh
python3 -m venv .venv
```
3. Source the activation script to start using it. 
    * If you use bash, execute
      ```sh
      source .venv/bin/activate
      ```
    * If you use fish, execute
      ```sh
      source .venv/bin/activate.fish
      ```
4. Install the required dependencies with
```sh
pip install -r requirements.txt
```
5. Create the `../.env` file from the `../.env.template` (see the repository root for file `.env.template`)
```sh
cp ../.env.template ../.env
```
6. Fill the environment variables in the `.env` file.
7. Execute the script with
```sh
python3 te.py
```
