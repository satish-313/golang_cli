package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "new Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskname := "new task"
	l.Add(taskname)

	if l[0].Task != taskname {
		t.Errorf("Expected %q, got %q instead.", taskname, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New task 1", "New task 2", "New task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}
	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected %d, got %q instead", 2, len(l))
	}

	if tasks[2] != l[1].Task {
		t.Errorf("Expected %q,got %q instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskname := "new task"
	l1.Add(taskname)

	if l1[0].Task != taskname {
		t.Errorf("Expected %q, got %q instead", taskname, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error in saving file: %s ", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error in getting file: %s ", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q", l1[0].Task, l2[0].Task)
	}
}
