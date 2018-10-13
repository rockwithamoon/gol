package gol

import "testing"

func TestSet_1(t *testing.T) {
	if v := SetLevel(INFO); !v {
		t.Error("Set failed with known level")
	}
}

func TestSet_2(t *testing.T) {
	SetTrace(true)
	if v := SetLevel(INFO); !v {
		t.Error("Set failed with known level")
	}
}

func TestSet_3(t *testing.T) {
	if v := SetLevel(100); v {
		t.Error("Set succeded on unkown level")
	}
}

func TestSet_4(t *testing.T) {
	SetTrace(false)
	if v, _ := Info("No trace here"); v > 21 {
		t.Error("Set failed to turn off trace")
	}
	SetTrace(true)
}

func TestFatal_Yes(t *testing.T) {
	SetLevel(FATAL)
	if v, _ := Fatal("Fatal message %v", 1); v == 0 {
		t.Error("Expected fatal message")
	}
}

func TestFatal_No(t *testing.T) {
	SetLevel(ERROR)
	if v, _ := Fatal("Can't rid off fatal messages %v", 2); v == 0 {
		t.Error("Stil expected fatal message")
	}
}

func TestError_Yes(t *testing.T) {
	SetLevel(ERROR)
	if v, _ := Error("Error message"); v == 0 {
		t.Error("Expected error message")
	}
}

func TestError_No(t *testing.T) {
	SetLevel(FATAL)
	if v, _ := Error("Error message"); v != 0 {
		t.Error("Did not expected error message")
	}
}

func TestWarn_Yes(t *testing.T) {
	SetLevel(WARN)
	if v, _ := Warn("Warn message"); v == 0 {
		t.Error("Expected warn message")
	}
}

func TestWarn_No(t *testing.T) {
	SetLevel(ERROR)
	if v, _ := Warn("Warn message"); v != 0 {
		t.Error("Did not expected warn message")
	}
}

func TestInfo_Yes(t *testing.T) {
	SetLevel(INFO)
	if v, _ := Info("Info message"); v == 0 {
		t.Error("Expected info message")
	}
}

func TestInfo_No(t *testing.T) {
	SetLevel(WARN)
	if v, _ := Info("Info message"); v != 0 {
		t.Error("Did not expected info message")
	}
}

func TestDebug_Yes(t *testing.T) {
	SetLevel(DEBUG)
	if v, _ := Debug("Debug message"); v == 0 {
		t.Error("Expected debug message")
	}
}

func TestDebug_No(t *testing.T) {
	SetLevel(INFO)
	if v, _ := Debug("Debug message"); v != 0 {
		t.Error("Did not expected debug message")
	}
}

func TestTrace_Yes(t *testing.T) {
	SetLevel(TRACE)
	if v, _ := Trace("Trace message"); v == 0 {
		t.Error("Expected trace message")
	}
}

func TestTrace_No(t *testing.T) {
	SetLevel(DEBUG)
	if v, _ := Trace("Trace message"); v != 0 {
		t.Error("Did not expected trace message")
	}
}
