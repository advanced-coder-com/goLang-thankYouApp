# ThankYou console App (Ubuntu)

This app generates Thanksgiving in random language.

## How to install

Just download compiled file `thankYou` (or compile from source) and make alias 


For example 

1) run `nano ~/.bashrc`

2) add `alias ThankYou='~/ThankYouApp/thankYou'`

3) save file and apply `source ~/.bashrc`

4) Run created command `ThankYou`

## Usage

Running ./thankYou or your custom alis command App will generate "thank you" in one random language.

### Other functions

#### Specific language

Run `./thankYou` with language code argument like `./thankYou It` Result will be in Italian language

#### Add new language

Run `./thankYou add` without arguments and follow instructions

#### Remove language

Run `./thankYou` delete with language code argument. `./thankYou delete It`. Italian language will be removed from database


