package mysql

import (
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