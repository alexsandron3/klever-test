# Upvote service

Simple microservice to upvote and/or downvote.

---
[Here](https://github.com/alexsandron3/klever-test/pulls?q=is%3Apr+is%3Aclosed) you can see my code evolution easily.

# How to run this project

- Clone this repository

```bash
   git clone git@github.com:alexsandron3/klever-test.git
```

- Acess the folder, download depencies and generate proto 

```bash
   cd klever-test && make generate
```

- Create .env and fill with your mongo credentials

- Run gRPC server

```bash
   make run
```

---