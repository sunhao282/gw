package backend

type MongoInfo struct {
	Id         string `bson:"id" json:"id"`
	GroupId    string `bson:"group_id" json:"group_id" form:"group_id" binding:"required"`
	Name       string `bson:"name" json:"name" form:"name" binding:"required"`
	Path       string `bson:"path" json:"path" form:"path" binding:"required"`
	To         string `bson:"to" json:"to" form:"to" binding:"required"`
	Method     string `bson:"method" json:"method" form:"method" binding:"oneof=GET POST get post"`
	Dns        int    `bson:"dns" json:"dns" form:"dns" binding:"oneof=0 1"`
	Flow       int    `bson:"flow" json:"flow" form:"flow" binding:"min=0"`
	CacheTime  int    `bson:"cache_time" json:"cache_time" form:"cache_time"`
	Timeout    int    `bson:"timeout" json:"timeout" form:"timeout" binding:"required"`
	Decay      int    `bson:"decay" json:"decay" form:"decay" binding:"oneof=0 1"`
	DecayTime  int    `bson:"decay_time" json:"decay_time" form:"decay_time"`
	CreateTime string `bson:"create_time" json:"create_time"`
	UpdateTime string `bson:"update_time" json:"update_time"`
}

type MongoGroup struct {
	Id         string `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name" form:"name" binding:"required"`
	Group      string `bson:"group" json:"group" form:"group" binding:"required"`
	CreateTime string `bson:"create_time" json:"create_time"`
	UpdateTime string `bson:"update_time" json:"update_time"`
}

type MongoAuth struct {
	Id         string `bson:"id" json:"id"`
	Src        string `bson:"src" json:"src" form:"src" binding:"required"`
	Key        string `bson:"key" json:"key" form:"key"`
	Content    string `bson:"content" json:"content" form:"content" binding:"required"`
	UserName   string `bson:"user_name" json:"user_name" form:"user_name" binding:"required"`
	UserEmail  string `bson:"user_email" json:"user_email" form:"user_email" binding:"required"`
	CreateTime string `bson:"create_time" json:"create_time"`
	UpdateTime string `bson:"update_time" json:"update_time"`
}

type MongoWgListApi struct {
	Id         string `bson:"id" json:"id" form:"id"`
	GroupId    string `bson:"group_id" json:"group_id" form:"group_id"`
	Name       string `bson:"name" json:"name" form:"name"`
}
