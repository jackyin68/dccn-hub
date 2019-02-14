package dbservice

import (
	dbcommon "github.com/Ankr-network/dccn-common/db"
	pb "github.com/Ankr-network/dccn-common/protos/usermgr/v1/micro"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mysql or mongodb. who is better.
type DBService interface {
	// Get get a user item by go_micro_srv_user mgr's email
	Get(email string) (*pb.User, error)
	// Create Creates a new user item if not exits
	Create(user *pb.User) error
	// Update updates user item
	Update(user *pb.User) error
	// UpdatePassword updates user password
	UpdatePassword(email string, password []byte) error
	// UpdateActivateStatus updates user activate status
	UpdateActivateStatus(userId string) error
	// Close closes db connection
	Close()
	// dropCollection for testing usage
	dropCollection()
}

// DB implements DBService
type DB struct {
	dbName         string
	collectionName string
	session        *mgo.Session
}

// New returns DBService.
func New(conf dbcommon.Config) (*DB, error) {
	session, err := dbcommon.CreateDBConnection(conf)
	if err != nil {
		return nil, err
	}

	return &DB{
		dbName:         conf.DB,
		collectionName: conf.Collection,
		session:        session,
	}, nil
}

func (p *DB) collection(session *mgo.Session) *mgo.Collection {
	return session.DB(p.dbName).C(p.collectionName)
}

// Get gets user item by email.
func (p *DB) Get(email string) (*pb.User, error) {
	session := p.session.Clone()
	defer session.Close()

	var user pb.User
	err := p.collection(session).Find(bson.M{"email": email}).One(&user)
	return &user, err
}

// Create creates a new user item if it not exists
// TODO: batch operations through bulk
func (p *DB) Create(user *pb.User) error {
	session := p.session.Clone()
	defer session.Close()
	return p.collection(session).Insert(user)
}

// Update updates user item.
func (p *DB) Update(user *pb.User) error {
	session := p.session.Clone()
	defer session.Close()
	return p.collection(session).Update(bson.M{"email": user.Email}, user)
}

func (p *DB) UpdatePassword(email string, password []byte) error {
	session := p.session.Clone()
	defer session.Close()
	return p.collection(session).Update(bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})
}

func (p *DB) UpdateActivateStatus(email string) error {
	session := p.session.Clone()
	defer session.Close()
	return p.collection(session).Update(bson.M{"email": email}, bson.M{"$set": bson.M{"isactivation": true}})
}

// Close closes the db connection.
func (p *DB) Close() {
	p.session.Close()
}

func (p *DB) dropCollection() {
	p.session.DB(p.dbName).C(p.collectionName).DropCollection()
}
