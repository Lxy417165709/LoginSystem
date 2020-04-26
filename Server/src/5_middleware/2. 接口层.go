package middleware

import "4_controls"

var Test 			= NewHandler(controls.Test, []middleWareType{TimesCounter, TimeCounter}).Format()
var Login 			= NewHandler(controls.Login, []middleWareType{TimesCounter,TimeCounter}).Format()
var Register 		= NewHandler(controls.Register, []middleWareType{TimesCounter, TimeCounter}).Format()
var UpdateUpi		= NewHandler(controls.UpdateUpi, []middleWareType{TimesCounter, TimeCounter}).Format()
var GetUpi			= NewHandler(controls.GetUpi, []middleWareType{TimesCounter, TimeCounter}).Format()
var SendRegisterVrc = NewHandler(controls.SendRegisterVrc, []middleWareType{TimesCounter, TimeCounter}).Format()
var GetPhoto 		= NewHandler(controls.GetPhoto, []middleWareType{TimesCounter, TimeCounter}).Format()
var UpdatePhoto 	= NewHandler(controls.UpdatePhoto, []middleWareType{TimesCounter, TimeCounter}).Format()
var GetUai 			= NewHandler(controls.GetUai, []middleWareType{TimesCounter, TimeCounter}).Format()
