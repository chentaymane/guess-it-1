```md
# 🎯 Guess-It

Guess-It is a statistical guessing project where the goal is to predict a **tight range** for the next number in a sequence based on previously received values.

The smaller and more accurate your range is, the **higher your score**.

---

## 📌 Objective

Your program receives numbers **one by one** through standard input (`stdin`).  
For each number, you must print a range (`min max`) where you believe the **next number** will fall.

- ✅ Correct range → you score points  
- ❌ Incorrect range → zero points  
- 📉 Larger ranges → fewer points  

The challenge is to find the **best balance between precision and safety**.

---

## 📈 Input Format

Numbers are provided line by line:

```

189
113
121
114
145
110
...

```

- X-axis → line index (0, 1, 2, ...)
- Y-axis → value itself

Each number is an input for which you must **predict the range of the next number**.

---

## 📤 Output Format

For each input number, output **one line**:

```

<lower_bound> <upper_bound>

```

### Example

```

Input: 189
Output: 120 200

Input: 113
Output: 160 230

```

⚠️ Ranges do not need to be perfect every time — tighter correct ranges give more points.

---

## 🧠 Strategy

You are encouraged to use statistics from the **math-skills** project:

- Average (Mean)
- Median
- Variance
- Standard Deviation

There is **no required algorithm** — experimentation is part of the challenge.

---

## 🧪 Testing & Scoring

- Your output is compared to the **next input value**
- If the value is inside your range → you score
- Smaller ranges = higher score
- Performance matters (large datasets)

Test datasets:
- Data 1
- Data 2
- Data 3

---

## 🗂 Project Structure

```

guess-it/
├── ai/
│   ├── big-range
│   └── ...
├── student/
│   ├── script.sh
│   └── (your solution files)
├── index.html
├── index.js
└── docker-compose.yml

````

---

## 👨‍🎓 Student Folder (Required)

To be evaluated, you **must**:

1. Create a folder named `student/`
2. Add only the files needed to run your program
3. Add an executable script called `script.sh`

### Example `script.sh`

```sh
#!/bin/sh
go run ./student/main.go
````

⚠️ Missing any of these steps will cause the test to fail.

---

## 🚀 Running the Tester

Start the environment:

```sh
docker compose up
```

Open your browser:

👉 [http://localhost:3000](http://localhost:3000)

---

## 🤖 Adding a Guesser

The tester requires a second guesser from the `ai/` folder.

Add it via URL:

```
?guesser=<guesser_name>
```

Example:

```
http://localhost:3000/?guesser=big-range
```

---

## ⚡ Quick Mode

* **Quick** → skips waiting time
* **Clean** → clears results (recommended after each test)

---

## 🛠 Supported Languages

You may implement your solution in:

* Go (Golang)
* JavaScript
* Rust
* Python

---

## 📚 Learning Outcomes

This project helps you practice:

* Statistical calculations
* Probability & prediction
* Performance optimization
* Shell scripting
* Reading from standard input

---

## 🏁 Final Notes

* Smaller ranges give more points
* Wrong guesses give zero
* There is no perfect solution
* Improve by testing and analyzing results

Good luck and happy guessing 🎲

```
```
