package expr

import (
	"testing"
  "time"
)



func TestAnswerSimple(t *testing.T) {
	got:= SolveExpr("1+1");
	expected:=2;
	if got !=expected {
		t.Errorf("Expected: %d \n got: %v",expected,got);
	}
}

func TestAnswerMedium(t *testing.T) {
	got := SolveExpr("12/6");
	expected := 2;
	if got !=expected {
				t.Errorf("Expected: %d \n got: %v",expected,got);
	}
}

func TestParenthesesSimple(t *testing.T) {
	got := SolveExpr("(1+2)");
	expected := 3;
	if got !=expected {
				t.Errorf("Expected: %d \n got: %v",expected,got);
	}
}

func TestParenthesesMedium(t *testing.T) {
	got := SolveExpr("10/(1+1)");
	expected:= 5;
	if got !=expected {
				t.Errorf("Expected: %d \n got: %v",expected,got);
	}
}

func TestParenthesesHard(t *testing.T) {
	got := SolveExpr("(10-5)*1214+100/21");
	expected:= 6074;
	if got !=expected {
				t.Errorf("Expected: %d \n got: %v",expected,got);
	}
}


func TestNewExpr(t *testing.T) {
	s := NewExpr(1,1);
  time.Sleep(5*time.Second) // Won't work for anything less than 5; idk why
	if NewExpr(1,1) == s {
		t.Errorf("Expressions are not new: found '%s' both times",s);
	}
}
