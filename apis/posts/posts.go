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
	inQuery := "%" + in.Query + "%"
	query := fmt.Sprintf(`SELECT
	P.ID AS ID,
	P.USER_UID AS USER_UID,
	JSON_BUILD_OBJECT('uid', U.UID, 'name', U.NAME, 'avatar', U.AVATAR) AS POSTED_BY,
	P.CONTENT AS CONTENT,
	P.TEMPLATE AS TEMPLATE,
	P.TYPE AS TYPE,
	TO_CHAR(P.CREATED_ON, 'HH12:MI:SS') AS CREATED_ON,
	P.META_DATA AS META_DATA,
	P.TAGS AS TAGS,
	P.CATEGORY AS CATEGORY,
	COALESCE(
		(
			SELECT
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'uid',
						TU.UID,
						'name',
						TU.NAME,
						'avatar',
						TU.AVATAR
					)
				)
			FROM
				USERS TU
			WHERE
				TU.UID = ANY (P.TAGGED_USERS)
		),
		'[]'
	) AS TAGGED_USERS,
	COALESCE(
		(
			SELECT
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'id',
						CMT.ID,
						'content',
						CMT.CONTENT,
						'attachment',
						COALESCE(
							(
								SELECT
									JSON_BUILD_OBJECT(
										'url',
										CMA.URL,
										'ref',
										CMA.REF,
										'name',
										CMA.NAME,
										'type',
										CMA.TYPE,
										'blur_hash',
										CMA.BLUR_HASH,
										'thumbnail',
										CMA.THUMBNAIL,
										'local_upload_ref',
										CMA.LOCAL_UPLOAD_REF
									)
								FROM
									ATTACHMENTS AS CMA
								WHERE
									CMT.ATTACHMENT_ID = CMA.ID
							),
							NULL
						),
						'user',
						COALESCE(
							(
								SELECT
									JSON_BUILD_OBJECT(
										'uid',
										CMU.UID,
										'name',
										CMU.NAME,
										'avatar',
										CMU.AVATAR
									)
								FROM
									USERS AS CMU
								WHERE
									CMT.USER_ID = CMU.ID
								LIMIT
									5
							),
							NULL
						),
						'type',
						CMT.TYPE,
						'created_on',
						CMT.CREATED_ON,
						'parent_id',
						CMT.PARENT_ID
					)
				)
			FROM
				COMMENTS CMT
			WHERE
				CMT.POST_ID = P.ID
		),
		'[]'
	) AS COMMENTS,
	COALESCE(
		(
			SELECT
				COUNT(CMT.USER_ID)
			FROM
				COMMENTS CMT
			WHERE
				CMT.POST_ID = P.ID
		),
		0
	) AS TOTAL_COMMENTS,
	COALESCE(
		(
			CASE
				WHEN COUNT(A.ID) > 0 THEN JSON_AGG(
					JSON_BUILD_OBJECT(
						'url',
						A.URL,
						'ref',
						A.REF,
						'name',
						A.NAME,
						'type',
						A.TYPE,
						'blur_hash',
						A.BLUR_HASH,
						'thumbnail',
						A.THUMBNAIL,
						'local_upload_ref',
						A.LOCAL_UPLOAD_REF
					)
				)
				ELSE '[]'
			END
		),
		'[]'
	) AS ATTACHMENTS,
	P.THUMBNAIL AS THUMBNAIL,
	COALESCE(CAST(P.IS_SHARED_ID AS VARCHAR), '0') AS IS_SHARED_ID,
	COALESCE(
		(
			CASE
				WHEN COUNT(PO.OPTION_ID) > 0 THEN JSON_AGG(
					JSON_BUILD_OBJECT(
						'option_id',
						PO.OPTION_ID,
						'option_type',
						PO.OPTION_TYPE,
						'metadata',
						PO.METADATA,
						'content',
						PO.CONTENT,
						'total_likes',
						COALESCE(
							(
								SELECT
									COUNT(POR.USER_ID)
								FROM
									POST_OPTION_RESPONSES AS POR
								WHERE
									POR.OPTION_ID = PO.OPTION_ID
							),
							0
						)
					)
				)
				ELSE '[]'
			END
		),
		'[]'
	) AS OPTIONS
FROM
	POSTS P
	JOIN USERS U ON P.USER_UID = U.UID
	LEFT JOIN ATTACHMENTS A ON P.ID = A.POST_ID
	LEFT JOIN POST_OPTIONS PO ON P.ID = PO.POST_ID
WHERE
	P.CONTENT LIKE '%s'
	AND P.IS_DELETE = FALSE
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
    P.ID = %s`, in.Uuid)
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
