# KraickList

Welcome to Haraj take home challenge!

In this repository you will find simple web app for fictional startup called KraickList. This app will allow users to search ads from given sample data located in `data.gz`.

Currently the app is just rough prototype. The search is case sensitive, limited to exact matches, & the search result is pretty much could be further improved.

You could see the live version of this app [here](https://guarded-bastion-39032.herokuapp.com/). Try searching for "iPhone" to see some results.

## Installation

- You could see the live version of this app [here](https://guarded-bastion-39032.herokuapp.com/)

OR

- Clone this repository

```bash
git clone gitgit@github.com:hilmihi/kraicklist.git
cd kraicklist
```

- Run the setup command to setup the containers.

```bash
make setup
```

- after that can check [http://127.0.0.1:3001/](http://127.0.0.1:3001/)

- Run the setup command to stop/destroy container.

```bash
make destroy
```

## Run unit test

```bash
make test
```

## Changes

1. Change FE a little using bootstrap.
2. Add func for specific search type. Search only by it's Title, Tags, Content or All.
3. Remove case sensitive when searching.
4. When we search more than one keyword, the result is based on when words we input is exist not only when our input have exact matched.
5. Sorting result based on how many keyword exist.
6. Add unit testing.

## About Sample Data

The data is translated from Arabic to English using Google Translate, so sometimes you will find funny translation on it. 🤣
