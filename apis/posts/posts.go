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
	inQuery := "%" + in.Query + "%"
	query := fmt.Sprintf(`SELECT
    P.ID AS ID,
    P.USER_UID AS USER_UID,
    JSON_BUILD_OBJECT(
        'uid', U.UID,
        'name', U.NAME,
        'avatar', U.AVATAR
    ) AS POSTED_BY,
    P.CONTENT AS CONTENT,
    P.TEMPLATE AS TEMPLATE,
    P.TYPE AS TYPE,
    to_char(P.created_on, 'HH12:MI:SS') AS CREATED_ON,
    P.META_DATA AS META_DATA,
    P.TAGS AS TAGS,
    P.CATEGORY AS CATEGORY,
    COALESCE((
        SELECT JSON_AGG(
            JSON_BUILD_OBJECT(
                'uid', TU.UID,
                'name', TU.NAME,
                'avatar', TU.AVATAR
            )
        )
        FROM USERS TU
        WHERE TU.UID = ANY (P.TAGGED_USERS)
    ), '[]') AS TAGGED_USERS,
    COALESCE((
        CASE 
            WHEN COUNT(A.ID) > 0 THEN JSON_AGG(
                JSON_BUILD_OBJECT(
                    'url', A.URL,
                    'ref', A.REF,
                    'name', A.NAME,
                    'type', A.TYPE,
                    'blur_hash', A.BLUR_HASH,
                    'thumbnail', A.THUMBNAIL,
                    'local_upload_ref', A.LOCAL_UPLOAD_REF
                )
            )
            ELSE '[]' 
        END
    ), '[]') AS ATTACHMENTS,
    P.THUMBNAIL AS THUMBNAIL,
    COALESCE(CAST(P.is_shared_id AS VARCHAR), '0') AS IS_SHARED_ID,
    COALESCE((
        CASE 
            WHEN COUNT(PO.OPTION_ID) > 0 THEN JSON_AGG(
                JSON_BUILD_OBJECT(
                    'option_id', PO.OPTION_ID,
                    'option_type', PO.OPTION_TYPE,
                    'metadata', PO.METADATA,
                    'content', PO.CONTENT,
                    'total_likes', COALESCE((
                        SELECT COUNT(POR.USER_ID)
                        FROM POST_OPTION_RESPONSES AS POR
                        WHERE POR.OPTION_ID = PO.OPTION_ID
                    ), 0)
                )
            )
            ELSE '[]'
        END
    ), '[]') AS OPTIONS
FROM
    POSTS P
JOIN USERS U ON P.USER_UID = U.UID
LEFT JOIN ATTACHMENTS A ON P.ID = A.POST_ID
LEFT JOIN POST_OPTIONS PO ON P.ID = PO.POST_ID
WHERE
    P.CONTENT LIKE '%s'
GROUP BY
    P.ID,
    U.UID,
    U.NAME,
    U.AVATAR
ORDER BY
    P.CREATED_ON DESC
LIMIT %d
OFFSET %d;`, inQuery, in.PageSize, in.PageNo)
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
