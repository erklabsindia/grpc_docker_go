package post_apis

import (
	"context"
	"database/sql"
	"fmt"
	pdb "worklen/apis"
	post "worklen/proto/proto/posts"
)

// ListPosts fetches a list of posts from the database
func ListPosts(ctx context.Context, db *pdb.PostgresDb, in *post.ListPostsRequest) (*post.PostResponse, error) {
	var postList []*post.Post
	// Your SQL query to fetch posts goes here
	var query string = `SELECT 
    p.id AS id,
    p.user_uid AS user_uid,
	json_object_agg('user',json_build_object(
        'uid', u.uid,
        'name', u.name,
        'avatar', u.avatar
    )) AS user,
    p.content AS content,
    p.template AS template,
    p.type AS type,
    to_char(p.created_on,'HH12:MI:SS') AS created_on,
    p.meta_data AS meta_data,
    p.tags AS tags,
    p.category AS category,
    (
        SELECT json_agg(json_build_object(
            'uid', tu.uid,
            'name', tu.name,
            'avatar', tu.avatar
        )) FROM users tu WHERE tu.uid = ANY(p.tagged_users)
    ) AS tagged_users,
    p.thumbnail AS thumbnail,
    COALESCE(cast(p.is_shared_id as varchar),'0') AS is_shared_id,
    json_agg(json_build_object(
        'url', a.url,
        'ref', a.ref,
        'name', a.name,
        'type', a.type,
        'blur_hash', a.blur_hash,
        'thumbnail', a.thumbnail,
        'local_upload_ref', a.local_upload_ref
    )) AS attachments,
    json_agg(json_build_object(
        'option_type', po.option_type,
        'metadata', po.metadata
    )) AS options
FROM 
    posts p
JOIN 
    users u ON p.user_uid = u.uid
LEFT JOIN 
    attachments a ON p.id = a.post_id
LEFT JOIN 
    post_options po ON p.id = po.post_id
GROUP BY 
    p.id, u.name, u.avatar;`

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
		// Scan values from the current row into Post struct
		err := rows.Scan(
			&p.Id,
			&p.UserUid,
			&p.User,
			&p.Content,
			&p.Template,
			&p.Type,
			&p.CreatedOn,
			&p.MetaData,
			&p.Tags,
			&p.Category,
			&p.TaggedUsers,
			&p.Thumbnail,
			&p.IsSharedId,
			&p.Attachments,
			&p.Options,
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
		Type:    "get",
		Post:    postList,
	}, nil
}

// CreatePost creates a new post in the database
func CreatePost(ctx context.Context, db *pdb.PostgresDb, in *post.Post) (*post.PostResponse, error) {
	// Your SQL query to insert a post goes here
	_, err := db.Exec(ctx, "INSERT INTO posts VALUES ($1, $2, $3, ...)", in.Id, in.UserUid, in.Content /* Add other fields */)
	if err != nil {
		return nil, err
	}

	// Return the response
	return &post.PostResponse{
		Message: "Post created successfully",
	}, nil
}

// GetPost fetches a single post from the database by ID
func GetPost(ctx context.Context, db *pdb.PostgresDb, in *post.GetPostRequest) (*post.PostResponse, error) {
	var postList []*post.Post
	// Your SQL query to fetch a post by ID goes here
	row := db.QueryRow(ctx, "SELECT * FROM posts WHERE id = $1", in.Uuid)

	// Populate the post object from the row
	var p post.Post
	if err := row.Scan(&p.Id, &p.UserUid, &p.Content /* Add other fields */); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post with ID %s not found", in.Uuid)
		}
		return nil, err
	}
	postList = append(postList, &p)

	// Construct and return the response
	return &post.PostResponse{
		Post: postList,
	}, nil
}

// UpdatePost updates an existing post in the database
func UpdatePost(ctx context.Context, db *pdb.PostgresDb, in *post.Post) (*post.PostResponse, error) {
	// Your SQL query to update a post goes here
	_, err := db.Exec(ctx, "UPDATE posts SET content = $2 WHERE id = $1", in.Id, in.Content /* Add other fields */)
	if err != nil {
		return nil, err
	}

	// Return the response
	return &post.PostResponse{
		Message: "Post updated successfully",
	}, nil
}

// DeletePost deletes a post from the database by ID
func DeletePost(ctx context.Context, db *pdb.PostgresDb, in *post.DeletePostRequest) (*post.PostResponse, error) {
	// Your SQL query to delete a post goes here
	_, err := db.Exec(ctx, "DELETE FROM posts WHERE id = $1", in.Uuid)
	if err != nil {
		return nil, err
	}

	// Return the response
	return &post.PostResponse{
		Message: "Post deleted successfully",
	}, nil
}
