package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/plumbing" // Removed unused import
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status of the current repository",
	Long:  `Displays paths that have differences between the index file and the current HEAD commit, paths that have differences between the working tree and the index file, and paths in the working tree that are not tracked by Git.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := "." // Current directory
		r, err := git.PlainOpen(path)
		if err != nil {
			if err == git.ErrRepositoryNotExists {
				fmt.Println("Error: Not a git repository (or any of the parent directories).")
				os.Exit(1)
			}
			fmt.Printf("Error opening repository: %s\n", err)
			os.Exit(1)
		}

		w, err := r.Worktree()
		if err != nil {
			fmt.Printf("Error getting worktree: %s\n", err)
			os.Exit(1)
		}

		// Get current branch
		head, err := r.Head()
		if err != nil {
			fmt.Printf("Error getting HEAD: %s\n", err)
			os.Exit(1)
		}

		if head.Name().IsBranch() {
			fmt.Printf("On branch %s\n", head.Name().Short())
		} else {
			fmt.Printf("HEAD detached at %s\n", head.Hash().String()[:7])
		}

		// Get repository status
		status, err := w.Status()
		if err != nil {
			fmt.Printf("Error getting status: %s\n", err)
			os.Exit(1)
		}

		if status.IsClean() {
			fmt.Println("nothing to commit, working tree clean")
			return
		}

		fmt.Println("\nChanges to be committed:")
		// Staged changes (Index to HEAD)
		stagedChanges := false
		for path, fileStatus := range status {
			if fileStatus.Staging == git.Added || fileStatus.Staging == git.Modified || fileStatus.Staging == git.Deleted || fileStatus.Staging == git.Renamed {
				fmt.Printf("\t%s: %s\n", getStatusText(fileStatus.Staging), path)
				stagedChanges = true
			}
		}
		if !stagedChanges {
			fmt.Println("\t(no changes staged)")
		}

		fmt.Println("\nChanges not staged for commit:")
		// Unstaged changes (Worktree to Index)
		unStagedChanges := false
		for path, fileStatus := range status {
			if fileStatus.Worktree == git.Modified || fileStatus.Worktree == git.Deleted {
				fmt.Printf("\t%s: %s\n", getStatusText(fileStatus.Worktree), path)
				unStagedChanges = true
			}
		}
		if !unStagedChanges {
			fmt.Println("\t(no changes modified in worktree)")
		}

		fmt.Println("\nUntracked files:")
		untrackedFiles := false
		for path, fileStatus := range status {
			if fileStatus.Worktree == git.Untracked && fileStatus.Staging == git.Untracked {
				fmt.Printf("\t%s\n", path)
				untrackedFiles = true
			}
		}
		if !untrackedFiles {
			fmt.Println("\t(no untracked files)")
		}

		// TODO: Add ahead/behind remote status (more complex, requires fetching remote refs)
	},
}

func getStatusText(statusCode git.StatusCode) string {
	switch statusCode {
	case git.Unmodified:
		return "unmodified"
	case git.Added:
		return "new file"
	case git.Deleted:
		return "deleted"
	case git.Modified:
		return "modified"
	case git.Renamed:
		return "renamed"
	case git.Copied:
		return "copied"
	case git.UpdatedButUnmerged:
		return "unmerged"
	case git.Untracked:
		return "untracked"
	default:
		return "unknown"
	}
}

func init() {
	repoCmd.AddCommand(statusCmd)

	// Here you can define flags for the status command, e.g., --short, --porcelain
}

