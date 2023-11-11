package cli

import (
  "QuantumMango/mathGame/expr"
  "fmt"
  "strings"
  "strconv"
  "github.com/charmbracelet/bubbles/textarea"
  tea "github.com/charmbracelet/bubbletea"
)


var streakStatements = []string{
  "",
  "",
  "Three in a row!",
  "Quadrouple in a row?!?",
  "5 in a row! Impressive!",
  "6 times in a row, you a calculator?",
  "7... I'm speechless",
  "8 in a row! You're on fire!",
  "9 consecutive successes!",
  "10 times! Keep it up!",
  "11 in a row! Unbelievable!",
  "A dozen in a row! Outstanding!",
  "Lucky 13! Keep the streak!",
  "14 times! You're unstoppable!",
  "Fifteen in a row! Amazing!",
  "Sweet 16! Keep the momentum!",
  "Seventeen times! Incredible!",
  "18 consecutive wins!",
  "Nineteen in a row! Wow!",
  "20 times! You're a champion!",
  "21 in a row! Legendary streak!",
  "22 straight successes!",
  "Twenty-three times! Phenomenal!",
  "24 consecutive wins!",
  "Twenty-five in a row! Remarkable!",
  "26 times! You're a superstar!",
  "27 consecutive successes!",
  "Twenty-eight in a row! Outstanding!",
  "29 times! Keep the streak alive!",
  "30 consecutive wins! Incredible!",
  "31 in a row! Unstoppable!",
  "32 straight victories! Amazing!",
  "33 times! You're a true champion!",
  "34 consecutive successes!",
  "Thirty-five in a row! Phenomenal!",
  "36 times! Keep up the great work!",
  "37 consecutive wins! You're incredible!",
  "38 in a row! Remarkable!",
  "Thirty-nine times! You're on fire!",
  "40 consecutive successes! You're a legend!",
  "I ran out of ideas, congrats",
}

type Game struct {
  question string
  answer int 
  status string
  qIndex int
  streak int
}


func newGameDefault() Game {
  question:= expr.NewExpr(1,3);
  return Game {
    question: question,
    answer: expr.SolveExpr(question),
    status: "",
    qIndex: 1,
  }
}

func getNewAnswer(s string) (int, error) {
  entries := strings.Split(s, "\n")

  // Find the last non-empty entry
  var lastEntry string
  for i := len(entries) - 1; i >= 0; i-- {
    if entries[i] != "" {
      lastEntry = entries[i]
      break
    }
  }

  // Convert the last entry to an integer
  result, err := strconv.Atoi(lastEntry)
  if err != nil {
    return 0, err
  }

  return result, nil}

func (g Game) getQuestion() string {
  return g.question
}

type errMsg error;

type InputModel struct {
  textarea textarea.Model
  err error
  game Game

}

func (m InputModel) Init() tea.Cmd {
  return textarea.Blink
}

func NewInputModel() InputModel {
  ti:= textarea.New();
  ti.Placeholder = "Enter the answer here"
  ti.Focus()
  ti.SetWidth(30);
  ti.SetHeight(1);
  ti.ShowLineNumbers = false;

  
  return InputModel {
    textarea: ti,
    err: nil,
    game: newGameDefault(),

  }
}

func (m InputModel) Update(msg tea.Msg) (tea.Model,tea.Cmd) {
  var cmds []tea.Cmd
  var cmd tea.Cmd

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.Type {
    case tea.KeyEsc:
      if m.textarea.Focused() {
        m.textarea.Blur()
      }
    case tea.KeyCtrlC:
      return m, tea.Quit
    case tea.KeyEnter:
      num,err := getNewAnswer(m.textarea.Value());
      if err != nil {
        m.game.status = "Invalid Input!";
        break;
      }
      m.textarea.Reset();
      if expr.SolveExpr(m.game.question) == num {
              q:=  expr.NewExpr((m.game.qIndex+1)%10-1,9);
        a := expr.SolveExpr(q);
        m.game.question = q;
        m.game.answer = a;
        m.game.qIndex +=1;
        m.game.streak += 1;
        m.game.status = "Correct!";
      }else {
        m.game.streak = 0;
        m.game.status = "Oops, wrong answer!";
      }
    default:
      if !m.textarea.Focused() {
        cmd = m.textarea.Focus()
        cmds = append(cmds, cmd)
      }
    }

  // We handle errors just like any other message
  case errMsg:
    m.err = msg
    return m, nil
  }

  m.textarea, cmd = m.textarea.Update(msg)
  cmds = append(cmds, cmd)
  return m, tea.Batch(cmds...)
}

func (m InputModel) View() string {
  g := m.game;


  streakStatement := "";
  if g.streak > 40 {
    streakStatement = "I ran out of sentence ideas, your streak: " + strconv.Itoa(g.streak);
  }else if g.streak-1 > 0 {
    streakStatement = streakStatements[g.streak-1];
  }else{
    streakStatement = streakStatements[g.streak]
  }
  
 

  return fmt.Sprintf(
    "Math Game. \nNote: Consider only integers. Ignore all fractional Parts. \n Example: 0.1+1=1 (ignore the frations)\n\nQuestion:\n%s\n%s\n%s\nYou Solved: %d question(s)\n%s\n\n%s",
    g.getQuestion(),
    m.textarea.View(),
    m.game.status,
    m.game.qIndex-1,
    streakStatement,
    "(ctrl+c to quit)",
    )+ "\n\n"

}


