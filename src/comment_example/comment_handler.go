package comment_example

import "github.com/manojpawar94/grpc/src/comment_example/commentpb"

func NewDummyComment() *commentpb.Comment {
	cmt := &commentpb.Comment{
		PostId:      1234,
		CommentId:   1,
		CommentText: "This is dummy commnet",
		Reactions:   []*commentpb.Reaction{{UserId: 12, ReactionType: commentpb.ReactionType_HAPPY}},
		Comments:    []*commentpb.Comment{},
	}

	return cmt
}
