package main

import ( "fmt" )

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func ( ai authenticationInfo) getBasicAuth() ( auth_string string )  {

	auth_string = fmt.Sprintf("Authorization: Basic %v:%v",ai.username,ai.password)

	return
}

func main(){

	user1_auth_info := authenticationInfo{username:"user1",password:"password1"}
	fmt.Println(user1_auth_info.getBasicAuth())

}