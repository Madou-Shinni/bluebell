package mysql

import "web_app/models"

// InsertInvitation 插入帖子记录
func InsertInvitation(i *models.Invitation) (err error) {
	sqlStr := `insert into invitation
              (id,invitation_id,title,content,author_id,community_id,status,create_time,update_time)
               values(?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, i.Id, i.InvitationId, i.Title, i.Content, i.AuthorId, i.CommunityId, i.Status, i.CreateTime, i.UpdateTime)
	return
}
