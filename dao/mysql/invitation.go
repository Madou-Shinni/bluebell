package mysql

import (
	"github.com/jmoiron/sqlx"
	"strings"
	"web_app/models"
)

// InsertInvitation 插入帖子记录
func InsertInvitation(i *models.Invitation) (err error) {
	sqlStr := `insert into invitation
              (id,invitation_id,title,content,author_id,community_id,status,create_time,update_time)
               values(?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, i.Id, i.InvitationId, i.Title, i.Content, i.AuthorId, i.CommunityId, i.Status, i.CreateTime, i.UpdateTime)
	return
}

// SelectInvitationById 根据帖子id查询帖子详情
func SelectInvitationById(iId int64) (i *models.Invitation, err error) {
	i = new(models.Invitation)
	sqlStr := `select id,invitation_id,title,content,author_id,community_id,status,create_time,update_time
               from invitation where id = ?`
	err = db.Get(i, sqlStr, iId)
	return
}

// SelectInvitationList 查询帖子列表
func SelectInvitationList(pageNum, pageSize int64) (i []*models.Invitation, err error) {
	sqlStr := `select id,invitation_id,title,content,author_id,community_id,status,create_time,update_time
               from invitation limit ?,?`
	i = make([]*models.Invitation, 0, 2) // 参数2：长度 参数3：容量
	err = db.Select(&i, sqlStr, (pageNum-1)*pageSize, pageSize)
	return
}

// SelectInvitationListByIds 根据给定的id列表查询帖子数据
func SelectInvitationListByIds(ids []string) (invitationList []*models.Invitation, err error) {
	sqlStr := `select invitation_id,title,content,author_id,community_id,create_time from invitation
               where invitation_id in (?)
			   order by FIND_IN_SET(invitation_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&invitationList, query, args...)
	return
}
