package logic

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	db "github.com/z3co/prot/db/gen"
)

type Instance struct {
	Store *db.Queries
}

func getFolderAndBranch() (db.CreateListParams, error) {
	folder, err := os.Getwd()
	out, errExec := exec.Command("git", "branch", "--show-current").Output()
	if errExec != nil {
		return db.CreateListParams{Folder: folder, Branch: "main"}, nil
	}
	if err != nil {
		return db.CreateListParams{}, fmt.Errorf("could not get working dir: %s", err)
	}
	branch := strings.TrimSpace(string(out))
	return db.CreateListParams{
		Folder: folder,
		Branch: branch,
	}, nil
}

func (s *Instance) CreateList(ctx context.Context, git bool) error {
	folder, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get working dir: %s", err)
	}
	_, err = s.Store.CreateList(ctx, db.CreateListParams{Folder: folder, Branch: "main"})
	if err != nil {
		return fmt.Errorf("could not create list: %s", err)
	}
	return nil
}

func (s *Instance) CreateListGit(ctx context.Context) error {
	folder, err := os.Getwd()
	out, errExec := exec.Command("git", "branch", "--show-current").Output()
	if errExec != nil {
		return fmt.Errorf("could not get git branch: %s", errExec)
	}
	if err != nil {
		return fmt.Errorf("could not get working dir: %s", err)
	}
	_, err = s.Store.CreateList(ctx, db.CreateListParams{Folder: folder, Branch: strings.TrimSpace(string(out))})
	if err != nil {
		return fmt.Errorf("could not create list: %s", err)
	}
	return nil
}

func (s *Instance) CreateTodo(ctx context.Context, description string) error {
	listInfo, err := getFolderAndBranch()
	if err != nil {
		return err
	}
	id, err := s.Store.GetListIdByFolderBranch(ctx, db.GetListIdByFolderBranchParams(listInfo))
	if err != nil {
		return fmt.Errorf("cannot find list for folder and branch: %s", err)
	}
	_, err = s.Store.CreateTodo(ctx, db.CreateTodoParams{
		Description: description,
		ListID:      id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Instance) GetTodos(ctx context.Context) ([]db.Todo, error) {
	listInfo, err := getFolderAndBranch()
	if err != nil {
		return []db.Todo{}, err
	}
	id, err := s.Store.GetListIdByFolderBranch(ctx, db.GetListIdByFolderBranchParams(listInfo))
	if err != nil {
		return []db.Todo{}, fmt.Errorf("cannot find list for folder and branch: %s", err)
	}
	todos, err := s.Store.GetTodosByListId(ctx, id)
	return todos, err
}

func (s *Instance) UpdateStatus(ctx context.Context, id int64) error {
	status, err := s.Store.GetTodoStatusById(ctx, id)
	done := status == 1
	if err != nil {
		return fmt.Errorf("could not fine todo with id %v", id)
	}
	if done {
		err := s.Store.UpdateTodoStatus(ctx, db.UpdateTodoStatusParams{ID: id, Done: 0})
		if err != nil {
			return fmt.Errorf("could change todo status", err)
		}
		return nil
	}
	err = s.Store.UpdateTodoStatus(ctx, db.UpdateTodoStatusParams{ID: id, Done: 1})
	if err != nil {
		return fmt.Errorf("could change todo status", err)
	}
	return nil
}
