package post_apis

import (
	"context"
	"database/sql"
	"fmt"

	post "grpc.worklen.com/proto/generated"
	pdb "grpc.worklen.com/server/apis"
)

// ListPosts fetches a list of posts from the database
func ListPosts(ctx context.Context, db *pdb.PostgresDb, in *post.ListPostsRequest) (*post.PostResponse, error) {
	var postList []*post.Post
	inQuery := "%" + in.Query + "%"
	query := getPostListQuery(inQuery, int(in.PageNo), int(in.PageSize))
	rows, err := db.Query(ctx, query)
	if err != nil {
		return &post.PostResponse{
			Message: err.Error(),
			Status:  "failed",
			Type:    "get",
			Post:    postList,
		}, nil
	}

	for rows.Next() {
		var p post.Post
		err := rows.Scan(
			&p.Id,            //0
			&p.UserUid,       //1
			&p.PostedBy,      //2
			&p.Content,       //3
			&p.Template,      //4
			&p.Type,          //5
			&p.CreatedOn,     //6
			&p.MetaData,      //7
			&p.Tags,          //8
			&p.Category,      //9
			&p.TaggedUsers,   //10
			&p.Comments,      //11
			&p.TotalComments, //12
			&p.Attachments,   //13
			&p.Thumbnail,     //14
			&p.IsSharedId,    //15
			&p.Options,       //16
		)
		if err != nil {
			return &post.PostResponse{
				Message: err.Error(),
				Status:  "failed",
				Type:    "get",
				Post:    postList,
			}, nil
		}
		postList = append(postList, &p)
	}
	defer rows.Close()
	return &post.PostResponse{
		Message: "Success !",
		Status:  "success",
		Type:    "list",
		Post:    postList,
	}, nil
}

// CreatePost creates a new post in the database
func CreatePost(ctx context.Context, db *pdb.PostgresDb, in *post.Post) (*post.PostResponse, error) {

	dt, err := db.Exec(ctx, `INSERT INTO posts(
		user_uid, content, template, type, meta_data, tags, category, tagged_users, thumbnail, is_shared_id, is_delete)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		in.UserUid,
		in.Content,
		in.Template,
		in.Type,
		in.MetaData,
		in.Tags,
		in.Category,
		in.TaggedUsers,
		in.Thumbnail,
		in.IsSharedId,
		false,
	)
	if err != nil {
		return &post.PostResponse{
			Message: err.Error(),
			Status:  "failed",
			Type:    "create",
			Post:    nil,
		}, nil
	}

	pst, err := GetPost(ctx, db, &post.GetPostRequest{Uuid: dt.String()})
	if err != nil {
		return pst, err
	}

	return pst, nil
}

// GetPost fetches a single post from the database by ID
func GetPost(ctx context.Context, db *pdb.PostgresDb, in *post.GetPostRequest) (*post.PostResponse, error) {
	var postList []*post.Post
	query := getPostQuery(in.Uuid)
	row := db.QueryRow(ctx, query)

	// Populate the post object from the row
	var p post.Post
	if err := row.Scan(
		&p.Id,          //0
		&p.UserUid,     //1
		&p.PostedBy,    //2
		&p.Content,     //3
		&p.Template,    //4
		&p.Type,        //5
		&p.CreatedOn,   //6
		&p.MetaData,    //7
		&p.Tags,        //8
		&p.Category,    //9
		&p.TaggedUsers, //10
		&p.Attachments, //11
		&p.Thumbnail,   //12
		&p.IsSharedId,  //13
		&p.Options,     //14
	); err != nil {
		if err == sql.ErrNoRows {
			return &post.PostResponse{
				Message: fmt.Sprintf("Post with ID %s not found", in.Uuid),
				Status:  "failed",
				Type:    "get",
				Post:    postList,
			}, nil
		}
		return &post.PostResponse{
			Message: err.Error(),
			Status:  "failed",
			Type:    "get",
			Post:    postList,
		}, nil
	}
	postList = append(postList, &p)

	return &post.PostResponse{
		Message: "Success !",
		Status:  "success",
		Type:    "get",
		Post:    postList,
	}, nil
}

// UpdatePost updates an existing post in the database
func UpdatePost(ctx context.Context, db *pdb.PostgresDb, in *post.Post) (*post.PostResponse, error) {
	_, err := db.Exec(ctx, "UPDATE posts SET content = $2 WHERE id = $1", in.Id, in.Content /* Add other fields */)
	if err != nil {
		return &post.PostResponse{
			Message: err.Error(),
			Status:  "failed",
			Type:    "update",
			Post:    nil,
		}, nil
	}

	return &post.PostResponse{
		Message: "Updated Scuuesfully !",
		Status:  "success",
		Type:    "update",
		Post:    nil,
	}, nil
}

// DeletePost deletes a post from the database by ID
func DeletePost(ctx context.Context, db *pdb.PostgresDb, in *post.DeletePostRequest) (*post.PostResponse, error) {
	query := fmt.Sprintf(`UPDATE posts SET is_delete=true WHERE posts.ID = %s;`, in.Uuid)
	_, err := db.Exec(ctx, query)
	if err != nil {
		return &post.PostResponse{
			Message: err.Error(),
			Status:  "failed",
			Type:    "delete",
			Post:    nil,
		}, nil
	}

	return &post.PostResponse{
		Message: "Deleted Scuuesfully !",
		Status:  "success",
		Type:    "delete",
		Post:    nil,
	}, nil
}
