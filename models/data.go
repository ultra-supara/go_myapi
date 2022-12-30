package models

import "time"

var (
	// ArticleList is a list of articles

	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "This is a comment of test1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "This is a comment of test2",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is a first article",
		UserName:    "ultra-supara",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "second article",
		Contents:  "This is a second article",
		UserName:  "ultra-supara",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
