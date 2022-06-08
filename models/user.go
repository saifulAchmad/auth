package models

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)
var DB *sql.DB

func ConnectDB() error{
	db,err:=sql.Open("sqlite3","./models/db/db")
	if err!=nil {
		return err
	}
	DB=db
	return nil
}

type User struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}





func GetUser()([]User,error){
	rows,err:= DB.Query(`SELECT * FROM user`)
	if err!=nil {
		return nil,err 
	}
	defer rows.Close()

	users:=make([]User,0)
	for rows.Next(){
		user:=User{}
		err:= rows.Scan(&user.Id,&user.Username,&user.Password)
		if err!=nil {
			return nil,err 
		}
		users= append(users, user)
	}
	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users,err
}


func GetUserById(id string)(User,error){
	sqlstmt,err:= DB.Prepare(`SELECT * FROM user WHERE id = ?`)
	if err!=nil {
		return User{},err
	}
	user:=User{}
	rows:=sqlstmt.QueryRow(id).Scan(&user.Id,&user.Username,&user.Password)
if rows!=nil {
	if rows==sql.ErrNoRows {
		return User{},nil
	}
	return User{}, rows

}
return user,nil
}

func AddUser(newUser User)(bool,error){
	tx,err:= DB.Begin()
	if err!=nil {
		return false,err
	}

	sqlstmt,err:=tx.Prepare(`INSERT INTO user (username,password)VALUES (?,?)`)
	if err!=nil {
		return false,err
	}
	defer sqlstmt.Close()
	_,Err:= sqlstmt.Exec(newUser.Username,newUser.Password)
	if Err!=nil {
		return false,err
	}
	tx.Commit()
	return true,nil
}





func UpdateUser(users User, id int)(bool,error){

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE user SET username = ? password = ? WHERE Id = ?")
	
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(users.Username,users.Password,users.Id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteUser( id int )(bool,error){
	tx,err:=DB.Begin()
	if err!=nil {
		return false,err
	}
	stmt,err:=DB.Prepare("DELETE from user WHERE id = ?")
	if err!=nil {
		return false,err
	}
	defer stmt.Close()

	_,err= stmt.Exec(id)
	if err!=nil {
		return false,err
	}
	tx.Commit()
	return true,nil
}