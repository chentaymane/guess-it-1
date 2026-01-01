```markdown
# Guess-It

A statistical prediction program that uses mathematical analysis to predict the range of upcoming numbers in a data sequence.

## 📋 Overview

Guess-It is a number prediction system that analyzes streaming numerical data and predicts ranges for future values. The program processes numbers line-by-line from standard input and outputs predicted ranges for the next number, balancing accuracy with range size for optimal scoring.

## 🎯 Objective

Given a sequence of numbers, the program must:
- Analyze incoming numerical data in real-time
- Calculate statistical patterns and trends
- Predict a range (lower and upper bounds) for the next number
- Optimize for the smallest accurate range to maximize score

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- A web browser
- Your implementation in the `student/` folder

### Project Structure


guess-it/
├── ai/                    # AI guesser implementations
│   ├── big-range
│   └── ...
├── student/               # Your solution goes here
│   ├── your_program.*     # Your implementation
│   └── script.sh          # Executable script to run your program
├── index.html
├── index.js
└── docker-compose.yml

## 🛠️ Installation & Usage

### 1. Setup Your Solution

Create a `student/` folder with:
- Your program files
- An executable `script.sh` that runs your program from the root directory

**Example `script.sh`:**
```bash
#!/bin/sh
node ./student/solution.js
```

Make it executable:
```bash
chmod +x student/script.sh
```

### 2. Start the Application

```bash
docker compose up
```

### 3. Access the Web Interface

Open your browser at [http://localhost:3000](http://localhost:3000)

### 4. Add a Guesser

Add a guesser parameter to the URL:
```
http://localhost:3000/?guesser=big-range
```

### 5. Run Tests

- Select a test data set button
- Click `Quick` to skip animations and see results immediately
- Use `Clean` button to clear displays between tests

## 📊 How It Works

### Input/Output Format

**Input:** Numbers are provided one per line via standard input

**Output:** For each number, output a range: `lower_limit upper_limit`

**Example:**
```
$ ./your_program
189                 # Input
120 200             # Your predicted range for next number (113)
113                 # Input
160 230             # Your predicted range for next number (121)
121                 # Input
110 140             # Your predicted range for next number (114)
114                 # Input
100 200             # Your predicted range for next number (145)
...
```

### Data Representation

The data represents a graph where:
- **X-axis:** Line numbers (0, 1, 2, 3, 4, 5, ...)
- **Y-axis:** Actual values (189, 113, 121, 114, 145, 110, ...)

## 🏆 Scoring System

- **Correct predictions:** Points awarded inversely proportional to range size
- **Smaller ranges:** Higher scores (if correct)
- **Larger ranges:** Lower scores (even if correct)
- **Balance:** Find the optimal trade-off between accuracy and range size

### Scoring Formula

```
Score increases when prediction is correct
Score increment is inversely proportional to range size
Smaller range = Higher score (if correct)
Incorrect prediction = No score increment
```

## 💻 Supported Languages

- **Go** (Golang)
- **JavaScript** (Node.js)
- **Rust**
- **Python**

## 🧪 Testing

### Test Datasets

The program is tested against multiple datasets:
- **Data 1**
- **Data 2**
- **Data 3**

### Running Tests

1. Download the official tester (zip file provided in project instructions)
2. Place your `student/` folder in the root directory
3. Follow the tester's instructions
4. Verify your `script.sh` works correctly from the root folder

### Testing Requirements

Your submission must include:
1. A folder named `student/`
2. All files needed to run your program
3. An executable shell script `script.sh` containing run commands

**Important:** The script must be executable and work from the root folder of the tester.

## 📈 Key Concepts

This project helps you learn about:

### Statistical Analysis
- Mean, median, mode
- Standard deviation
- Variance
- Moving averages

### Probability Calculations
- Confidence intervals
- Distribution analysis
- Prediction intervals
- Probability density functions

### Pattern Recognition
- Trend detection
- Seasonality analysis
- Anomaly detection
- Time series analysis

### Algorithm Optimization
- Balancing accuracy vs. range size
- Real-time data processing
- Performance optimization
- Memory efficiency

## 🎓 Learning Outcomes

- **Statistical and Probability Calculation**
- **Data Analysis and Pattern Recognition**
- **Algorithm Design and Optimization**
- **Shell Scripting**
- **Real-time Data Processing**

## 📦 Deliverables

Your submission should include:

```
student/
├── your_program_files
└── script.sh              # Executable shell script
```

## ⚡ Performance Considerations

- **Speed:** Program will be extensively tested - prioritize performance
- **Memory:** Handle large datasets efficiently
- **Accuracy:** Balance between range size and prediction accuracy
- **Real-time:** Process streaming data without delays

## 🔗 Repository

Project Repository: [https://learn.zone01oujda.ma/git/achent/guess-it-1](https://learn.zone01oujda.ma/git/achent/guess-it-1)

## 🛠️ Troubleshooting

### Common Issues

**Script doesn't run:**
- Ensure `script.sh` is executable: `chmod +x student/script.sh`
- Verify the shebang line is correct: `#!/bin/sh`
- Check that paths are relative to root folder

**Docker issues:**
- Ensure Docker daemon is running
- Check port 3000 is not already in use
- Try `docker compose down` then `docker compose up`

**Browser console errors:**
- Clear browser cache
- Ensure guesser parameter is in URL
- Check Dev Tools console for specific errors

## 📚 Additional Resources

- Review the math-skills exercise for statistical foundations
- Study time series analysis techniques
- Research prediction interval calculation methods
- Learn about confidence intervals and their applications

## 📝 Notes

- The web interface uses large datasets - use the `Clean` button after each test
- Test locally with the provided tester before final submission
- Focus on finding the optimal balance between range size and accuracy
- Performance matters - optimize your calculations

## 📄 License

Educational project for Zone01 Oujda

---

**Good luck with your implementation! May your predictions be accurate and your ranges be small! 🎯**
```
