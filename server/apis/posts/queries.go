package post_apis

import "fmt"

func getPostListQuery(inQuery string, pageNo int, pageSize int) string {
	query := fmt.Sprintf(`SELECT P.ID AS ID, P.USER_UID AS USER_UID,
	JSON_BUILD_OBJECT('uid', U.UID, 'name', U.NAME, 'avatar', U.AVATAR) AS POSTED_BY,
	P.CONTENT AS CONTENT, P.TEMPLATE AS TEMPLATE, P.TYPE AS TYPE, TO_CHAR(P.CREATED_ON, 'HH12:MI:SS') AS CREATED_ON,
	P.META_DATA AS META_DATA, P.TAGS AS TAGS, P.CATEGORY AS CATEGORY,
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
			FROM USERS TU WHERE TU.UID = ANY (P.TAGGED_USERS)
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
								FROM ATTACHMENTS AS CMA
								WHERE CMT.ATTACHMENT_ID = CMA.ID
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
								FROM USERS AS CMU
								WHERE CMT.USER_ID = CMU.ID LIMIT 5
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
			FROM COMMENTS CMT
			WHERE CMT.POST_ID = P.ID
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
OFFSET %d;`, inQuery, pageSize, pageNo)
	return query
}

func getPostQuery(uid string) string {
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
    P.ID = %s`, uid)
	return query
}
